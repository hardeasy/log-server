package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/satori/go.uuid"
	"log"
	"log-server/internal/dto"
	"log-server/internal/models"
	"log-server/internal/utils"
	"time"
)

const (
	ES_Return_Max_Total = 10000 //ES使用from+size方式，为防止深分页最多只能返回的最大数量
	MAPPING = `{
	"mappings": {
		"_doc": {
			"properties": {
				"id": {
					"type":"keyword"
				},
				"time":{
					"type":"date"
				},
				"level":{
					"type":"keyword"
				},
				"content":{
					"type":"text"
				},
				"appcode":{
					"type":"keyword"
				}
			}
		}
	}
}`
)

type LogDao struct {

}

func (this LogDao) GetListEs(esClient *elastic.Client, d *dto.GeneralListDto) (rList []*models.Log, rSum int) {
	rList = []*models.Log{}
	rSum = 0

	logIndexName := fmt.Sprintf("%s*", this.GetLogIndexNamePrefix())

	search := esClient.Search().
		Index(logIndexName).   // search in index "twitter"
		Sort("time", false). // sort by "user" field, ascending
		Type("_doc").
		From(d.Offset).Size(d.Limit).   // take documents 0-9
		Pretty(true)       // pretty print request and response JSON

	if appCodeList,ok := d.Q["appCode"].([]string); ok && len(appCodeList) > 0 {
		boolQuery := elastic.NewBoolQuery()
		for _, appcode := range appCodeList {
			boolQuery.Should(elastic.NewTermQuery("appcode", appcode))
		}
		search.Query(boolQuery)
	}

	level := d.Q["level"].(string)
	if len(level) > 0 {
		termQuery := elastic.NewTermQuery("level", level)
		search.Query(termQuery)
	}

	content := d.Q["content"].(string)
	if len(content) > 0 {
		fullQuery := elastic.NewMatchQuery("content", content)
		search.Query(fullQuery)
	}

	ctx := context.Background()
	searchResult, err := search.Do(ctx)
	if err != nil {
		fmt.Println(d,err)
		return
	}
	rSum = int(searchResult.Hits.TotalHits)
	if rSum > ES_Return_Max_Total {
		rSum = ES_Return_Max_Total - d.Limit
	}
	if rSum > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var log models.Log
			err := json.Unmarshal(*hit.Source, &log)
			if err != nil {
				continue
			}
			log.Id = hit.Id
			rList = append(rList, &log)
		}
	}
	return
}

func (this LogDao) GetById(esClient *elastic.Client, id string) *models.Log {
	logIndexName := fmt.Sprintf("%s*", this.GetLogIndexNamePrefix())
	search := esClient.Search().Index(logIndexName).
		Type("_doc").
		From(0).Size(1).   // take documents 0-9
		Query(elastic.NewTermQuery("id", id))
	ctx := context.Background()
	searchResult, err := search.Do(ctx)
	if err != nil {
		return nil
	}
	total := int(searchResult.Hits.TotalHits)
	if total > 0{
		hit := searchResult.Hits.Hits[0]
		var log models.Log
		err := json.Unmarshal(*hit.Source, &log)
		if err != nil {
			return nil
		}
		return &log
	}
	return nil
}

func (this LogDao) AddPushLog(esClient *elastic.Client, d dto.PushLogDto) error {
	cxt := context.Background()
	id := uuid.NewV4().String()
	logModel := &models.Log{
		Id: id,
		Level: d.Level,
		Time:  time.Unix(int64(d.Time), 0),
		Content:  d.Content,
		Appcode: d.Appcode,
	}
	logIndexName := this.GetActiveIndexName()
	this.GetExistsAndCreate(esClient, logIndexName)
	_, err := esClient.Index().Index(logIndexName).Type("_doc").Id(id).BodyJson(logModel).Do(cxt)
	if err != nil {
		log.Println("pushlog error", err.Error())
	}
	return nil
}

func (LogDao) GetLogIndexNamePrefix() string {
	return "log"
}

func (this LogDao) GetExistsAndCreate(esClient *elastic.Client, indexName string) {
	ctx := context.Background()
	exists, _ := esClient.IndexExists(indexName).Do(ctx)
	if !exists {
		_, err:= esClient.CreateIndex(indexName).BodyString(MAPPING).Do(ctx)
		if err != nil {
			log.Println("create Index error", err.Error())
		}
	}
}

func (this LogDao) GetActiveIndexName() string {
	prefix := this.GetLogIndexNamePrefix()
	return fmt.Sprintf("%s-%s", prefix, time.Now().Format(utils.DateFormart))
}

func (this LogDao) GetIndexList(esClient *elastic.Client) []dto.Index {
	list := []dto.Index{}
	logIndexName := fmt.Sprintf("%s*", this.GetLogIndexNamePrefix())
	ctx := context.Background()
	result, err := esClient.CatIndices().Index(logIndexName).Do(ctx)
	if err != nil {
		log.Println("create Index error", err.Error())
		return list
	}
	for _, item := range result {
		list = append(list, dto.Index{
			Health:       item.Health,
			Status:       item.Status,
			Index:        item.Index,
			Uuid:         item.UUID,
			Pri:          item.Pri,
			Rep:          item.Rep,
			DocsCount:    item.DocsCount,
			DocsDeleted:  item.DocsDeleted,
			StoreSize:    item.StoreSize,
			PriStoreSize: item.PriStoreSize,
		})
	}
	return list
}

func (this LogDao) DeleteIndex(esClient *elastic.Client, indexName string) error {
	ctx := context.Background()
	_, err := esClient.DeleteIndex(indexName).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

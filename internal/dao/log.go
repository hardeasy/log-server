package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"log-server/internal/dto"
	"log-server/internal/models"
)

const (
	ES_Return_Max_Total = 10000 //ES使用from+size方式，为防止深分页最多只能返回的最大数量
)

type LogDao struct {

}

func (this LogDao) GetListEs(esClient *elastic.Client, d *dto.GeneralListDto) (rList []*models.Log, rSum int) {
	rList = []*models.Log{}
	rSum = 0

	appCode := d.Q["appCode"]

	logIndexName := this.GetLogIndexName(appCode)

	search := esClient.Search().
		Index(logIndexName).   // search in index "twitter"
		//Sort("time", false). // sort by "user" field, ascending
		Type("_doc").
		From(d.Offset).Size(d.Limit).   // take documents 0-9
		Pretty(true)       // pretty print request and response JSON

	level := d.Q["level"]
	if len(level) > 0 {
		termQuery := elastic.NewTermQuery("level", level)
		search.Query(termQuery)
	}

	content := d.Q["content"]
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

func (this LogDao) GetById(esClient *elastic.Client, appCode string, id string) *models.Log {
	cxt := context.Background()
	logIndexName := this.GetLogIndexName(appCode)
	result, err := esClient.Get().Index(logIndexName).Type("_doc").Id(id).Do(cxt)
	if err != nil {
		return nil
	}
	if result.Found {
		var log *models.Log
		err := json.Unmarshal(*result.Source, &log)
		if err == nil {
			log.Id = result.Id
			return log
		}
	}
	return nil
}

func (this LogDao) AddPushLog(esClient *elastic.Client, d dto.PushLogDto) error {
	cxt := context.Background()
	log := &models.Log{
		Id: "",
		Level: d.Level,
		Time:  d.Time,
		Content:  d.Content,
		Appcode: d.Appcode,
	}
	logIndexName := this.GetLogIndexName(d.Appcode)
	esClient.Index().Index(logIndexName).Type("_doc").BodyJson(log).Do(cxt)
	return nil
}

func (LogDao) GetLogIndexNamePrefix() string {
	return "log"
}

func (this LogDao) GetLogIndexName(appcode string) string {
	prefix := this.GetLogIndexNamePrefix()
	if len(appcode) == 0 {
		return fmt.Sprintf("%s*", prefix)
	}
	return fmt.Sprintf("%s-%s", prefix, appcode)
}

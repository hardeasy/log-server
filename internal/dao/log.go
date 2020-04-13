package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic"
	"log-server/internal/dto"
	"log-server/internal/models"
	"time"
)

type LogDao struct {

}

func (this LogDao) GetListEs(esClient *elastic.Client, d *dto.GeneralListDto) (rList []*models.Log, rSum int) {
	rList = []*models.Log{}
	rSum = 0
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("level", "error")

	index := fmt.Sprintf("%s-%s-*", this.GetNowLogIndexNamePrefix(), time.Now().Format("2006"))
	searchResult, _ := esClient.Search().
		Index(index).   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("time", false). // sort by "user" field, ascending
		From(d.Offset).Size(d.Limit).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             //
	rSum = int(searchResult.Hits.TotalHits)
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
	return;
}

func (LogDao) GetListDb(db *gorm.DB, d *dto.GeneralListDto) (rList []*models.Log, rSum int) {
	if url, ok := d.Q["url"]; ok && len(url) > 0 {
		db = db.Where("url like ?", fmt.Sprintf("%s%%", url))
	}
	if category, ok := d.Q["category"]; ok && len(category) > 0 {
		db = db.Where("category = ?", category)
	}
	if content, ok := d.Q["content"]; ok && len(content) > 0 {
		db = db.Where("content like ?", fmt.Sprintf("%%%s%%", content))
	}
	orderBy := "id desc"
	if len(d.Order)  > 0 {
		orderBy = d.Order
	}
	db.Order(orderBy).Offset(d.Offset).Limit(d.Limit).Model(&models.Log{}).Count(&rSum)
	db.Order(orderBy).Offset(d.Offset).Limit(d.Limit).Find(&rList)
	return
}

func (this LogDao) AddPushLog(esClient *elastic.Client, d dto.PushLogDto) error {
	cxt := context.Background()
	log := &models.Log{
		Id: "",
		Level: d.Level,
		Time:  d.Time,
		Data:  d.Data,
	}
	nowDayLogIndex := this.GetNowLogIndexName()
	esClient.Index().Index(nowDayLogIndex).Type("_doc").BodyJson(log).Do(cxt)
	return nil
}

func (LogDao) GetNowLogIndexNamePrefix() string {
	return "log"
}

func (this LogDao) GetNowLogIndexName() string {
	date := time.Now().Format("2006-01")
	return fmt.Sprintf("%s-%s", this.GetNowLogIndexNamePrefix(), date)
}
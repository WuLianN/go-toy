package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
)

func (d *Dao) VisitStatisticsByMonth(year string) []model.Result {
	list := []model.Result{}
	d.engine.Raw("SELECT DATE_FORMAT(visit_time,'%Y-%m') AS date, COUNT(visit_time) AS total FROM `statistics_visit` WHERE YEAR(visit_time) = ? GROUP BY(date)", year).Scan(&list)

	return list
}

func (d *Dao) VisitStatisticsByOneMonth(year int, month int) model.Result {
	var oneMonthResult model.Result 
	d.engine.Raw("SELECT DATE_FORMAT(visit_time,'%Y-%m') AS date, COUNT(visit_time) AS total FROM `statistics_visit` WHERE YEAR(visit_time) = ? AND MONTH(visit_time) = ? GROUP BY(date)", year, month).Scan(&oneMonthResult)

	return oneMonthResult
}
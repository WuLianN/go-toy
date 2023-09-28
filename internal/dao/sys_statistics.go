package dao

type result struct {
	Date string `json:"date"`
	Total int `json:"total"`
}

func (d *Dao) VisitStatisticsByMonth(year string) []result {
	list := []result{}
	d.engine.Raw("SELECT DATE_FORMAT(visit_time,'%Y-%m') AS date, COUNT(visit_time) AS total FROM `statistics_visit` WHERE YEAR(visit_time) = ? GROUP BY(date)", year).Scan(&list)

	return list
}
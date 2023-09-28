package service

import (
	"github.com/WuLianN/go-toy/pkg/convert"
	"strconv"
)


func (svc *Service) GetVisitStatistics(year string) ([12]string, [12]int){
	list := svc.dao.VisitStatisticsByMonth(year)

	valueList := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	dateList := [12]string{}

	for i := 0; i < 12; i++ {
		if i < 9 {
			dateList[i] = year + "-0" + strconv.Itoa(i + 1)
		} else {
			dateList[i] = year + "-" + strconv.Itoa(i + 1)
		}
	}

	if len(list) > 0 {
		for _, value := range list {
			month := value.Date[5:7]
			simpifyMonth := convert.StrTo(month).MustInt()

			switch simpifyMonth {
			case 1:
				valueList[0] = value.Total
			case 2:
				valueList[1] = value.Total
			case 3:
				valueList[2] = value.Total
			case 4:
				valueList[3] = value.Total
			case 5:
				valueList[4] = value.Total
			case 6:
				valueList[5] = value.Total
			case 7:
				valueList[6] = value.Total
			case 8:
				valueList[7] = value.Total
			case 9:
				valueList[8] = value.Total
			case 10:
				valueList[9] = value.Total
			case 11:
				valueList[10] = value.Total
			case 12:
				valueList[11] = value.Total
			}
		}
	}

	return dateList, valueList
}

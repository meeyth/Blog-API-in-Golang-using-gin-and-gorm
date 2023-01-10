package utils

import (
	"strconv"
	"strings"
	"time"
)

func CreateDate(strTime string) (time.Time, error) {
	dateDetails := strings.Split(strTime, ":")

	year, err1 := strconv.Atoi(dateDetails[0])
	month, err2 := strconv.Atoi(dateDetails[1])
	date, err3 := strconv.Atoi(dateDetails[2])

	if err1 != nil || err2 != nil || err3 != nil {
		panic(err1.Error() + err2.Error() + err3.Error())
	}

	createDate := time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.Local)

	return createDate, nil
}

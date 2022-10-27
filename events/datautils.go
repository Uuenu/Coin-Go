package events

import (
	"strconv"
	"telegram-coin-go/storage"
	"time"
)

func CheckTime(lrTime, timeNow time.Time) bool {
	lrDay, lrMonth, lrYear := lrTime.Date()
	nowDay, nowMonth, nowYear := timeNow.Date()

	return lrDay == nowDay && lrMonth == nowMonth && lrYear == nowYear
}

func RecData(chatID int, sum float64, s storage.Storage) (map[string]string, time.Time, error) {
	result := make(map[string]string, 0)
	lastDebit, lastCredit, recTime, err := s.LastRecord(chatID) // get Debit/Credit from last document in collection
	if err != nil {
		return nil, recTime, err
	}

	iDebit, err := strconv.Atoi(lastDebit)
	if err != nil {

	}

	iCredit, err := strconv.Atoi(lastCredit)
	if err != nil {

	}

	result["Debit"] = strconv.Itoa(iDebit - int(sum))
	result["Credit"] = strconv.Itoa(iCredit - int(sum))
	result["Sum"] = strconv.Itoa(int(sum))
	//result["Time"] = recTime.GoString()
	//result["Text"] = text
	return result, recTime, nil
}

package events

import (
	"strconv"
	lib "telegram-coin-go/lib/e"
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
	data, recTime, err := s.LastRecord(chatID) // get Debit, Credit and Time from last document in collection
	if err != nil {
		return nil, recTime, lib.Wrap("can't get last record from db", err)
	}

	fDebit, err := strconv.ParseFloat(data["Debit"], 64)
	if err != nil {
		return nil, recTime, lib.Wrap("can't get debit from db", err)
	}

	fCredit, err := strconv.ParseFloat(data["Credit"], 64)
	if err != nil {
		return nil, recTime, lib.Wrap("can't get credit from db", err)
	}

	result["Debit"] = strconv.FormatFloat(fDebit-sum, 'f', 5, 64)
	result["Credit"] = strconv.FormatFloat(fCredit-sum, 'f', 5, 64)
	result["Debit"] = ""
	result["Credit"] = ""
	result["Sum"] = strconv.FormatFloat(sum, 'f', 5, 64)
	return result, recTime, nil
}

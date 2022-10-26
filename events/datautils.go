package events

import (
	"strconv"
	"telegram-coin-go/storage"
)

func RecData(chatID int, sum string, s storage.Storage) (map[string]string, error) {
	result := make(map[string]string, 0)
	lastDebit, lastCredit, err := s.LastCredit(chatID) // get Debit/Credit from last document in collection
	if err != nil {
		return nil, err
	}

	iDebit, err := strconv.Atoi(lastDebit)
	if err != nil {

	}

	iCredit, err := strconv.Atoi(lastCredit)
	if err != nil {

	}

	isum, err := strconv.Atoi(sum)
	if err != nil {

	}

	result["Debit"] = strconv.Itoa(iDebit - isum)
	result["Credit"] = strconv.Itoa(iCredit - isum)
	result["Sum"] = strconv.Itoa(isum)
	//result["Text"] = text
	return result, nil
}

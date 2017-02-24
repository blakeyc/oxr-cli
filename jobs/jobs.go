package jobs

import (
	"log"
	"strings"

	"github.com/mvillalba/go-openexchangerates/oxr"
)

// GetLatest (appID, baseCurrency, output, fields)
// Fetches the latest rates data from the historical API endpoint and writes
// to file.
func GetLatest(appID string, baseCurrency string, output string, fields []string) (err error) {

	log.Print("info: starting latest job")

	var results = make([]*oxr.Rates, 1)
	client := oxr.New(appID)

	r, err := client.LatestWithOptions(strings.ToUpper(baseCurrency), nil)
	if err != nil {
		return err
	}

	results[0] = r

	log.Printf("info: recieved %d rates, timestamp %d", len(r.Rates), r.Timestamp)

	err = WriteToFile(output, results, fields)
	if err != nil {
		return err
	}
	log.Print("info: completed latest job")

	return nil
}

// GetHistorical (appID, baseCurrency, output, fields, dates)
// Fetches the specified date rates from the historical API endpoint and writes
// to file.
func GetHistorical(appID string, baseCurrency string, output string, fields []string, dates []string) (err error) {

	log.Print("info: starting historical job")

	var results = make([]*oxr.Rates, len(dates))
	client := oxr.New(appID)

	for i, date := range dates {

		log.Printf("info: fetching historical rates for %s", date)
		r, errClient := client.HistoricalWithOptions(date, strings.ToUpper(baseCurrency), nil)
		if errClient != nil {
			return errClient
		}

		results[i] = r

		log.Printf("info: recieved %d rates, timestamp %d", len(r.Rates), r.Timestamp)

	}

	err = WriteToFile(output, results, fields)
	if err != nil {
		return err
	}

	log.Print("info: completed historical job")

	return nil
}

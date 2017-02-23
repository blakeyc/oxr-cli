package jobs

import (
	"fmt"
	"log"
	"strings"

	"github.com/mvillalba/go-openexchangerates/oxr"
)

// GetLatest (appID, baseCurrency, file)
// Fetches the latest rates data from the historical API endpoint and writes
// to file.
func GetLatest(appID string, baseCurrency string, output string) (err error) {

	log.Print("info: starting latest job")

	var results = make([]*oxr.Rates, 1)
	client := oxr.New(appID)

	r, err := client.LatestWithOptions(strings.ToUpper(baseCurrency), nil)
	if err != nil {
		return err
	}

	results[0] = r

	log.Printf("info: recieved %d rates, timestamp %d", len(r.Rates), r.Timestamp)

	err = WriteToFile(output, results)
	if err != nil {
		return err
	}
	log.Print("info: completed latest job")

	return
}

// GetHistorical (appID, date, baseCurrency, file)
// Fetches the specified date rates from the historical API endpoint and writes
// to file.
func GetHistorical(appID string, dates []string, baseCurrency string, output string) (err error) {

	log.Print("info: starting historical job")

	fmt.Print(dates)

	var results = make([]*oxr.Rates, len(dates))
	client := oxr.New(appID)

	for i, date := range dates {

		log.Printf("info: fetching historical rates for %s", date)
		r, err := client.HistoricalWithOptions(date, strings.ToUpper(baseCurrency), nil)
		if err != nil {
			return err
		}

		results[i] = r

		log.Printf("info: recieved %d rates, timestamp %d", len(r.Rates), r.Timestamp)

	}

	err = WriteToFile(output, results)
	if err != nil {
		return err
	}

	log.Print("info: completed historical job")

	return
}

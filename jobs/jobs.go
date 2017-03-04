package jobs

import (
	"log"
	"strings"
	"time"

	"github.com/mvillalba/go-openexchangerates/oxr"
)

// GetLatest (appID, baseCurrency, output, fields)
// Fetches the latest rates data from the historical API endpoint and writes
// to file.
func GetLatest(appID string, baseCurrency string, output string, fields []string) (err error) {

	log.Print("info: starting latest job")

	results := make([]oxr.Rates, 1)
	client := oxr.New(appID)

	r, err := client.LatestWithOptions(strings.ToUpper(baseCurrency), nil)
	if err != nil {
		return err
	}

	results[0] = oxr.Rates{
		Disclaimer: r.Disclaimer,
		License:    r.License,
		Timestamp:  r.Timestamp,
		Base:       r.Base,
		Rates:      r.Rates,
	}

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

	results := make([]oxr.Rates, len(dates))
	client := oxr.New(appID)

	for i, date := range dates {

		log.Printf("info: fetching historical rates for %s", date)
		r, errClient := client.HistoricalWithOptions(date, strings.ToUpper(baseCurrency), nil)
		if errClient != nil {
			return errClient
		}

		results[i] = oxr.Rates{
			Disclaimer: r.Disclaimer,
			License:    r.License,
			Timestamp:  r.Timestamp,
			Base:       r.Base,
			Rates:      r.Rates,
		}

		log.Printf("info: recieved %d rates, timestamp %d", len(r.Rates), r.Timestamp)

	}

	err = WriteToFile(output, results, fields)
	if err != nil {
		return err
	}

	log.Print("info: completed historical job")

	return nil
}

// GetTimeSeries (appID, baseCurrency, output, fields, startDate, endDate)
// Fetches the latest rates data from the historical API endpoint and writes
// to file.
func GetTimeSeries(appID string, baseCurrency string, output string, fields []string, startDate string, endDate string) (err error) {

	log.Print("info: starting time series job")

	client := oxr.New(appID)

	r, err := client.TimeSeriesWithOptions(startDate, endDate, strings.ToUpper(baseCurrency), nil)
	if err != nil {
		return err
	}

	results := make([]oxr.Rates, len(r.Rates))

	i := 0
	for date, rates := range r.Rates {

		parseDate, _ := time.Parse("2006-01-02", date)
		timestamp := parseDate.Unix()

		log.Printf("Date: %s, TimeStamp: %d", date, timestamp)

		results[i] = oxr.Rates{
			Disclaimer: r.Disclaimer,
			License:    r.License,
			Timestamp:  int(timestamp),
			Base:       r.Base,
			Rates:      rates,
		}
		i++

	}
	log.Printf("info: recieved %d days of rates", len(r.Rates))

	err = WriteToFile(output, results, fields)
	if err != nil {
		return err
	}
	log.Print("info: completed time series job")

	return nil
}

package jobs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mvillalba/go-openexchangerates/oxr"
)

// GetHistorical (appID, date, baseCurrency, output)
// Takes an openexchangerates app_id as an input, queries the API
// for the latest exchange rates and saves to the required file.
func GetHistorical(appID string, date string, baseCurrency string, output string) (err error) {

	log.Printf("info: starting historical job - %s", date)

	base := strings.ToUpper(baseCurrency)

	client := oxr.New(appID)

	r, err := client.HistoricalWithOptions(date, base, nil)
	if err != nil {
		return err
	}

	file, err := os.Create(output)
	if err != nil {
		return err
	}

	defer file.Close()

	log.Printf("info: recieved %d rates", len(r.Rates))
	log.Print("info: writing to file")

	w := bufio.NewWriter(file)

	for k := range r.Rates {
		fmt.Fprintf(w, "%s,", date)
		fmt.Fprintf(w, "%s,", base)
		fmt.Fprintf(w, "%s,", k)
		fmt.Fprintf(w, "%s\n", r.Rates[k])
	}

	w.Flush()

	log.Printf("info: completed historical job - %s", date)

	return
}

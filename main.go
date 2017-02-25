package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/blakeyc/oxr-cli/jobs"
	"github.com/blakeyc/oxr-cli/utils"
)

// VERSION of application
const VERSION = "1.0.0"

var (
	showVersion  = flag.Bool("version", false, "print application version")
	output       = flag.String("output", "", "destination of the output file")
	job          = flag.String("job", "", "task type to perform (latest|historical)")
	appID        = flag.String("app_id", "", "your open exchange rates app_id")
	baseCurrency = flag.String("base", "USD", "which currency to use as the base, defaults to USD")
	dates        = flag.String("dates", "", "the dates to get historical rates for (YYYY-MM-DD)")
	fields       = flag.String("fields", "base,currency,rate,timestamp,date", "pick which fields to include in output (base,currency,rate,timestamp,date)")
)

func main() {
	var err error
	flag.Usage = usage
	flag.Parse()

	if *showVersion {
		version()
		return
	}

	// Check the required params have been supplied
	if *appID == "" {
		log.Fatal("error: missing --app_id")
	}

	if *job == "" {
		log.Fatal("error: missing --job")
	}

	if *output == "" {
		log.Fatal("error: missing --output")
	}

	// Compile array of fields to be used for output
	fields, errFields := utils.ExtractFields(*fields)
	if errFields != nil {
		log.Fatal(errFields)
	}

	// Execute the correct job from the supplied param
	switch *job {

	case "latest":

		err = jobs.GetLatest(*appID, *baseCurrency, *output, fields)

	case "historical":

		if *dates == "" {
			log.Fatal("error: missing --dates")
		}

		dates, errDates := utils.ExtractDates(*dates)
		if errDates != nil {
			log.Fatal(errDates)
		}
		err = jobs.GetHistorical(*appID, *baseCurrency, *output, fields, dates)

	default:

		log.Fatal("error: unrecognized job type, should be on of (latest|historical)")

	}

	if err != nil {
		log.Fatal(err)
	}

}

// print usage
func usage() {
	fmt.Fprint(os.Stderr, "usage: oxr [flags]\n")
	flag.PrintDefaults()
}

// print version
func version() {
	fmt.Printf("v%s\n", VERSION)
}

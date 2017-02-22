package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/blakeyc/oxr-cli/jobs"
)

// VERSION of application
const VERSION = "0.0.1"

var (
	showVersion  = flag.Bool("version", false, "print application version")
	output       = flag.String("output", "", "destination of the output file")
	job          = flag.String("job", "", "task type to perform (latest|daily|historical)")
	appID        = flag.String("app_id", "", "your open exchange rates app_id")
	baseCurrency = flag.String("base", "USD", "which currency to use as the base, defaults to USD")
	date         = flag.String("date", "", "the date to get historical rates for (YYYY-MM-DD)")
)

func main() {
	var err error
	flag.Usage = usage
	flag.Parse()

	if *showVersion {
		version()
		return
	}

	if *appID == "" {
		log.Fatal("error: missing --app_id")
	}

	if *job == "" {
		log.Fatal("error: missing --job")
	}

	if *output == "" {
		log.Fatal("error: missing --output")
	}

	switch *job {
	case "latest":
		err = jobs.GetLatest(*appID, *baseCurrency, *output)

	case "historical":
		if *date == "" {
			log.Fatal("error: missing --date")
		}
		err = jobs.GetHistorical(*appID, *date, *baseCurrency, *output)
	}

	if err != nil {
		log.Fatal(err)
	}

}

// print usage message
func usage() {
	fmt.Fprint(os.Stderr, "usage: gorates [flags]\n")
	flag.PrintDefaults()
}

// print application version
func version() {
	fmt.Printf("v%s\n", VERSION)
}

package jobs

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mvillalba/go-openexchangerates/oxr"
)

// WriteToFile (output, r)
// Writes the fetched data to file.
func WriteToFile(output string, results []*oxr.Rates, fields []string) (err error) {

	log.Printf("info: writing to file %s", output)

	file, err := os.Create(output)
	if err != nil {
		return err
	}

	defer file.Close()

	w := bufio.NewWriter(file)

	for _, r := range results {
		for k := range r.Rates {
			fmt.Fprintf(w, "%s", r.Base)
			fmt.Fprint(w, ",")
			fmt.Fprintf(w, "%s", k)
			fmt.Fprint(w, ",")
			fmt.Fprintf(w, "%s", r.Rates[k])
			fmt.Fprint(w, ",")
			fmt.Fprintf(w, "%d", r.Timestamp)
			fmt.Fprint(w, "\n")
		}
	}

	w.Flush()

	return
}

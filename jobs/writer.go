package jobs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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

			line := make([]string, len(fields))

			for i, field := range fields {

				switch field {
				case "base":
					line[i] = r.Base
				case "currency":
					line[i] = k
				case "rate":
					line[i] = string(r.Rates[k])
				case "timestamp":
					line[i] = strconv.Itoa(r.Timestamp)
				case "date":
					line[i] = time.Unix(int64(r.Timestamp), int64(0)).Format("2006-01-02")
				}

			}

			output := strings.Join(line, ",")

			fmt.Fprintf(w, "%s\n", output)
		}
	}

	w.Flush()

	return nil
}

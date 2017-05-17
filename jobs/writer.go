package jobs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/mvillalba/go-openexchangerates/oxr"
)

// ByTimestamp Interface for sort functions
type ByTimestamp []oxr.Rates

func (s ByTimestamp) Len() int {
	return len(s)
}
func (s ByTimestamp) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByTimestamp) Less(i, j int) bool {
	return s[i].Timestamp < s[j].Timestamp
}

// WriteToFile (output, r)
// Writes the fetched data to file.
func WriteToFile(output string, results []oxr.Rates, fields []string) (err error) {

	log.Printf("info: writing to file %s", output)

	file, err := os.Create(output)
	if err != nil {
		return err
	}

	defer file.Close()

	w := bufio.NewWriter(file)

	// Sort based on timestamp field
	sort.Sort(ByTimestamp(results))

	for _, r := range results {

		keys := []string{}
		for key := range r.Rates {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for k := range keys {
			key := keys[k]

			line := make([]string, len(fields))

			for i, field := range fields {

				switch field {
				case "base":
					line[i] = r.Base
				case "currency":
					line[i] = key
				case "rate":
					line[i] = string(r.Rates[key])
				case "timestamp":
					line[i] = strconv.Itoa(r.Timestamp)
				case "date":
					line[i] = time.Unix(int64(r.Timestamp), int64(0)).UTC().Format("2006-01-02")
				}

			}

			output := strings.Join(line, ",")

			fmt.Fprintf(w, "%s\n", output)
		}
	}

	w.Flush()

	return nil
}

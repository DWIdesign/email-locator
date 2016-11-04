package tld

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

// CountryTLD is responsible for matching top-level domain with country it
// represents.
type CountryTLD interface {
	CheckTLD(string) string
}

type countryTLD []countryCode

type countryCode struct {
	tld     string
	country string
}

// CheckTLD matches top-level domain with country it represents.
func (cslice countryTLD) CheckTLD(tld string) string {
	tld = strings.ToUpper(tld)
	for _, c := range cslice {
		if c.tld == tld {
			return c.country
		}
	}
	return ""
}

// LoadCSV parses a CSV file into a slice of tld-representation pairs
func LoadCSV(file string) countryTLD {
	f, err := os.Open(file)
	if err != nil {
		log.Println(err)
	}
	r := csv.NewReader(f)

	records, err := r.ReadAll()

	countries := make([]countryCode, len(records))
	for i, r := range records {
		countries[i] = countryCode{r[0], r[1]}
	}

	return countries
}

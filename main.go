package main

import (
	"fmt"
	"github.com/DWIdesign/email-locator/tld"
	"log"

	"encoding/csv"
	"os"
)

var countries = tld.LoadCSV("tld/country-locations-clean.csv")

func main() {
	testData := loadTestData()

	for _, testmail := range testData {
		location, err := tld.LocateEmail(testmail, countries)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(location)
	}
}

func loadTestData() []string {
	f, err := os.Open("tld/test_mail.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)

	elems, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	ret := make([]string, len(elems))
	for i, e := range elems {
		ret[i] = e[0]
	}
	return ret
}

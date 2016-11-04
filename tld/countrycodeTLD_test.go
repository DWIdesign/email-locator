package tld

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"
)

func TestTLD(t *testing.T) {
	s := `PL,Poland
DE,Germany
GB,United Kingdom
GS,South Georgia and the South Sandwich Islands
`
	csv := newFakeCSV(t, s)
	defer csv.close()

	codes := LoadCSV(csv.path)

	country := codes.CheckTLD("PL")
	if country != "Poland" {
		t.Errorf("Expected Poland, got %s", country)
	}
}

type FakeCSV struct {
	file *os.File
	path string
}

func newFakeCSV(t *testing.T, csv string) FakeCSV {

	ret := FakeCSV{}
	ret.file = makeUniqueFile(t, `^test([0-9]*).csv$`)
	ret.path = ret.file.Name()

	_, err := ret.file.WriteString(csv)
	checkErr(t, err)
	return ret
}

func (csv FakeCSV) close() {
	os.Remove(csv.path)
	csv.file.Close()
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func getDirNames(t *testing.T) []string {
	f, err := os.Open("/")
	checkErr(t, err)
	defer f.Close()

	names, err := f.Readdirnames(0)
	checkErr(t, err)

	return names
}

func makeUniqueFile(t *testing.T, form string) *os.File {
	names := getDirNames(t)

	rx, err := regexp.Compile(form)
	checkErr(t, err)
	max := 0
	for _, name := range names {
		namesm := rx.FindStringSubmatch(name)
		if namesm != nil {
			num, _ := strconv.Atoi(namesm[1])
			if num >= max {
				max = num + 1
			}
		}
	}
	path := fmt.Sprintf("test%d.csv", max)
	file, err := os.Create(path)
	checkErr(t, err)
	return file
}

package tld

import "testing"

type testpair struct {
	vaule    string
	expected string
}

var tests = []testpair{
	{"ABC@gmail.PL", "Poland"},
	{"fdslkj@fdjslk.fdh.GS", "South Georgia and the South Sandwich Islands"},
	{"ABC@.PL", ""},
	{"@gmail.COM", ""},
}

func TestLocator(t *testing.T) {
	fakeTLD := NewFakeTLD()

	for _, pair := range tests {
		country, _ := LocateEmail(pair.vaule, fakeTLD)
		if country != pair.expected {
			t.Errorf("For %s, Expected %s, got %s", pair.vaule, pair.expected, country)
		}
	}
}

type fakeTLD map[string]string

// NewFakeTLD mocks CountryTLD
func NewFakeTLD() fakeTLD {
	ret := make(map[string]string)
	ret["PL"] = "Poland"
	ret["GB"] = "United Kingdom"
	ret["GS"] = "South Georgia and the South Sandwich Islands"
	ret["COM"] = "Commercial"

	return ret
}

func (f fakeTLD) CheckTLD(tld string) string {
	return f[tld]
}

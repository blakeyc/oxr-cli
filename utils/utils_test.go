package utils

import "testing"

func TestExtractDates(t *testing.T) {

	t.Log("Testing extracting two dates from str... (expected array of length 2)")
	dates, err := ExtractDates("2017-01-01,2017-01-02")
	if err != nil {
		t.Errorf("Expected no errors to be thrown, one was: %s", err)
	}
	if len(dates) != 2 {
		t.Errorf("Expected array to be length 2 instead it was %d", len(dates))
	}

	t.Log("Testing passing no dates ... (expected empty response)")
	dates, err = ExtractDates("")
	if dates != nil {
		t.Errorf("Expected dates to equal nil: %s", dates)
	}
	if err == nil {
		t.Error("Expected error to be thrown")
	}

	t.Log("Testing passing bad date ... (expected empty response)")
	dates, err = ExtractDates("20160101")
	if dates != nil {
		t.Errorf("Expected dates to equal nil: %s", dates)
	}
	if err == nil {
		t.Error("Expected error to be thrown")
	}

	t.Log("Testing passing wrong delimiter ... (expected empty response)")
	dates, err = ExtractDates("2017-01-01|2017-01-02")
	if len(dates) != 0 {
		t.Errorf("Expected array to be length 0 instead it was %d", len(dates))
	}
	if err == nil {
		t.Error("Expected error to be thrown")
	}

	t.Log("Testing passing different date formats ... (expected empty response)")
	dates, err = ExtractDates("2017-01-01,20170102")
	if len(dates) != 0 {
		t.Errorf("Expected array to be length 0 instead it was %d", len(dates))
	}
	if err == nil {
		t.Error("Expected error to be thrown")
	}

}

func TestExtractFields(t *testing.T) {

	t.Log("Testing extracting two fields from str... (expected array of length 2)")
	fields, err := ExtractFields("base,currency")
	if err != nil {
		t.Errorf("Expected no errors to be thrown, one was: %s", err)
	}
	if len(fields) != 2 {
		t.Errorf("Expected array to be length 2 instead it was %d", len(fields))
	}

	t.Log("Testing passing no fields ... (expected empty response)")
	fields, err = ExtractFields("")
	if fields != nil {
		t.Errorf("Expected fields to equal nil: %s", fields)
	}
	if err == nil {
		t.Error("Expected error to be thrown")
	}

	t.Log("Testing passing unknown field... (expected empty response)")
	fields, err = ExtractFields("base,time")
	if fields != nil {
		t.Errorf("Expected fields to equal nil: %s", fields)
	}
	if err == nil {
		t.Error("Expected error to be thrown")
	}

}

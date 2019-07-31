package format

import (
	"testing"
)

type item struct {
	Name    string
	Age     int
	Married bool
}

func TestFormatTable(t *testing.T) {

	items := []interface{}{
		item{"hank", 32, true},
		item{"peggy", 31, true},
		item{"bobby", 1, false},
	}

	rows, err := Table(TableOpts{
		Rows:       items,
		Columns:    []string{"Name", "Age"},
		ShowHeader: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"===========",
		"NAME  | AGE",
		"===========",
		"hank  | 32 ",
		"peggy | 31 ",
		"bobby | 1  ",
	}

	for i, row := range rows {
		if row != expected[i] {
			t.Errorf("Got: '%s' Expected: '%s'", row, expected[i])
		}
	}
}

func TestFormatTableNoHeader(t *testing.T) {

	items := []interface{}{
		item{"a", 0, true},
		item{"abcd", 31, true},
		item{"abcdef", 1, false},
	}

	rows, err := Table(TableOpts{
		Rows:      items,
		Columns:   []string{"Name", "Married"},
		Separator: " . ",
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"a      . true ",
		"abcd   . true ",
		"abcdef . false",
	}

	for i, row := range rows {
		if row != expected[i] {
			t.Errorf("Got: '%s' Expected: '%s'", row, expected[i])
		}
	}
}

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Tag = string

type Expense struct {
	Date   time.Time
	Amount float64
	Class  string
	Tags   []Tag
}

func UTCDate(year int, month time.Month, day int) time.Time {
	timeZone, _ := time.LoadLocation("UTC")
	return time.Date(year, month, day, 0, 0, 0, 0, timeZone)
}

func displayExpenses(expenses []*Expense, wrap bool) {
	t := table.NewWriter()
	// if we don't wrap each header with [] they will show in reverse order
	if wrap {
		t.AppendHeader(table.Row{"#", "[תאריך]", "[סכום]", "[מחלקה]", "[תגים]"})
	} else {
		t.AppendHeader(table.Row{"#", "תגים", "מחלקה", "סכום", "תאריך"})
	}
	for i, e := range expenses {
		dateWithoutTime := strings.Split(e.Date.String(), " ")[0]
		// we format here to wrap with [] otherwise order will break
		eClass := e.Class
		if wrap {
			eClass = fmt.Sprintf("[%s]", e.Class)
		}
		t.AppendRows([]table.Row{
			{
				i,
				dateWithoutTime,
				e.Amount,
				eClass,
				e.Tags,
			},
		})
	}
	if wrap {
		t.AppendFooter(table.Row{"", "[סהכ]", 30})
	} else {
		t.AppendFooter(table.Row{"", "סהכ", 30})
	}
	t.SetCaption("Wrapped: %v", wrap)

	fmt.Printf("%s\n\n", t.Render())
}

func main() {
	testExpenses := []*Expense{
		{
			Date:   UTCDate(2021, 03, 18),
			Amount: 5.0,
			Class:  "מחלקה1",
			Tags:   []Tag{"תג1", "תג2"},
		},
		{
			Date:   UTCDate(2021, 04, 19),
			Amount: 5.0,
			Class:  "מחלקה1",
			Tags:   []Tag{"תג1"},
		},
		{
			Date:   UTCDate(2021, 05, 20),
			Amount: 5.0,
			Class:  "מחלקה2",
			Tags:   []Tag{"תג1"},
		},
	}

	displayExpenses(testExpenses, true)
	displayExpenses(testExpenses, false)
}

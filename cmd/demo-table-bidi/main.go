package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	wrapped = false
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

func displayExpenses(expenses []*Expense) {
	t := table.NewWriter()
	t.AppendHeader(generateHeader())
	for i, e := range expenses {
		dateWithoutTime := strings.Split(e.Date.String(), " ")[0]
		eClass := processBiDi(e.Class)
		eTags := processBiDi(strings.Join(e.Tags, " "))
		t.AppendRow(table.Row{i, dateWithoutTime, e.Amount, eClass, eTags})
	}
	t.AppendFooter(generateFooter())
	t.SetCaption("Wrapped: %v", wrapped)

	fmt.Printf("%s\n\n", t.Render())
}

func generateFooter() table.Row {
	row := table.Row{"", "סהכ", 30}
	row[1] = processBiDi(fmt.Sprint(row[1]))
	return row
}

func generateHeader() table.Row {
	row := table.Row{"#"}
	for _, col := range []string{"תאריך", "סכום", "מחלקה", "תגים"} {
		row = append(row, processBiDi(col))
	}
	return row
}

func processBiDi(str string) string {
	if wrapped {
		return fmt.Sprintf("[%s]", str)
	}
	return str
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

	wrapped = true
	displayExpenses(testExpenses)
	wrapped = false
	displayExpenses(testExpenses)
}

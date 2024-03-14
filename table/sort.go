package table

import (
	"sort"
	"strconv"
)

// SortBy defines What to sort (Column Name or Number), and How to sort (Mode).
type SortBy struct {
	// Name is the name of the Column as it appears in the first Header row.
	// If a Header is not provided, or the name is not found in the header, this
	// will not work.
	Name string
	// Number is the Column # from left. When specified, it overrides the Name
	// property. If you know the exact Column number, use this instead of Name.
	Number int

	// Mode tells the Writer how to Sort. Asc/Dsc/etc.
	Mode SortMode
}

// SortMode defines How to sort.
type SortMode int

const (
	// Asc sorts the column in Ascending order alphabetically.
	Asc SortMode = iota
	// AscNumeric sorts the column in Ascending order numerically.
	AscNumeric
	// Dsc sorts the column in Descending order alphabetically.
	Dsc
	// DscNumeric sorts the column in Descending order numerically.
	DscNumeric
)

type rowsSorter struct {
	rows          []rowStr
	sortBy        []SortBy
	sortedIndices []int
}

// getSortedRowIndices sorts and returns the row indices in Sorted order as
// directed by Table.sortBy which can be set using Table.SortBy(...)
func (t *Table) getSortedRowIndices() []int {
	sortedIndices := make([]int, len(t.rows))
	for idx := range t.rows {
		sortedIndices[idx] = idx
	}

	if t.sortBy != nil && len(t.sortBy) > 0 {
		sort.Sort(rowsSorter{
			rows:          t.rows,
			sortBy:        t.parseSortBy(t.sortBy),
			sortedIndices: sortedIndices,
		})
	}

	return sortedIndices
}

func (t *Table) parseSortBy(sortBy []SortBy) []SortBy {
	var resSortBy []SortBy
	for _, col := range sortBy {
		colNum := 0
		if col.Number > 0 && col.Number <= t.numColumns {
			colNum = col.Number
		} else if col.Name != "" && len(t.rowsHeader) > 0 {
			for idx, colName := range t.rowsHeader[0] {
				if col.Name == colName {
					colNum = idx + 1
					break
				}
			}
		}
		if colNum > 0 {
			resSortBy = append(resSortBy, SortBy{
				Name:   col.Name,
				Number: colNum,
				Mode:   col.Mode,
			})
		}
	}
	return resSortBy
}

func (rs rowsSorter) Len() int {
	return len(rs.rows)
}

func (rs rowsSorter) Swap(i, j int) {
	rs.sortedIndices[i], rs.sortedIndices[j] = rs.sortedIndices[j], rs.sortedIndices[i]
}

func (rs rowsSorter) Less(i, j int) bool {
	realI, realJ := rs.sortedIndices[i], rs.sortedIndices[j]
	for _, sortBy := range rs.sortBy {
		rowI, rowJ, colIdx := rs.rows[realI], rs.rows[realJ], sortBy.Number-1
		// extract the values/cells from the rows for comparison
		iVal, jVal := "", ""
		if colIdx < len(rowI) {
			iVal = rowI[colIdx]
		}
		if colIdx < len(rowJ) {
			jVal = rowJ[colIdx]
		}
		// compare and choose whether to continue
		shouldContinue, returnValue := rs.lessColumns(iVal, jVal, sortBy)
		if !shouldContinue {
			return returnValue
		}
	}
	return false
}

func (rs rowsSorter) lessColumns(iVal string, jVal string, sortBy SortBy) (bool, bool) {
	if iVal == jVal {
		return true, false
	} else if sortBy.Mode == Asc {
		return false, iVal < jVal
	} else if sortBy.Mode == Dsc {
		return false, iVal > jVal
	}

	iNumVal, iErr := strconv.ParseFloat(iVal, 64)
	jNumVal, jErr := strconv.ParseFloat(jVal, 64)
	if iErr == nil && jErr == nil {
		if sortBy.Mode == AscNumeric {
			return false, iNumVal < jNumVal
		} else if sortBy.Mode == DscNumeric {
			return false, jNumVal < iNumVal
		}
	}
	return true, false
}

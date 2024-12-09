package table

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPager(t *testing.T) {
	expectedOutput := `+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| PASSENGERID | SURVIVED | PCLASS | NAME                                                | SEX    | AGE | SIBSP | PARCH | TICKET           | FARE    | CABIN | EMBARKED |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| 1           | 0        | 3      | Braund, Mr. Owen Harris                             | male   | 22  | 1     | 0     | A/5 21171        | 7.25    |       | S        |
| 2           | 1        | 1      | Cumings, Mrs. John Bradley (Florence Briggs Thayer) | female | 38  | 1     | 0     | PC 17599         | 71.2833 | C85   | C        |
| 3           | 1        | 3      | Heikkinen, Miss. Laina                              | female | 26  | 0     | 0     | STON/O2. 3101282 | 7.925   |       | S        |
| 4           | 1        | 1      | Futrelle, Mrs. Jacques Heath (Lily May Peel)        | female | 35  | 1     | 0     | 113803           | 53.1    | C123  | S        |
| 5           | 0        | 3      | Allen, Mr. William Henry                            | male   | 35  | 0     | 0     | 373450           | 8.05    |       | S        |
| 6           | 0        | 3      | Moran, Mr. James                                    | male   |     | 0     | 0     | 330877           | 8.4583  |       | Q        |
| 7           | 0        | 1      | McCarthy, Mr. Timothy J                             | male   | 54  | 0     | 0     | 17463            | 51.8625 | E46   | S        |
| 8           | 0        | 3      | Palsson, Master. Gosta Leonard                      | male   | 2   | 3     | 1     | 349909           | 21.075  |       | S        |
| 9           | 1        | 3      | Johnson, Mrs. Oscar W (Elisabeth Vilhelmina Berg)   | female | 27  | 0     | 2     | 347742           | 11.1333 |       | S        |
| 10          | 1        | 2      | Nasser, Mrs. Nicholas (Adele Achem)                 | female | 14  | 1     | 0     | 237736           | 30.0708 |       | C        |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+`
	expectedOutputP1 := `+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| PASSENGERID | SURVIVED | PCLASS | NAME                                                | SEX    | AGE | SIBSP | PARCH | TICKET           | FARE    | CABIN | EMBARKED |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| 1           | 0        | 3      | Braund, Mr. Owen Harris                             | male   | 22  | 1     | 0     | A/5 21171        | 7.25    |       | S        |
| 2           | 1        | 1      | Cumings, Mrs. John Bradley (Florence Briggs Thayer) | female | 38  | 1     | 0     | PC 17599         | 71.2833 | C85   | C        |
| 3           | 1        | 3      | Heikkinen, Miss. Laina                              | female | 26  | 0     | 0     | STON/O2. 3101282 | 7.925   |       | S        |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+`
	expectedOutputP2 := `+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| PASSENGERID | SURVIVED | PCLASS | NAME                                                | SEX    | AGE | SIBSP | PARCH | TICKET           | FARE    | CABIN | EMBARKED |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| 4           | 1        | 1      | Futrelle, Mrs. Jacques Heath (Lily May Peel)        | female | 35  | 1     | 0     | 113803           | 53.1    | C123  | S        |
| 5           | 0        | 3      | Allen, Mr. William Henry                            | male   | 35  | 0     | 0     | 373450           | 8.05    |       | S        |
| 6           | 0        | 3      | Moran, Mr. James                                    | male   |     | 0     | 0     | 330877           | 8.4583  |       | Q        |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+`
	expectedOutputP3 := `+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| PASSENGERID | SURVIVED | PCLASS | NAME                                                | SEX    | AGE | SIBSP | PARCH | TICKET           | FARE    | CABIN | EMBARKED |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| 7           | 0        | 1      | McCarthy, Mr. Timothy J                             | male   | 54  | 0     | 0     | 17463            | 51.8625 | E46   | S        |
| 8           | 0        | 3      | Palsson, Master. Gosta Leonard                      | male   | 2   | 3     | 1     | 349909           | 21.075  |       | S        |
| 9           | 1        | 3      | Johnson, Mrs. Oscar W (Elisabeth Vilhelmina Berg)   | female | 27  | 0     | 2     | 347742           | 11.1333 |       | S        |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+`
	expectedOutputP4 := `+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| PASSENGERID | SURVIVED | PCLASS | NAME                                                | SEX    | AGE | SIBSP | PARCH | TICKET           | FARE    | CABIN | EMBARKED |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+
| 10          | 1        | 2      | Nasser, Mrs. Nicholas (Adele Achem)                 | female | 14  | 1     | 0     | 237736           | 30.0708 |       | C        |
+-------------+----------+--------+-----------------------------------------------------+--------+-----+-------+-------+------------------+---------+-------+----------+`

	tw := NewWriter()
	tw.AppendHeader(testTitanicHeader)
	tw.AppendRows(testTitanicRows)
	compareOutput(t, expectedOutput, tw.Render())

	p := tw.Pager(PageSize(3))
	assert.Equal(t, 1, p.Location())
	compareOutput(t, expectedOutputP1, p.Render())
	compareOutput(t, expectedOutputP2, p.Next())
	compareOutput(t, expectedOutputP2, p.Render())
	assert.Equal(t, 2, p.Location())
	compareOutput(t, expectedOutputP3, p.Next())
	compareOutput(t, expectedOutputP3, p.Render())
	assert.Equal(t, 3, p.Location())
	compareOutput(t, expectedOutputP4, p.Next())
	compareOutput(t, expectedOutputP4, p.Render())
	assert.Equal(t, 4, p.Location())
	compareOutput(t, expectedOutputP4, p.Next())
	compareOutput(t, expectedOutputP4, p.Render())
	assert.Equal(t, 4, p.Location())
	compareOutput(t, expectedOutputP3, p.Prev())
	compareOutput(t, expectedOutputP3, p.Render())
	assert.Equal(t, 3, p.Location())
	compareOutput(t, expectedOutputP2, p.Prev())
	compareOutput(t, expectedOutputP2, p.Render())
	assert.Equal(t, 2, p.Location())
	compareOutput(t, expectedOutputP1, p.Prev())
	compareOutput(t, expectedOutputP1, p.Render())
	assert.Equal(t, 1, p.Location())
	compareOutput(t, expectedOutputP1, p.Prev())
	compareOutput(t, expectedOutputP1, p.Render())
	assert.Equal(t, 1, p.Location())

	compareOutput(t, expectedOutputP1, p.GoTo(0))
	compareOutput(t, expectedOutputP1, p.Render())
	assert.Equal(t, 1, p.Location())
	compareOutput(t, expectedOutputP1, p.GoTo(1))
	compareOutput(t, expectedOutputP1, p.Render())
	assert.Equal(t, 1, p.Location())
	compareOutput(t, expectedOutputP2, p.GoTo(2))
	compareOutput(t, expectedOutputP2, p.Render())
	assert.Equal(t, 2, p.Location())
	compareOutput(t, expectedOutputP3, p.GoTo(3))
	compareOutput(t, expectedOutputP3, p.Render())
	assert.Equal(t, 3, p.Location())
	compareOutput(t, expectedOutputP4, p.GoTo(4))
	compareOutput(t, expectedOutputP4, p.Render())
	assert.Equal(t, 4, p.Location())
	compareOutput(t, expectedOutputP4, p.GoTo(5))
	compareOutput(t, expectedOutputP4, p.Render())
	assert.Equal(t, 4, p.Location())

	sb := strings.Builder{}
	p.SetOutputMirror(&sb)
	p.Render()
	compareOutput(t, expectedOutputP4, sb.String())
}

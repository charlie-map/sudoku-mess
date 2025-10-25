package practice

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestShouldBuild2x2Grid(t *testing.T) {
	expected :=
		"+---+---+\n" +
			"| 1 | 2 |\n" +
			"+---+---+\n" +
			"| 3 | 4 |\n" +
			"+---+---+\n"

	result := buildGridWithNumbers(2, 2, []int{1, 2, 3, 4})

	assert.Equal(t, result, expected)
}

func TestShouldBuild2x2GridWithEmpties(t *testing.T) {
	expected :=
		"+---+---+\n" +
			"| 1 |   |\n" +
			"+---+---+\n" +
			"|   | 4 |\n" +
			"+---+---+\n"

	result := buildGridWithNumbers(2, 2, []int{1, 0, 0, 4})

	assert.Equal(t, result, expected)
}

func TestShouldBuild3x3Grid(t *testing.T) {
	expected :=
		"+---+---+---+\n" +
			"| 1 | 2 | 3 |\n" +
			"+---+---+---+\n" +
			"| 4 | 5 | 6 |\n" +
			"+---+---+---+\n" +
			"| 7 | 8 | 9 |\n" +
			"+---+---+---+\n"

	result := buildGridWithNumbers(3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	assert.Equal(t, result, expected)
}

func TestShouldBuild5x5Grid(t *testing.T) {
	expected :=
		"+---+---+---+---+---+\n" +
			"| 1 | 2 | 3 | 4 | 5 |\n" +
			"+---+---+---+---+---+\n" +
			"| 6 | 7 | 8 | 9 |10 |\n" +
			"+---+---+---+---+---+\n" +
			"|11 |12 |13 |14 |15 |\n" +
			"+---+---+---+---+---+\n" +
			"|16 |17 |18 |19 |20 |\n" +
			"+---+---+---+---+---+\n" +
			"|21 |22 |23 |24 |25 |\n" +
			"+---+---+---+---+---+\n"

	result := buildGridWithNumbers(5, 5, []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25,
	})

	assert.Equal(t, result, expected)
}

package practice

import (
	"fmt"

	"sudoko-speed.charlie.city/utils"
)

func BuildMissingNumberTest(width int, height int) (ouput string, answer int) {
	length := width * height
	numbers := utils.RandomNumbersList(length)

	indexToClear := utils.GetRandomNumber(0, length-1)

	answer = numbers[indexToClear]
	numbers[indexToClear] = 0

	output := buildGridWithNumbers(width, height, numbers)

	return output, answer
}

func buildGridWithNumbers(width int, height int, numbers []int) string {
	output := ""

	for i := range height {
		output = output + buildSeperator(width) + buildRow(width, i, numbers)
	}

	output = output + buildSeperator(width)

	return output
}

func buildSeperator(width int) string {
	output := ""

	for range width {
		output = output + "+---"
	}

	return output + "+\n"
}

func buildRow(width int, rowNum int, numbers []int) string {
	output := ""

	for i := range width {
		pos := rowNum*width + i
		value := fmt.Sprintf("%d", numbers[pos])

		if numbers[pos] == 0 {
			value = " "
		}

		endValue := fmt.Sprintf("| %s ", value)
		switch len(value) {
		case 2:
			endValue = fmt.Sprintf("|%s ", value)
		case 3:
			endValue = fmt.Sprintf("|%s", value)
		}

		output = output + endValue
	}

	return output + "|\n"
}

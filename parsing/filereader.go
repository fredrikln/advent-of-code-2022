package parsing

import (
	"os"
	"strconv"
	"strings"
)

func readfile(filename string) string {
	content, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return strings.Trim(string(content), "\n")
}

func ReadRowsToStringSlice(filename string) []string {
	return strings.Split(readfile(filename), "\n")
}

func stringSliceToIntSlice(input []string) []int {
	out := make([]int, 0, len(input))

	for _, value := range input {
		number, err := strconv.Atoi(value)

		if err != nil {
			panic(err)
		}

		out = append(out, number)
	}

	return out
}

func ReadRowsToIntSlice(filename string) []int {
	lines := ReadRowsToStringSlice(filename)
	return stringSliceToIntSlice(lines)
}

func ReadFirstLineToIntSlice(filename string) []int {
	line := readfile(filename)

	values := strings.Split(line, " ")

	return stringSliceToIntSlice(values)
}

func ReadFirstLineToString(filename string) string {
	line := readfile(filename)

	return line
}

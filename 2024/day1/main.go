package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mapSlice(slice []string, fn func(string) (int, error)) ([]int, error) {
	result := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		num, err := fn(slice[i])
        if err != nil {
            return nil, err
        }
		result[i] = num
	}
	return result, nil
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)
	inputRows := strings.Split(string(data), "\n")
    rowsLength := len(inputRows) - 1
	listOne := make([]int, rowsLength)
	listTwo := make([]int, rowsLength)
	for i := 0; i < rowsLength; i++ {
		row := inputRows[i]
		stringNums := strings.Fields(row)
        intNums, err := mapSlice(stringNums, strconv.Atoi)
        check(err)

        listOne[i] = intNums[0]
        listTwo[i] = intNums[1]
	}
    sort.Ints(listOne)
    sort.Ints(listTwo)

    totalDistance := 0
    for i, num := range listOne {
        distance := num - listTwo[i]
        if distance < 0 {
            distance = -distance
        }
        totalDistance += distance
    }
    fmt.Println("Answer:", totalDistance)
}

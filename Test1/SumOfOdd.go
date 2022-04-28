package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var lines [2]string
	var firstIdxLine = 0
	var secondIdxLine = 1

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	for i := 0; i < 2; i++ {
		line := readLine(reader)
		lines[i] = line
	}

	numsSize, err := strconv.Atoi(lines[firstIdxLine])
	checkError(err)
	fmt.Println(numsSize)

	numsInt := parseSlice(lines[secondIdxLine])
	sort.Ints(numsInt[:])

	numsStr := make([]string, len(numsInt))
	for i, s := range numsInt {
		numsStr[i] = strconv.Itoa(s)
	}

	fmt.Println(strings.Join(numsStr, " "))

	if len(numsInt) != numsSize {
		panic("Invalid input")
	}

	fmt.Println(sumOdd(numsInt, numsSize))
}

func readLine(reader *bufio.Reader) string {
	strLine, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	str := strings.TrimRight(string(strLine), "\r\n")

	return str
}

func parseSlice(input string) []int {
	strSlice := strings.Split(input, " ")
	intSlice := make([]int, len(strSlice))
	for i := range intSlice {
		intSlice[i], _ = strconv.Atoi(strSlice[i])
	}

	return intSlice
}

func sumOdd(nums []int, numsSize int) int {
	sum := 0
	for i := 0; i < numsSize; i += 2 {
		sum += nums[i]
	}

	return sum
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

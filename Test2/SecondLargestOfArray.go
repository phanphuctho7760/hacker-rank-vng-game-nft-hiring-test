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

	fmt.Println(getSecondLargestItem(numsInt))
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

func getSecondLargestItem(nums []int) int {
	length := len(nums)
	if length == 0 {
		panic("Slice nums cannot be 0-size.")
	}

	if length == 1 {
		return nums[0]
	}

	return nums[length-2]
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

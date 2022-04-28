package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/big"
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

	primeSlice := getSlicePrime(numsInt, numsSize)

	fmt.Println(lastDigitOfSlice(primeSlice))
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

func checkPrimeNumber(num int) bool {
	if big.NewInt(int64(num)).ProbablyPrime(0) {
		return true
	} else {
		return false
	}
}

func getSmallestItem(nums []int) int {
	length := len(nums)
	if length == 0 {
		panic("Slice nums cannot be 0-size.")
	}

	return nums[0]
}

func getLargestItem(nums []int) int {
	length := len(nums)
	if length == 0 {
		panic("Slice nums cannot be 0-size.")
	}

	if length == 1 {
		return nums[0]
	}

	return nums[length-1]
}

func getSlicePrime(nums []int, numsSize int) []int {
	var resSlice []int
	for i := 0; i < numsSize; i++ {
		if checkPrimeNumber(nums[i]) {
			resSlice = append(resSlice, nums[i])
		}
	}

	return resSlice
}

func lastDigitOfSlice(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	smallestNum := getSmallestItem(nums)
	largestNum := getLargestItem(nums)

	fmt.Println(nums)
	fmt.Println(smallestNum)
	fmt.Println(largestNum)

	return int(math.Pow(float64(smallestNum), float64(largestNum)))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

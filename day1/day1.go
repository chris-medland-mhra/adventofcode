package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readFileToArray(file string) []string {
	fileContent, _ := os.Open(file)
	scanner := bufio.NewScanner(fileContent)
	var arr []string
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}

func stringToInt(input string) string {
	replacer := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e",
	)
	return replacer.Replace(input)
}

func main() {
	inputArray := readFileToArray("testinput.text")
	var total int
	for _, value := range inputArray {
		value = stringToInt(stringToInt(value))
		fmt.Println("value", value)
		var numberString string
		for _, char := range value {
			if unicode.IsNumber((char)) {
				numberString = numberString + string(char)
			}
		}
		fmt.Println("numberString", numberString)
		numberString = string(numberString[0]) + string(numberString[len(numberString)-1])
		number, _ := strconv.Atoi(numberString)
		total = total + number
	}
	fmt.Println("TOTAL", total)
	//fmt.Println(stringToInt("eighttwothree"))
}

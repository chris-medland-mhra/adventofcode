package main

import (
    "fmt"
    "os"
    "bufio"
    "slices"
    "strings"
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

func returnArrs (input []string) int {
    total := 0 
    for _, v := range input {
        count := 0
        lineArr := strings.Fields(v)
        splitIndex := slices.Index(lineArr, "|")
        winners := lineArr[0:splitIndex]
        checkers := lineArr[splitIndex+1:len(lineArr)]
        for _, v := range winners {
            if slices.Contains(checkers, v){
                if count == 0 {
                    count = 1
                } else {
                count = count << 1
                }
            }
        }
        total = total + count
    }
    return total 
}

func main () {
	inputArray := readFileToArray("input.txt")
    fmt.Println(returnArrs(inputArray))
}


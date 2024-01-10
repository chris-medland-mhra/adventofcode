
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

func returnCount (input []string, start, end int) int {

    total := end - start  
    for i := start; i < end; i++ { 
        wins := numWinners(input[i])
        if wins > 0 {
            total = total + returnCount(input, i+1, i+wins+1)
        }
    }
    return total 
}

func numWinners (line string) int {
    count := 0
    lineArr := strings.Fields(line)
    splitIndex := slices.Index(lineArr, "|")
    winners := lineArr[0:splitIndex]
    checkers := lineArr[splitIndex+1:len(lineArr)]
    for _, v := range winners {
        if slices.Contains(checkers, v){
            count = count + 1
        }
    }
    return count

}

func main () {
	inputArray := readFileToArray("input.txt")
    fmt.Println(returnCount(inputArray, 0, len(inputArray)))
}


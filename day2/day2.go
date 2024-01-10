package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "strconv"
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

func lineToValueArr(line string) []string {
    r := strings.Fields(line)
    return r
}

func calculateGame(game []string) int {

    resultMap := make(map[string]int)
    for i := 0; i < len(game); i = i + 2 {
        v, _ := strconv.Atoi(game[i])
        _, ok := resultMap[game[i+1]]
        if ok {
            if v > resultMap[game[i+1]] {
            resultMap[game[i+1]] = v 
            }
        } else {
            resultMap[game[i+1]] = v 
        }
    }
    fmt.Println(resultMap)
    return resultMap["red"] * resultMap["green"] * resultMap["blue"]
}

func calculateGameLine(game []string) bool {

    result := true
    for i:=0; i<len(game); i=i+2 {
        v, _ := strconv.Atoi(game[i])
        switch color := game[i+1]; color {
        case "red" :
            if v > 12 {
                result = false
            }
        case "blue" :
            if v > 14 {
                result = false
            }
        case "green" :
            if v > 13 {
                result = false
            }
        }
    }
    return result
} 

func main() {
	inputArray := readFileToArray("input.txt")

    count := 0
    for _, line := range inputArray {
        fmt.Println(lineToValueArr(line))
        fmt.Println(calculateGame(lineToValueArr(line)))
        count = count + calculateGame(lineToValueArr(line))
    }
    fmt.Println(count)

}

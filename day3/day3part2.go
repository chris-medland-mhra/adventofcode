package main 

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "slices"
)

type Coord struct {
    X           int
    Y           int
}

type Numloc struct {
    value       int
    neighbours  []Coord 
}

func readFileToArray(file string) []string {
    fileContent, _ := os.Open(file)
    scanner := bufio.NewScanner(fileContent)
    var arr []string
    for scanner.Scan() {
        arr = append(arr, scanner.Text())
    }
    return arr
}

func getNumbers(line string) []string {
    s := strings.Split(line, "")
    return s
}

func toMatrix(input []string) [][]string {
    matrix := make([][]string, len(input))
    for i, v := range input {
        matrix[i] = getNumbers(v) 
    } 
    return matrix
}

func inBounds(x, y, length, height int) bool {
    return x >= 0 && x < height && y >= 0 && y < length
}

func getNeighbours(numsize, length, height, x, y int) []Coord  {
    neighbours := make([]Coord, 0)
    for i := 0; i < numsize; i++ {
        for dy := -1; dy <= 1; dy++ {
            for dx := -1; dx <= 1; dx++ {
                if dy == 0 && dx == 0 {
                    continue
                }
                if i > 0 && i < numsize -1 {
                    if dy == -1 || dy == 1 {
                        continue
                    }
                }
                c := Coord{x+dx, y+dy}
                if inBounds(c.X, c.Y, length, height) {
                    neighbours = append(neighbours, c)
                }
            }
        }
        y = y + 1
    }
    return neighbours
}

func getNumbersLoc(input [][]string) []Numloc{
    var numarr []Numloc
    for i:=0;i<len(input);i++ {
        for j:=0;j<len(input[i]);j++ {
            if _, err := strconv.Atoi(input[i][j]); err == nil {
                b := true
                s := input[i][j]
                count := 1
                startpoint := Coord{i, j}
                for b {
                    if (j + 1) < len(input[i]) {
                        if _, err := strconv.Atoi(input[i][j+1]); err == nil {
                            j = j + 1
                            s = s + input[i][j]
                            count = count + 1
                        } else {
                            b = false
                        }

                    } else {
                        b = false
                    }
                }
                n, _ := strconv.Atoi(s)
                pointWithNeighbours := Numloc{n, getNeighbours(count, len(input[i]), len(input), startpoint.X, startpoint.Y)}
                numarr = append(numarr, pointWithNeighbours)
            }
        }
    }
    return numarr
}

func checkNeighbours(neighbours []Coord, matrix [][]string) bool {
    for _, v := range neighbours {
        if _, err := strconv.Atoi(matrix[v.X][v.Y]); err != nil {
            if matrix[v.X][v.Y] != "." {
                return true
            }
        }
    }
    return false
}

func getGearCoords(input [][]string) []Coord { 
    gearCoords := make([]Coord, 0)
    for i:=0;i<len(input);i++ {
        for j:=0;j<len(input[i]);j++ {
            if input[i][j] == "*" {
                g := Coord{i , j}
                gearCoords = append(gearCoords, g)
            }
        }
    }
    return gearCoords
}

func getGearRatios(gearCoords []Coord, numValCoords []Numloc) int {
    total := 0 
    for _, gv := range gearCoords {
        count := make([]int, 0)
        for _, nv := range numValCoords {
            if slices.Contains(nv.neighbours, gv) {
                count = append(count, nv.value)
            }
        }
        if len(count) == 2 {
            total = total + (count[0] * count[1])
        }
    }
    return total 
}

func main () {
    inputArray := readFileToArray("realinput.txt")
    matrix := toMatrix(inputArray)
    valuesWithCoords := getNumbersLoc(matrix)
    fmt.Println(getGearRatios(getGearCoords(matrix), valuesWithCoords))
}


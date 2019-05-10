package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    a := ReadInput()
    fmt.Println("Here is the result:", AddAll(a))
}

func ReadInput() []float64 {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter digits to be added (space separated): ")
    text, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }

    textSlice := strings.Fields(text)
    floatsSlice := make([]float64, 0)
    for _, elem := range textSlice {
        i, err := strconv.ParseFloat(elem, 64)
        if err != nil {
            fmt.Println(err)
        }
        floatsSlice = append(floatsSlice, i)
    }

    return floatsSlice
}

func AddAll(floatsSlice []float64) float64 {
    var sum float64
    for _, v := range floatsSlice {
        sum += v
    }
    return sum
}
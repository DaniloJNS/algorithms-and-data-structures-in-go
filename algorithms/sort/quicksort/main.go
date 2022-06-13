package main

import (
	"fmt"
	"os"
	"strconv"
)

func particionar(numbers []int, pivot int) (minors []int, bigger []int) {
    for _, e := range numbers {
        if pivot > e {
            bigger = append(bigger, e)
       } else {
           minors = append(minors, e)
       }
    }
    return minors, bigger
}

func quicksort(numbers []int) []int {
    if len(numbers) < 1 {
        return numbers
    }

    n := make([]int , len(numbers))
    copy(n, numbers)

    pivotIndex := len(numbers) / 2
    pivot := n[pivotIndex]

    n = append(numbers[:pivotIndex], numbers[pivotIndex+1:]...)

    minors, bigger := particionar(n, pivot)
    
    return append(append(quicksort(minors), pivot), quicksort(bigger)...)
}

func main() {
    input := os.Args[1:]
    
    numbers := make([]int, len(input))

    for i, in:= range input {
        number, err := strconv.Atoi(in)  

        if err != nil {
            fmt.Printf("%s is not a number\n", in)
            panic(err)
            
        }

        numbers[i] = number
    }

    fmt.Println(numbers)
}

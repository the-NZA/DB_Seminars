package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/go-pkgz/lgr"
)

func main() {
	arr1 := []int{71, 10, 20, 30, 0, 55, 63, 4, 5, 9, 10, 12, 14, 18, 21, 22, 23, 24, 25, 26, 29, 30, 36, 39, 42, 43, 45, 46, 47, 48,  52, 56, 57, 59, 61, 63, 64, 65, 68, 70, 72, 77, 80, 82, 83, 85, 86, 88, 91, 94, 96, 98, 99, 100}
	arr2 := []int{50, 40, 32, 7, 13, 58, 129, 112}

	res, err := solve(arr1, arr2)
	if err != nil {
		lgr.Fatalf("Error during solving problem: %v\n", err)
	}
	fmt.Printf("Min sub is: %d\n", res)

        arr1 = []int{10, 16, 28, 2}
        arr2 = []int{20, 56, 74 }
	res, err = solve(arr1, arr2)
	if err != nil {
		lgr.Fatalf("Error during solving problem: %v\n", err)
	}
	fmt.Printf("Min sub is: %d\n", res)
        
        arr1 = []int{80, 5, 31, 6, 17, 71}
        arr2 = []int{96, 19,1 }
	res, err = solve(arr1, arr2)
	if err != nil {
		lgr.Fatalf("Error during solving problem: %v\n", err)
	}
	fmt.Printf("Min sub is: %d\n", res)
}

func solve(arr1, arr2 []int) (int, error) {
	arr1Len := len(arr1)
	arr2Len := len(arr2)
	if arr1Len == 0 || arr2Len == 0 {
		return 0, fmt.Errorf("Both slices must have at least 1 item")
	}

        sort.Slice(arr1, func(i,j int) bool{
                return arr1[i] < arr1[j]
        })

        sort.Slice(arr2, func(i,j int) bool{
                return arr2[i] < arr2[j]
        })

        mergedLen := arr1Len + arr2Len
        i,j := 0,0

        min := int(math.Abs(float64(arr1[0]) - float64(arr2[0])))

        for k := 0; k <  mergedLen; k += 1 {
                sub := 0
                if i > arr1Len - 1 {
                        sub = int(math.Abs(float64(arr1[i - 1]) - float64(arr2[j])))
                        j++
                } else if j > arr2Len - 1 {
                        sub = int(math.Abs(float64(arr1[i ]) - float64(arr2[j - 1])))
                        i++
                } else if arr1[i] < arr2[j] {
                        sub = int(math.Abs(float64(arr1[i]) - float64(arr2[j])))
                        i++
                } else {
                        sub = int(math.Abs(float64(arr1[i]) - float64(arr2[j])))
                        j++
                }

                if sub < min {
                        min = sub
                }
        }



        return min, nil
}

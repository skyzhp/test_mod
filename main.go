package main

import (
	"fmt"
	"github.com/skyzhp/util"
	a "github.com/skyzhp/util/a"
	"time"
)

func main() {
	fmt.Println(util.Add(1, 1))
	a.LogTest("111")

	arr := []int{}
	arr = append(arr, 1)
	fmt.Printf("%p\n", arr)
	testArray(arr)
	fmt.Println(arr)

	arr2 := make([]int, 100)

	for {
		for i := 0; i < 100; i++ {
			//index := i
			go func() {
				arr2[0] = 3
			}()
		}
		time.Sleep(time.Second)
	}

}
func testArray(arr []int) {
	fmt.Println(arr)
	fmt.Printf("%p\n", arr)
	arr[0] = 2
}

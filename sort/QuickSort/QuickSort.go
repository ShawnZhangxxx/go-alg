package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	arr := []int{1,3,4,6,7,12,18,10}
	//fmt.Print(rand.Intn(len(arr)))

	//QuickSort(arr)
	//fmt.Println(Partion(arr,5,0,7))
	//fmt.Println(arr)
	QuickSort(arr,0,len(arr) -1)
	fmt.Println(arr)

}

func QuickSort(arr []int,start int,end int) {
	if start >= end || start < 0 || end > len(arr) - 1  {
		return
	}
	p := getRand(start,end)

	p = Partion(arr,p,start,end)
	QuickSort(arr,start,p-1)
	QuickSort(arr,p+1,end)


}

func getRand(start int,end int) int {
	rand.Seed(time.Now().UnixNano())
	if end == start {
		return start
	}
	return rand.Intn(end - start) + start
}

/**
 *
 */
func Partion(arr []int,p int,start int,end int) int  {
	pivot := arr[p]
	for start < end {
		for arr[start] < pivot {
			start ++
		}
		for arr[end] > pivot {
			end --
		}
		if start < end {
			arr[start],arr[end] = arr[end],arr[start]
		}
	}
	return start
}

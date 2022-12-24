package main

import "fmt"

func main()  {
	arr := []int{1,3,4,6,7,2,18,10}
	SelectSort(arr)
	fmt.Println(arr)
}

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i],arr[j] = arr[j],arr[i]
			}
		}


	}

}

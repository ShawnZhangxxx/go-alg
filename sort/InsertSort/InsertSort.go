package main

import "fmt"

func main()  {
	arr := []int{11,1,3,4,6,7,2,18,10}
	InsertSort(arr)
	fmt.Println(arr)
}

func InsertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j],arr[j-1] = arr[j-1],arr[j]
			}else {
				break
			}
		}
	}
}

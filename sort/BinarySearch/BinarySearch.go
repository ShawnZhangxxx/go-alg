package main

import "fmt"

func BinarySearch(arr []int,v int) int  {
	start := 0
	end := len(arr) - 1

	mid := (start + end) / 2

	for i := 0; i < len(arr); i++ {
		if start == end {
			return -2
		}
		if arr[mid] == v {
			return mid
		}else if arr[mid] > v {
			end = mid -1
		}else {
			start = mid + 1
		}
		mid = (start + end) / 2

	}

	return -1
}


func BinarySearchRecursion(arr []int,start int, end int ,v int) int  {

	if start > end {
		return  -1
	}

	mid := (start + end) / 2

	if arr[mid] == v {
		return mid
	}else if arr[mid] > v {
	  return 	BinarySearchRecursion(arr,start,mid - 1,v)
	}else {
	  return 	BinarySearchRecursion(arr,mid + 1,end,v)
	}

}

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	//index := BinarySearch(arr,0)
	index := BinarySearchRecursion(arr,0,len(arr) - 1,11)
	fmt.Println(index)
}
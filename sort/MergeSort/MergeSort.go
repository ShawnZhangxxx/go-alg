package main

import "fmt"

func main()  {
	arr := []int{1,3,4,666,7,2,18,10}
	arr = MergeSort(arr)
	fmt.Println(arr)
}

func MergeSort(nums []int)[]int {
		if len(nums) < 2  {
			return nums
		}
		i := len(nums) / 2
		left := MergeSort(nums[0:i])
		right := MergeSort(nums[i:])

		result := Merge(left,right)
		return result
	}


//func MergeSort(arr []int,start int,end int)[] int {
//	if start<0 || end >len(arr) - 1  {
//		return nil
//	}
//	if start == end{
//		return []int{arr[start]}
//	}
//
//	m := (start + end) /2
//	left := MergeSort(arr,start,m)
//	right := MergeSort(arr,m+1,end)
//	result := Merge(left,right)
//	return result
//}
func Merge(left, right []int)[]int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i<l && j<r{
		if left[i] > right[j]{
			result = append(result, right[j])
			j++
		}else {
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}

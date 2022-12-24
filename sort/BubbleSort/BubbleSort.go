package main

import "fmt"

func main()  {
	arr := []int{1,3,4,6,7,2,18,10}
    BubbleSort(arr)
	fmt.Println(arr)
}

func BubbleSort(arr []int) {
	flag := true
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) -i -1; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
				flag = false
			}
		}
		//判断数据是否是有序
		if flag {
			return
		} else {
			flag = true
		}
	}

}

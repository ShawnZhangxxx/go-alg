package main

import "fmt"

//左孩子:2*i + 1  右孩子2*i+2
// 父节点 i - 1 /2

func main() {
	arr := []int{0, 1, 2, 4, 5, 6, 3,88,99,11,1,3,5,1,2,77}
	//fmt.Println(-1/2)
	for i := 0; i < len(arr); i++ { //建堆 o(n) 从堆顶开始
		Heap(arr, i, len(arr))
	}
	fmt.Println(arr)

	l := len(arr) - 1
	for i := l; i >= 0; i-- {//排序弹出
		Heapify(arr, i)
	}
}

func HeapSort(arr []int, start int, end int) {

}

func Heap(arr []int, end int, heapSize int) { //建堆过程
	//if end < heapSize {
		for arr[end] > arr[(end-1)/2] && end > 0 { //与父节点比较
			arr[end], arr[(end-1)/2] = arr[(end-1)/2], arr[end]
			end = (end - 1) / 2
		}
	//}
}

func Heapify(arr []int, end int) {
	fmt.Println(arr[0])
	arr[0] = arr[end]
	i := 0
	for (2*i + 1) < end {
		var max int
		if (2*i+2 < end) && (arr[2*i+2] > arr[2*i+1]) {
			max = 2*i + 2
		} else {
			max = 2*i + 1
		}
		if arr[max] > arr[i] {
			arr[max], arr[i] = arr[i], arr[max]
			i = max
		} else {
			break
		}

	}
	//return arr[:end]
}

package main

func ShellSort(arr []int) {  //1.第一次分组,间隔4个进行插入排序2.第二次间隔2 ,最后间隔1进行插入排序     排序次数和交换次数变少
	//根据增量（len(arr)/2）每次变成上一次的1/2
	for inc := len(arr) / 2; inc > 0; inc /= 2 {

		for i := inc; i < len(arr); i++ {  //不这么写会写出四次循环
			temp := arr[i]

			//根据增量从数据到0进行比较
			for j := i - inc; j >= 0; j -= inc {
				//满足条件交换数据
				if temp < arr[j] {
					arr[j], arr[j+inc] = arr[j+inc], arr[j]
				} else {
					break
				}
			}
		}
	}
}

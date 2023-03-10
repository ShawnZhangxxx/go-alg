package main

import (
	"fmt"
	"time"
)

func test(ch chan int, quit chan bool) {

}

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	go func() { // 子go 程获取数据
		for {
			fmt.Println("sssssssssssss ")
			select {   //如果没有输入channel就会阻塞住
			case num := <-ch1:
				fmt.Println("num1 = ", num)
			case <-ch2:
				fmt.Println("num2= ", ch2)
			//default:
			//	fmt.Println("dddddddddddd")

				// return
				// runtime.Goexit()
			}
			fmt.Println("eeeeeeeeeeee ")
		}
	}()
	ch1 <- 1
	ch1 <- 3
	ch2 <- 2
	time.Sleep(time.Second * 2)
	fmt.Println("finish!!!")
}

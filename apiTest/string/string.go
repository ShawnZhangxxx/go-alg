package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 101
	var y string
	var z string
	y = string(x)  //打印 e
	fmt.Println(y)
	z = strconv.Itoa(x)  //数字转字符串
	fmt.Println(z)
}


package main

import (
  "fmt"
  "time"
)

// 指针
// 1.题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
func add(num *int) {
	*num += 10
}

func main() {
	number := 2
	add(&number)
	fmt.Println("修改后的值：", number)
}

// 2.实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func multiply(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	multiply(&numbers)
	fmt.Println(numbers)
}

// Goroutine
// 1.编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
func odds() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数:", i)
	}
}

func evens() {
	for j := 2; j <= 10; j += 2 {
		fmt.Println("偶数:", j)
	}
}
func main() {
	go odds()
	go evens()
	time.Sleep(100 * time.Millisecond)
}

package main

import (
  "fmt"
  "time"
  "math"
  "sync/automic"
)

// 指针
// 1.编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
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

// 2.设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
func main() {
	// 任务1：计算1到100000的总和
	task1 := func() {
		sum := 0
		for i := 1; i <= 100000; i++ {
			sum += i
		}
		fmt.Println("任务1结果：", sum)
	}
	// 任务2：计算1到50000的平方和
	task2 := func() {
		sum := 0
		for i := 1; i <= 50000; i++ {
			sum += i * i
		}
		fmt.Println("任务2结果：", sum)
	}
	ch := make(chan bool, 2)

	// 启动任务1
	go func() {
		start := time.Now()
		task1()
		fmt.Println("任务1耗时：", time.Since(start))
		ch <- true
	}()

	// 启动任务2
	go func() {
		start := time.Now()
		task2()
		fmt.Println("任务2耗时：", time.Since(start))
		ch <- true
	}()
	<-ch
	<-ch
}

// 面向对象
// 1.定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
type Shape interface {
	Area()       float64
	Perimeter()  float64
}

type Rectangle struct {
	width  float64
	height float64
}
func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rect := Rectangle{
		width:  5.5,
		height: 3.2,
	}

	circle := Circle{
		radius: 4.6,
	}

	fmt.Println("矩形面积：", rect.Area())
	fmt.Println("矩形周长：", rect.Perimeter())

	fmt.Println("圆形面积：", circle.Area())
	fmt.Println("圆形周长：", circle.Perimeter())
}   

// 2.使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	Person
}

func (e Employee) PrintInfo() {
	fmt.Printf("ID: %d, 姓名: %s, 年龄: %d\n", e.EmployeeID, e.Name, e.Age)
}

func main() {
	info := Employee{
		Person: Person{
			Name: "张三",
			Age:  18,
		},
		EmployeeID: 101,
	}

	info.PrintInfo()
}

// Channel
// 1.编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for num := range ch {
			fmt.Printf("%d ", num)
		}
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("\n程序结束")
}


// 2.实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
func main() {
	ch := make(chan int, 5)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			select {
			case num, ok := <-ch:
				if !ok {
					return
				}
				fmt.Printf("%d ", num)
			}
		}
	}()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\n数据处理完成")
}

// 锁机制
// 1.编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	counter := SafeCounter{}
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("Final count: %d\n", counter.GetCount())
}

// 2.使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func main() {
	var count int64

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1) 
			}
		}()
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Final count: %d\n", atomic.LoadInt64(&count))
}

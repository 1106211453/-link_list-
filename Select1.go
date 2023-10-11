package main

import (
	"fmt"
	"time"
)

func SL() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(10 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
func SL2() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine，分别从两个通道中获取数据
	go func() {
		for {
			ch1 <- "from 1"
		}
	}()
	go func() {
		for {
			ch2 <- "from 2"
		}
	}()

	// 使用 select 语句非阻塞地从两个通道中获取数据
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			// 如果两个通道都没有可用的数据，则执行这里的语句
			fmt.Println("no message received")
		}
	}
}
func SL3(){
	 var numbers []int
	    printSlice(numbers)

		   /* 允许追加空切片 */
		      numbers = append(numbers, 0)
			     printSlice(numbers)

				    /* 向切片添加一个元素 */
					   numbers = append(numbers, 1)
					      printSlice(numbers)

						     /* 同时添加多个元素 */
							    numbers = append(numbers, 2,3,4)
								   printSlice(numbers)

								      /* 创建切片 numbers1 是之前切片的两倍容量*/
									     numbers1 := make([]int, len(numbers), (cap(numbers))*2)

										    /* 拷贝 numbers 的内容到 numbers1 */
											   copy(numbers1,numbers)
											      printSlice(numbers1)  
												  }

												  func printSlice(x []int){
												     fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)										
}


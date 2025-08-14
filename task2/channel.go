package main

import (
	"fmt"
	"time"
)

/*
*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func main1() {
	channel1()
	channel2()
	time.Sleep(3 * time.Second)
	return
}

func channel1() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println(v)
			} else {
				ch = nil
			}
		default:
			if ch == nil {
				fmt.Println("channel is nil!")
			}
		}
		if ch == nil {
			break
		}
	}
}

func channel2() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			select {
			case v, ok := <-ch:
				if ok {
					fmt.Println(v)
				} else {
					ch = nil
				}
			default:
				if ch == nil {
					fmt.Println("channel is nil!")
				}
			}
			if ch == nil {
				break
			}
		}
	}()
}

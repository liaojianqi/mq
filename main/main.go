package main

import (
	"fmt"
	"broker"
	"strconv"
)

func main() {
	topic := "sum10"
	b := broker.NewBroker()
	for i := 1; i <= 100; i++ {
		b.Produce(topic, []byte(fmt.Sprintf("%d", i)))
	}
	finish := make(chan int64, 10)
	for i := 0; i < 10; i ++ {
		go func(i int) {
			sum := int64(0)
			for j := 0; j < 10; j++ {
				buf := b.Consume(topic, int64(i * 10 + j))
				t, _ := strconv.ParseInt(string(buf), 10, 64)
				sum += t
			}
			finish <- sum
		}(i)
	}
	sum := int64(0)
	for i := 0; i < 10; i++ {
		t := <- finish
		sum += t
	}
	fmt.Println(sum)
}
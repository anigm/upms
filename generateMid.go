package main

import (
	"fmt"

	"time"

	"github.com/zheng-ji/goSnowFlake"
)

func main() {
	// Params: Given the workerId, 0 < workerId < 1024
	iw, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 100; i++ {
		if id, err := iw.NextId(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id)
		}
	}
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		if id, err := iw.NextId(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(i, id)
		}
	}
	fmt.Println(goSnowFlake.ParseId(103574008141713412))
}

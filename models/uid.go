package models

import (
	"fmt"
)

func GetOne() int64 {
	if id, err := Worker.NextId(); err != nil {
		fmt.Println(err)
		return -1
	} else {
		return id
	}
}

func GetMany(count int) []int64 {
	var ids []int64 = make([]int64, count)
	for i := 0; i < count; i++ {
		if id, err := Worker.NextId(); err != nil {
			fmt.Println(err)
		} else {
			ids[i] = id
		}
	}
	return ids
}

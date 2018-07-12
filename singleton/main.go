package main

import (
	"fmt"
	"sync"
)

// 全局实例
type singleton struct {
	test string
}

var once sync.Once
var instance *singleton

// 获取实例对象
func getInstace() *singleton {
	once.Do(func() {
		instance = &singleton{""}
	})
	return instance
}

func main() {
	i1 := getInstace()
	fmt.Println(*i1)
	i1.test = "1"
	i2 := getInstace()
	fmt.Println(*i2)
}

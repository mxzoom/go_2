package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandDiv(val int) int { //функция возвращает результат деления двух случайных чисел в диапазоне от 0 до val
	defer func() {
		if rcvr := recover(); rcvr != nil {
			fmt.Printf("recovered, %v \n", rcvr)
		}
	}()
	rand.Seed(time.Now().UnixMicro())
	randInt1 := rand.Intn(val)
	randInt2 := rand.Intn(val)
	time.Sleep(50 * time.Millisecond)
	return randInt1 / randInt2

}
func main() {

	for i := 100; i != 0; i-- {
		fmt.Println(RandDiv(20))
	}

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type myErr struct {
	time time.Time
	err  string
}

func (myerr *myErr) error() string {
	return fmt.Sprintf("error - %s, was discovered in %s", myerr.err, myerr.time)
}

func newErr(error string) myErr {
	return myErr{
		time: time.Now(),
		err:  error,
	}
}

func RandDiv(val int) int { //функция возвращает результат деления двух случайных чисел в диапазоне от 0 до val

	rand.Seed(time.Now().UnixMicro())
	randInt1 := rand.Intn(val)
	randInt2 := rand.Intn(val)
	time.Sleep(50 * time.Millisecond)
	return randInt1 / randInt2

}
func main() {
	defer func() {
		if rcvr := recover(); rcvr != nil {
			err := newErr("divided by zero")
			fmt.Println(err.error())
		}
	}()
	for i := 100; i != 0; i-- {
		fmt.Println(RandDiv(20))
	}

}

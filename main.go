package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"math/rand"

	"time"
)

var carNumber = 5
var responses = make(chan string, carNumber)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println(mirroredQuery(wg))
	wg.Wait()
}

func mirroredQuery(wg *sync.WaitGroup) string {
	for i := 1; i < carNumber+1; i++ {
		car := i
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			addCar(time.Duration(rand.Intn(5))*time.Second, car, wg)
		}(wg)
	}

	winner := <-responses // возврат самого быстрого ответа
	return winner
}

func addCar(delay time.Duration, carNumber int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Задержка старта...", delay)
	time.Sleep(delay)
	speed := rand.Intn(40)+110 // Км в час
	fmt.Println("Car", carNumber, "has speed:", speed)


	passed := 0

	for i := 0; ; i++ {
		time.Sleep(time.Second)
		passed = passed + speed
		if passed > 1500 {
			break
		}
	}

	responses <- strings.Join([]string{"the car ", " finished first!"}, strconv.Itoa(carNumber))
	runtime.Gosched()
}

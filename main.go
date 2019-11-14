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
	speed := rand.Intn(40)+110 // Км в час
	fmt.Println("Car", carNumber, "has speed:", speed)
	defer wg.Done()
	fmt.Println("Задержка старта...", delay)
	time.Sleep(delay)

	passed := 0

	for i := 0; ; i++ {
		time.Sleep(250*time.Millisecond)
		passed = passed + speed
		if passed > 6000 {
			fmt.Println("The car", carNumber, " finished!")
			break
		}
	}

	responses <- strings.Join([]string{"the car ", " won!"}, strconv.Itoa(carNumber))
	runtime.Gosched()
}

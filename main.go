package main

import (
	"fmt"

	"strconv"
	"strings"


	"math/rand"

	"time"
)

var carNumber = 5
var responses = make(chan string, carNumber)

func main() {

	fmt.Println(mirroredQuery())

}

func mirroredQuery() string {
	for i := 1; i < carNumber+1; i++ {
		car := i

		go func() {
			addCar(time.Duration(rand.Intn(5))*time.Second, car)
		}()
	}

	winner := <-responses // возврат самого быстрого ответа
	return winner
}

func addCar(delay time.Duration, carNumber int) {
	speed := rand.Intn(40)+110 // Км в час
	fmt.Println("Car", carNumber, "has speed:", speed)

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

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"time"
)

var startTime time.Time

func main() {

	for {
		startTime = time.Now()
		fmt.Println(mirroredQuery(startTime))
		time.Sleep(2000 * time.Millisecond)
	}
}

func mirroredQuery(startTime time.Time) string {

	responses := make(chan string, 3)
	go func() {
		responses <- request("http://mail.ru", startTime)
	}()
	go func() {
		responses <- request("http://google.ru", startTime)
	}()
	go func() {
		responses <- request("http://www.ru", startTime)
	}()

	tmpString := <-responses // возврат самого быстрого ответа

	return tmpString
}

func request(hostname string, startTime time.Time) string {


	client := http.Client{}
	response, err := client.Get(hostname)
	if err != nil {
		log.Println(err)
		return "Error!"
	}
	defer response.Body.Close()

	timeDifference := time.Now().Sub(startTime)

	return strings.Join([]string{hostname, response.Status, timeDifference.String()}, " - ")
}

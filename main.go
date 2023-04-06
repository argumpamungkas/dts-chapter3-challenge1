package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/novalagung/gubrak/v2"
)

func main() {

	statusWater := func(n int) {
		fmt.Print("status water: ")
		if n <= 5 {
			fmt.Println("aman")
		} else if n >= 6 && n <= 8 {
			fmt.Println("siaga")
		} else {
			fmt.Println("bahaya")
		}
	}

	statusWind := func(n int) {
		fmt.Print("status wind: ")
		if n <= 6 {
			fmt.Println("aman")
		} else if n >= 7 && n <= 15 {
			fmt.Println("siaga")
		} else {
			fmt.Println("bahaya")
		}
	}

	for {
		numberWater := gubrak.RandomInt(1, 100)
		numberWind := gubrak.RandomInt(1, 100)
		PostData(numberWater, numberWind)
		statusWater(numberWater)
		statusWind(numberWind)

		fmt.Println(strings.Repeat("=", 20))
		time.Sleep(15 * time.Second)
	}
}

func PostData(numWater, numWind int) {
	data := map[string]interface{}{
		"water": numWater,
		"wind":  numWind,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatal("err", err)
	}

	url := "https://jsonplaceholder.typicode.com/posts"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal("err", err)
	}

	result, err := client.Do(req)
	if err != nil {
		log.Fatal("err", err)
	}

	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Fatal("err", err)
	}

	log.Println(string(body))
}

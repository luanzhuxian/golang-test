package main

import (
	"fmt"
	"golang-test/weather.com/weather"
)

func main() {
	result := weather.CurrentWeather("101010100")
	fmt.Println(result)
}

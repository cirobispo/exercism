package main

import (
	"fmt"
	"meetup"
	"time"
)

func main() {
	fmt.Println( meetup.Day(meetup.First, time.Monday, 1, 2022) )
}
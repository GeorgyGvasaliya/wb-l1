package main

import (
	"time"
)

func OwnSleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)
}

func main() {
	OwnSleep(2)
}

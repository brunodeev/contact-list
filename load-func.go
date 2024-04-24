package main

import (
	"fmt"
	"time"
)

func LoadFunc() {
	points := []string{".", "..", "..."}
	for {
		for _, point := range points {
		    fmt.Printf("\rRodando%s\033[K", point)
		    time.Sleep(time.Millisecond * 10)
		}
	}
}

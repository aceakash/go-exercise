package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	go play("Ed")
	go play("Natasha")

	time.Sleep(time.Second * 30)
	fmt.Println("Time up!")
}

func play(player string) {
	total := 0
	for {
		roll := (rand.Int() % 6) + 1
		if roll == 6 {
			fmt.Printf("%s (%d) rolled a 6, rolling again immediately\n", player, total)
		} else {
			waitSec := 6 - roll
			fmt.Printf("%s (%d) rolled a %d, waiting %d sec\n", player, total, roll, waitSec)
			time.Sleep(time.Second * time.Duration(waitSec))
		}
		total += roll
	}
}

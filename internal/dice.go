package internal

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var onlyOnce sync.Once
var pickDice int

func rollDice() {
	dice := []int{1, 2, 3, 4, 5, 6}
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	pickDice = dice[rand.Intn(len(dice))]
	fmt.Println("pickDice:", pickDice)
}

func guessDice(text string) (bool, string, string, int) {
	if text != "small" && text != "big" {
		panic("You can choose only one big or small")
	}
	if text == "small" && pickDice <= 3 {
		oldPickDice := pickDice
		rollDice()
		return true, text, "win", oldPickDice
	}
	if text == "big" && pickDice > 3 {
		oldPickDice := pickDice
		rollDice()
		return true, text, "win", oldPickDice

	}
	oldPickDice := pickDice
	rollDice()
	return false, text, "lost", oldPickDice
}

func stopRollingDice() {
	pickDice = 0
}

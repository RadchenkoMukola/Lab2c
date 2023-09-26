package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

func fight(a int, b int, o chan int) {
	if rand.Intn(2) == 1 {
		o <- b
	} else {
		o <- a
	}
}

func main() {
	l := list.New()
	l.Init()

	for i := 0; i < 32; i++ {
		l.PushBack(0)
		l.PushBack(1)
	}

	round := 1

	for l.Len() > 1 {
		o := make(chan int)

		val := l.Front()

		fmt.Printf("Round %d:\n", round)

		for val != nil {
			go fight(val.Value.(int), val.Next().Value.(int), o)
			val = val.Next().Next()
		}

		iterations := l.Len() / 2

		l = list.New()
		l.Init()

		for i := 0; i < iterations; i++ {
			winner := <-o
			winnerName := "Гуань-Янь" // За замовчуванням вважаємо, що переміг Гуань-Янь (значення 0).

			if winner == 1 {
				winnerName = "Гуань-Інь" // Якщо переможець - 1, то він є Гуань-Інь.
			}

			fmt.Printf("   Fight %d: Переміг боєць монастирю %s\n", i+1, winnerName)
			l.PushBack(winner)
		}

		round++
	}

	finalWinnerName := "Гуань-Янь" // За замовчуванням вважаємо, що фінальний переможець - Гуань-Янь (значення 0).

	if l.Front().Value.(int) == 1 {
		finalWinnerName = "Гуань-Інь" // Якщо фінальний переможець - 1, то він є Гуань-Інь.
	}

	fmt.Printf("Переможець монастирь %s забирає статую боддісатви.\n", finalWinnerName)
}

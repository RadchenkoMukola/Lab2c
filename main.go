package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	monastery1 := []int{10, 15, 8, 20}
	monastery2 := []int{5, 25, 30, 12}

	// Об'єднуємо список двох монастирів в один список.
	energy := append(monastery1, monastery2...)

	// Створюємо канал для обміну результатами між потоками.
	resultChan := make(chan int)

	// Створюємо WaitGroup для синхронізації горутин.
	var wg sync.WaitGroup

	// Створюємо список для зберігання індексів ченців.
	indexes := make([]int, len(energy))
	for i := range indexes {
		indexes[i] = i
	}

	// Перемішуємо індекси ченців для рандомізації парування.
	rand.Shuffle(len(indexes), func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})

	// Розпочинаємо боротьбу між ченцями в багатопоточному режимі.
	for i := 0; i < len(indexes); i += 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			index1, index2 := indexes[i], indexes[i+1]
			energy1, energy2 := energy[index1], energy[index2]
			winner := battle(energy1, energy2)
			fmt.Printf("Бій між ченцем з енергією %d і ченцем з енергією %d. Переможець: ченец з енергією %d\n", energy1, energy2, winner)
			resultChan <- winner
		}(i)
	}

	// Закриваємо канал після завершення всіх боїв.
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Очікуємо результати боїв та визначаємо переможця в фінальному бою.
	finalWinner := -1
	for winner := range resultChan {
		if finalWinner == -1 {
			finalWinner = winner
		} else {
			finalWinner = battle(finalWinner, winner)
		}
	}

	// Визначаємо, який монастир переміг на основі переможця фінального бою.
	var winningMonastery string
	if finalWinner <= len(monastery1) {
		winningMonastery = "Гуань-Інь"
	} else {
		winningMonastery = "Гуань-Янь"
	}

	fmt.Printf("Переможець монастиря %s забирає статую боддісатви.\n", winningMonastery)
}

// Функція для визначення переможця між двома ченцями.
func battle(energy1, energy2 int) int {
	if energy1 > energy2 {
		return energy1
	}
	return energy2
}

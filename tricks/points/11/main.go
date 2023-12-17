package main

import "fmt"

type player struct {
	name  string
	score int
}

func incrementScore(s int) {
	s += 10
}

func incrementScorePoint(s *int) {
	newScore := *s + 10
	s = &newScore // Тут идет переопределение копии аргумента внутри функции
}

func incrementScorePoint2(p *player) {
	p.score += 10
}

func main() {
	score := 20
	incrementScore(score)
	// Что выведет ?
	fmt.Println("The score is", score) // Prints: "The score is 20"

	// А сейчас?
	incrementScorePoint(&score)
	fmt.Println("The score is", score) // Prints: "The score is 20"

	// Initialize a player struct and assign it to the variable p1.
	p1 := player{name: "Alice", score: 20}
	// Pass a pointer to p1 to incrementScore().
	incrementScore(&p1)
	fmt.Printf("The score for %s is %d", p1.name, p1.score) // Prints: "The score for Alice is 30"

}

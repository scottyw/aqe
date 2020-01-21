package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	questions, answers := generate(20)

	fmt.Println("------------------------------------------------")
	fmt.Println("Questions")
	fmt.Println("------------------------------------------------")
	fmt.Println("")

	for i, q := range questions {
		fmt.Printf("Q%d: %s\n\n", i+1, q)
	}

	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("Answers")
	fmt.Println("------------------------------------------------")
	fmt.Println("")

	for i, a := range answers {
		fmt.Printf("Q%d: %d\n\n", i+1, a)
	}

}

func generate(count int) ([]string, []int) {
	questions := make([]string, count)
	answers := make([]int, count)

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	for i := 0; i < count; i++ {
		q, a := generateRandomQuestion(r)
		questions[i] = q
		answers[i] = a
	}

	return questions, answers
}

func generateRandomQuestion(r *rand.Rand) (string, int) {
	return questionGenerators[r.Intn(len(questionGenerators))](r)
}

var questionGenerators = []func(r *rand.Rand) (string, int){
	generateMultiplication,
	generateAddition1,
	generateAddition2,
	generateSubtraction,
	generateDivision,
}

func generateMultiplication(r *rand.Rand) (string, int) {
	x := n(r, 3)
	y := n(r, 2)
	q := fmt.Sprintf("Multiply %d and %d.", x, y)
	a := x * y
	return q, a
}

func generateAddition1(r *rand.Rand) (string, int) {
	w := n(r, 2) * 1000
	x := n(r, 2) * 100
	y := n(r, 2) * 10
	z := n(r, 2)
	q := fmt.Sprintf("Add %d, %d, %d and %d.", w, x, y, z)
	a := w + x + y + z
	return q, a
}

func generateAddition2(r *rand.Rand) (string, int) {
	x := n(r, 3)
	y := n(r, 3)
	z := n(r, 2)
	q := fmt.Sprintf("Add %d, %d and %d.", x, y, z)
	a := x + y + z
	return q, a
}

func generateSubtraction(r *rand.Rand) (string, int) {
	x := n(r, 5)
	y := n(r, 5)
	if x > y {
		tmp := y
		y = x
		x = tmp
	}
	q := fmt.Sprintf("Subtract %d from %d.", x, y)
	a := y - x
	return q, a
}

func generateDivision(r *rand.Rand) (string, int) {
	x := r.Intn(10) + 11
	y := r.Intn(40) + 11
	q := fmt.Sprintf("Divide %d by %d.", x*y, x)
	a := y
	return q, a
}

// Generate a random number that is N digits long
func n(r *rand.Rand, n int) int {
	max := int(math.Pow10(n))
	offset := int(math.Pow10(n - 1))
	return r.Intn(max-offset+1) + offset
}

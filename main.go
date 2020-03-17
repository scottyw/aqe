package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	questions, answers := generate(p4, 20)

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

func generate(qs []func(r *rand.Rand) (string, int), count int) ([]string, []int) {
	questions := make([]string, count)
	answers := make([]int, count)

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	for i := 0; i < count; i++ {
		q, a := generateRandomQuestion(r, qs)
		questions[i] = q
		answers[i] = a
	}

	return questions, answers
}

func generateRandomQuestion(r *rand.Rand, qs []func(r *rand.Rand) (string, int)) (string, int) {
	return qs[r.Intn(len(qs))](r)
}

var p6 = []func(r *rand.Rand) (string, int){
	generateMultiplication2,
	generateAddition1,
	generateAddition2,
	generateSubtraction2,
	generateDivision,
}

var p4 = []func(r *rand.Rand) (string, int){
	generateMultiplication1,
	generateAddition3,
	generateSubtraction1,
}

func generateMultiplication1(r *rand.Rand) (string, int) {
	x := n(r, 3)
	y := n(r, 1)
	q := fmt.Sprintf("Multiply %d by %d.", x, y)
	a := x * y
	return q, a
}

func generateMultiplication2(r *rand.Rand) (string, int) {
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

func generateAddition3(r *rand.Rand) (string, int) {
	x := n(r, 3)
	y := n(r, 3)
	q := fmt.Sprintf("Add %d and %d.", x, y)
	a := x + y
	return q, a
}

func generateSubtraction1(r *rand.Rand) (string, int) {
	x := n(r, 3)
	y := n(r, 3)
	if x > y {
		tmp := y
		y = x
		x = tmp
	}
	q := fmt.Sprintf("Subtract %d from %d.", x, y)
	a := y - x
	return q, a
}

func generateSubtraction2(r *rand.Rand) (string, int) {
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

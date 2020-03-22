package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type generator func(r *rand.Rand) (string, int)

func main() {

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	fmt.Printf(generateGrid(r) + "\n\n")

	questions, answers := generate(r, p6, 20)

	for i, q := range questions {
		fmt.Printf("Q%d: %s\n\n", i+1, q)
	}

	for i, a := range answers {
		fmt.Printf("A%d: %d \t[%s]\n\n", i+1, a, questions[i])
	}

}

func generate(r *rand.Rand, qs []generator, count int) ([]string, []int) {
	questions := make([]string, count)
	answers := make([]int, count)

	for i := 0; i < count; i++ {
		q, a := generateRandomQuestion(r, qs)
		questions[i] = q
		answers[i] = a
	}

	return questions, answers
}

func generateRandomQuestion(r *rand.Rand, qs []generator) (string, int) {
	return qs[r.Intn(len(qs))](r)
}

//
// Generator sets
//

var p6 = []generator{
	generateMultiplication(2, 2),
	generateMultiplication(3, 2),
	generateAdditionWithZeros,
	generateAdditionTriple(3, 2, 2),
	generateAdditionTriple(3, 3, 2),
	generateSubtraction(4, 4),
	generateSubtraction(4, 5),
	generateSubtraction(5, 5),
	generateDivision(50),
	generateDivision(99),
}

var p4 = []generator{
	generateMultiplication(2, 1),
	generateMultiplication(3, 1),
	generateAddition(3, 2),
	generateAddition(3, 3),
	generateSubtraction(3, 2),
	generateSubtraction(3, 3),
	generateDivision(12),
}

//
// Problem generators
//

func generateMultiplication(xl, yl int) generator {
	return func(r *rand.Rand) (string, int) {
		x := n(r, xl)
		y := n(r, yl)
		q := fmt.Sprintf("Multiply %d and %d.", x, y)
		a := x * y
		return q, a
	}
}

func generateAdditionWithZeros(r *rand.Rand) (string, int) {
	w := n(r, 2) * 1000
	x := n(r, 2) * 100
	y := n(r, 2) * 10
	z := n(r, 2)
	q := fmt.Sprintf("Add %d, %d, %d and %d.", w, x, y, z)
	a := w + x + y + z
	return q, a
}

func generateAddition(xl, yl int) generator {
	return func(r *rand.Rand) (string, int) {
		x := n(r, xl)
		y := n(r, yl)
		q := fmt.Sprintf("Add %d and %d.", x, y)
		a := x + y
		return q, a
	}
}

func generateAdditionTriple(xl, yl, zl int) generator {
	return func(r *rand.Rand) (string, int) {
		x := n(r, xl)
		y := n(r, yl)
		z := n(r, zl)
		q := fmt.Sprintf("Add %d, %d and %d.", x, y, z)
		a := x + y + z
		return q, a
	}
}

func generateSubtraction(xl, yl int) generator {
	return func(r *rand.Rand) (string, int) {
		x := n(r, xl)
		y := n(r, yl)
		if x > y {
			tmp := y
			y = x
			x = tmp
		}
		q := fmt.Sprintf("Subtract %d from %d.", x, y)
		a := y - x
		return q, a
	}
}

func generateDivision(max int) generator {
	return func(r *rand.Rand) (string, int) {
		x := r.Intn(12) + 1      // 1-12
		y := r.Intn(max-10) + 11 // 11-50
		q := fmt.Sprintf("Divide %d by %d.", x*y, x)
		a := y
		return q, a
	}
}

// Generate a random number that is N digits long
func n(r *rand.Rand, n int) int {
	max := int(math.Pow10(n))
	offset := int(math.Pow10(n - 1))
	return r.Intn(max-offset+1) + offset
}

func generateGrid(r *rand.Rand) string {
	x := r.Perm(12)
	y := r.Perm(12)
	return fmt.Sprintf(grid,
		x[0]+1, x[1]+1, x[2]+1, x[3]+1, x[4]+1, x[5]+1,
		x[6]+1, x[7]+1, x[8]+1, x[9]+1, x[10]+1, x[11]+1,
		y[0]+1, y[1]+1, y[2]+1, y[3]+1, y[4]+1, y[5]+1,
		y[6]+1, y[7]+1, y[8]+1, y[9]+1, y[10]+1, y[11]+1)
}

var grid = `
┏━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┓
┃  x ┃ %2d ┃ %2d ┃ %2d ┃ %2d ┃ %2d ┃ %2d ┃ %2d ┃ %2d ┃ %2d ┃ %2d ┃ %2d ┃ %2d ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ %2d ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫
┃ SQ ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃    ┃
┗━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┛
`

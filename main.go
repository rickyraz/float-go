package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// Demonstrating floating-point precision issues in Go
	// x := 0.1
	// for i := 0; i < 10; i++ {
	// 	x += 0.1
	// }
	// fmt.Printf("%.17f\n", x)

	// Demonstrating integer overflow in Go
	// x := float64(9007199254740992) // 2^53
	// y := x + 1

	// fmt.Printf("x = %.0f\n", x)
	// fmt.Printf("y = %.0f\n", y)
	// fmt.Printf("x == y? %v\n", x == y)

	// Using arbitrary-precision decimal arithmetic to avoid precision issues
	x := decimal.NewFromFloat(0)
	for i := 0; i < 10; i++ {
		x = x.Add(decimal.NewFromFloat(0.1))
	}
	fmt.Println(x.String()) // "1"

	// Demonstrating the difference between creating decimals from float vs string

	// Dari float (approximate)
	d1 := decimal.NewFromFloat(0.1)
	fmt.Printf("NewFromFloat(0.1) = %s\n", d1.String())

	// Dari string (exact)
	d2, _ := decimal.NewFromString("0.1")
	fmt.Printf("NewFromString(0.1) = %s\n", d2.String())

	// Apakah sama?
	fmt.Printf("Equal? %v\n", d1.Equal(d2))

	BestPractice()
}

func BestPractice() {
	fmt.Println("=== Test 1: Kode asli Anda ===")
	x1 := decimal.NewFromFloat(0.1)
	for range 10 {
		x1 = x1.Add(decimal.NewFromFloat(0.1))
	}
	fmt.Printf("Result: %s (Expected: 1.1)\n\n", x1.String())

	fmt.Println("=== Test 2: Start dari 0 ===")
	x2 := decimal.NewFromFloat(0)
	for range 10 {
		x2 = x2.Add(decimal.NewFromFloat(0.1))
	}
	fmt.Printf("Result: %s (Expected: 1)\n\n", x2.String())

	fmt.Println("=== Test 3: NewFromString (Best Practice) ===")
	x3 := decimal.Zero
	point1, _ := decimal.NewFromString("0.1")
	for range 10 {
		x3 = x3.Add(point1)
	}
	fmt.Printf("Result: %s (Expected: 1)\n\n", x3.String())

	fmt.Println("=== Test 4: Multiply (Paling Efisien) ===")
	point1V2, _ := decimal.NewFromString("0.1")
	x4 := point1V2.Mul(decimal.NewFromInt(10))
	fmt.Printf("Result: %s (Expected: 1)\n", x4.String())
}

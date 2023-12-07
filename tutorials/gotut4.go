package tutorials

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
	Tutorial four covers the numbers and using the built in Math package
*/

func TutorialFour() {
	fmt.Println("5 + 4", 5+4)
	fmt.Println("5 - 4", 5-4)
	fmt.Println("5 * 4", 5*4)
	fmt.Println("5 / 4", 5/4)
	fmt.Println("5 % 4", 5%4)

	mInt := 1
	mInt++

	fmt.Println("float precision = ", 0.2222222222222222222)

	seedSecs := time.Now().Unix()
	randSeed := rand.New(rand.NewSource(seedSecs))
	randNumFromSeed := randSeed.Intn(50) + 1
	fmt.Println("Random:", randNumFromSeed)

	fmt.Println("Abs(-10) =", math.Abs(-10))
	fmt.Println("Pow(4, 2) =", math.Pow(4, 2))
	fmt.Println("Sqrt(16) =", math.Sqrt(16))
	fmt.Println("Cbrt(8) =", math.Cbrt(8))
	fmt.Println("Ceil(4.4) =", math.Ceil(4.4))
	fmt.Println("Floor(4.4) =", math.Floor(4.4))
	fmt.Println("Round(4.4) =", math.Round(4.4))
	fmt.Println("Log2(8) =", math.Log2(8))
	fmt.Println("Log10(100) =", math.Log10(100))
	fmt.Println("Log(7.389) =", math.Log(7.389))
	fmt.Println("Max(5, 4) =", math.Max(5, 4))
	fmt.Println("Min(5, 4) =", math.Min(5, 4))

	r90 := 90 * math.Pi / 100
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	fmt.Println("Sin(90) =", math.Sin(90))
}

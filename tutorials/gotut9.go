package tutorials

import (
	"fmt"
	"log"
)

func TutorialNine() {
	sayHello()
	log.Println(getSum(5,4))
	
	ans, err := getQuotient(5, 4)
	if err == nil {
		fmt.Printf("%.2f / %.2f = %f\n", 5.0, 4.0, ans)
	} else {
		log.Println(err)
	}

	f1, f2 := getTwo(5)
	fmt.Printf("%d %d\n", f1, f2)

	log.Println("Unknown sum:", getSum2(1, 2, 3, 4))

	vArr := []int{1, 2, 3, 4}
	log.Println("Array sum:", getArraySum(vArr))

	f3 := 5
	log.Println("f3 before func:", f3)
	cVal := changeVal(f3)
	log.Println("f3 after func:", cVal)
}

func sayHello() {
	log.Println("Hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("you can't divide by 0")
	} else {
		return x / y, nil
	}
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getSum2(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	
	for _, num := range arr {
		sum += num
	}

	return sum
}

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

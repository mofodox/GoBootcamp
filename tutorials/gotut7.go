package tutorials

import (
	"fmt"
	"log"
)

func TutorialSeven() {
	var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	log.Println("index 0:", arr2[0])
	log.Println("array length:", len(arr2))

	for i := 0; i < len(arr2); i++ {
		log.Println(arr2[i])
	}

	for i, v := range arr2 {
		fmt.Printf("%d: %d\n", i, v)
	}

	// multi dimensional
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			log.Println(arr3[i][j])
		}
	}

	// Runes
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune array: %d\n", v)
	}

	byteArr := []byte{
		'a',
		'b',
		'c',
	}

	bStr := string(byteArr[:])
	log.Println("I'm a string:", bStr)
}
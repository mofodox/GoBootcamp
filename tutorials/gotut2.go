package tutorials

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
	Tutorial two covers the variables and datatypes
*/

func TutorialTwo() {
	// var vName string = "Kai"
	// var v1, v2 = 1.2, 3.4
	// var v3 = "Hello"

	// v1 = 4.5 // assigning a new value to v1 variable
	fmt.Println(reflect.TypeOf(25))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello"))
	fmt.Println(reflect.TypeOf("👨‍💻"))

	cV1 := 1.5
	cV2 := int(cV1)
	fmt.Println(cV2)

	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3)
	fmt.Println(cV4, err, reflect.TypeOf(cV4))

	cV5 := 5000000
	cV6 := strconv.Itoa(cV5)
	fmt.Println(cV6)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		fmt.Println(cV8)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	fmt.Println(cV9)
}

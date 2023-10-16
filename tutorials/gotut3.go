package tutorials

import "fmt"

/*
	Tutorial three covers the if statement and formatting
*/

/*
	Formatted Output: fmt.Printf / fmt.Sprintf

	%d => Integer
	%c => Character
	%f => Float
	%t => Boolean
	%s => String
	%o => Base 8
	%x => Base 16
	%v => Guesses based on data type
	%T => Type of supplied value
*/

func TutorialThree() {
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		fmt.Println("Important birthday")
	} else if (iAge == 21) || (iAge == 50) {
		fmt.Println("Important birthday")
	} else if iAge >= 65 {
		fmt.Println("Important birthday")
	} else {
		fmt.Println("Not important birthday")
	}
}

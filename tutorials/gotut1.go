package tutorials

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
	Tutorial one is the introduction of Go language
*/

func TutorialOne() {
	fmt.Println("Hello Go!")

	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println("Hello", name)
	} else {
		log.Fatal(err)
	}
}

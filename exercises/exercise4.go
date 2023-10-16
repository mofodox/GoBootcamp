package exercises

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Exercise4Tutorial() {
	// Enter number 1
	firstInput := bufio.NewReader(os.Stdin)
	// Enter number 2
	secondInput := bufio.NewReader(os.Stdin)

	fmt.Println("Type your first number:")
	firstNum, err := firstInput.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	firstNum = strings.TrimSpace(firstNum)

	fmt.Println("Type your second number:")
	secondNum, err := secondInput.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	secondNum = strings.TrimSpace(secondNum)

	iFirstNum, err := strconv.Atoi(firstNum)
	if err != nil {
		log.Fatal(err)
	}

	iSecondNum, err := strconv.Atoi(secondNum)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d + %d = %d\n", iFirstNum, iSecondNum, iFirstNum+iSecondNum)
}

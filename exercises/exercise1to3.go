package exercises

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ExerciseOneToThreeTutorial() {
	fmt.Println("What is your age?")

	reader := bufio.NewReader(os.Stdin)
	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	age = strings.TrimSpace(age)

	intTrimmedAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}

	if intTrimmedAge < 5 {
		fmt.Println("Too young for school")
	} else if intTrimmedAge == 5 {
		fmt.Println("Go to Kindergarten")
	} else if (intTrimmedAge > 5) && (intTrimmedAge <= 17) {
		grade := intTrimmedAge - 5
		fmt.Printf("Go to Grade %d\n", grade)
	} else {
		fmt.Println("Go to College")
	}
}

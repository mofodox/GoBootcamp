package tutorials

import (
	"fmt"
	"log"
	"time"
)

/*
	Tutorial six covers on working in built in time package
*/

func TutorialSix() {
	now := time.Now()
	log.Println(now.Year(), now.Month(), now.Day())
	log.Println(now.Hour(), now.Minute(), now.Second())
	loc, err := time.LoadLocation("EET")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Time in Palestine %s\n", now.In(loc))

	locEST, _ := time.LoadLocation("EST")
	locUTC, _ := time.LoadLocation("UTC")
	locMST, _ := time.LoadLocation("MST")
	fmt.Printf("EST: %s\n", now.In(locEST))
	fmt.Printf("UTC: %s\n", now.In(locUTC))
	fmt.Printf("MST: %s\n", now.In(locMST))

	birthDate := time.Date(1991, time.July, 26, 15, 33, 0, 0, time.Local)
	diff := now.Sub(birthDate)
	fmt.Printf("Days Alive: %d days\n", int(diff.Hours() / 24))
	fmt.Printf("Hours Alive: %d hours\n", int(diff.Hours()))
}
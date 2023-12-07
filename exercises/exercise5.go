package exercises

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

/*
	Hangman game
	---
	head =  0
	body = /|\
	legs = / \
	---

	Guess secret word = M__K__
	Show incorrect guesses = A

	Guess a letter: Y
	Sorry You Lose! The word is MONKEY
	Yes, the secret word is MONKEY

	Please enter only one letter
	Please enter a letter
	Please enter a letter you haven't guessed
*/

var hmArr = [7]string{
	" +---+\n" +
	"     |\n" +
	"     |\n" +
	"     |\n" +
	"    ===\n",
	" +---+\n" +
	" 0   |\n" +
	"     |\n" +
	"     |\n" +
	"    ===\n",
	" +---+\n" +
	" 0   |\n" +
	"/|     |\n" +
	"     |\n" +
	"    ===\n",
	" +---+\n" +
	" 0   |\n" +
	"/|\\ |\n" +
	"     |\n" +
	"    ===\n",
	" +---+\n" +
	" 0   |\n" +
	"/|\\ |\n" +
	"/     |\n" +
	"    ===\n",
	" +---+\n" +
	" 0   |\n" +
	"/|\\ |\n" +
	"/ \\ |\n" +
	"    ===\n",
}

var wordArr = [7]string{
	"JAZZ",
	"ZIGZAG",
	"ZILCH",
	"ZIPPER",
	"ZODIAC",
	"ZOMBIE",
	"FLUFF",
}

var randWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

func Exercise5Tutorial() {
	fmt.Println(getRandWord())

	for {
		// Show game board
		showBoard()

		// Get a letter from the user
		guess := getUserLetter()

		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)
			if sliceHasEmpty(correctLetters) {
				fmt.Println("More letters to guess")
			} else {
				fmt.Println("Yes the secret word is", randWord)
				break
			}
		} else {
			guessedLetters += guess
			wrongGuesses = append(wrongGuesses, guess)

			if len(wrongGuesses) >= 6 {
				fmt.Println("Sorry you are dead! The word is", randWord)
				break
			}
		}
	}
}

func getRandWord() string {
	seedSecs := time.Now().Unix()
	rand.New(rand.NewSource(seedSecs))
	randWord = wordArr[rand.Intn(7)]
	correctLetters = make([]string, len(randWord))
	return randWord
}

func showBoard() {
	fmt.Println(hmArr[len(wrongGuesses)])
	fmt.Print("Secret Word: ")

	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}

	fmt.Print("\nIncorrect guesses: ")

	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}

	fmt.Println()
}

func getUserLetter() string {
	reader := bufio.NewReader(os.Stdin)
	var guess string

	for true {
		fmt.Printf("\nGuess a letter: ")

		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)
		
		var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(guess)

		if len(guess) != 1 {
			fmt.Println("Please enter only one letter")
		} else if !IsLetter {
			fmt.Println("Please enter a letter")
		} else if strings.Contains(guessedLetters, guess) {
			fmt.Println("Please enter a letter tha you haven't guessed")
		} else {
			return guess
		}
	}

	return guess
}

func updateCorrectLetters(letter string) {
	indexMatches := getAllIndexes(randWord, letter)
	for _, v := range indexMatches {
		correctLetters[v] = letter
	}
}

func getAllIndexes(str, subStr string) (indices []int) {
	if (len(subStr) == 0) || (len(str) == 0) {
		return indices
	}

	offset := 0
	for {
		i := strings.Index(str[offset:], subStr)
		if i == -1 {
			return indices
		}
		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}
}

func sliceHasEmpty(slice []string) bool {
	for _, v := range slice {
		if len(v) == 0 {
			return true
		}
	}

	return false
}

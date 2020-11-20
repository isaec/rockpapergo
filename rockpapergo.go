package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //seed the random
	choices := [4]string{"", "rock", "paper", "scissors"}
	//"" is due to older sins

	compInt := rand.Intn(3) + 1

	var playChoice string

	fmt.Println("please choose from: 1:rock, 2:paper, 3:scissors, I choose " + choices[compInt] + " btw")

	fmt.Scanln(&playChoice)

	playInt, err := strconv.Atoi(playChoice)

	if err != nil || (3 < playInt || playInt < 1) {
		fmt.Println("number, 1 to 3. run again.")
		os.Exit(1) //ty misterturtle
	}

	playChoice = choices[playInt]

	if playInt == compInt { //there is one case where this doesnt work... but meh
		fmt.Println("tie. " + choices[playInt] + " = " + choices[compInt])
	} else if playInt-1 == compInt || playInt+2 == compInt {
		fmt.Println("you win, " + choices[playInt] + " beats " + choices[compInt])
	} else {
		fmt.Println("you loose, " + choices[playInt] + " looses to " + choices[compInt])
	}

}

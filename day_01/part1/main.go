package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalFuel int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		fuel := calcFuel(mass)
		totalFuel = totalFuel + fuel
		log.Printf("mass %d fuel %d", mass, fuel)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("total fuel %d", totalFuel)
}

func calcFuel(mass int) int {
	return (mass / 3) - 2
}

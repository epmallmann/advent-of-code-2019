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
		fuel := calcFuelRecursive(mass)
		totalFuel = totalFuel + fuel
		log.Printf("mass %d fuel %d", mass, fuel)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("total fuel %d", totalFuel)
}

func calcFuelRecursive(mass int) int {
	fuel := calcFuel(mass)
	if fuel > 0 {
		aux := calcFuelRecursive(fuel)
		if aux > 0 {
			fuel += aux
		}
	}
	return fuel
}

func calcFuel(mass int) int {
	return (mass / 3) - 2
}

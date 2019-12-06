package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	output := calculateIntcode(input)

	log.Print(output)
}

func calculateIntcode(input string) string {
	log.Printf("input %s", input)

	stringArr := strings.Split(input, ",")
	var intArr = []int{}
	for _, i := range stringArr {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		intArr = append(intArr, j)
	}

	outputArr := append(intArr[:0:0], intArr...)

	opcode := intArr[0]
	log.Printf("opcode %d", opcode)
	if opcode == 99 {
		return ""
	}

	start := 0

	for {
		partToAnalyze := outputArr[start:]

		log.Printf("partToAnalyze %v", partToAnalyze)

		opcode = partToAnalyze[0]

		if opcode == 99 {
			log.Print("opcode=99 breaking")
			break
		}

		partToAnalyze = outputArr[start : start+4]

		doTheThing(partToAnalyze, outputArr)

		start += 4
	}

	var outputStrArr []string
	for _, i := range outputArr {
		outputStrArr = append(outputStrArr, strconv.Itoa(i))
	}

	return strings.Join(outputStrArr, ",")
}

func doTheThing(part []int, outputArr []int) {
	log.Printf("slice %v", part)

	opcode := part[0]

	if opcode == 99 {
		return
	}

	input1 := outputArr[part[1]]
	input2 := outputArr[part[2]]
	output := part[3]

	log.Printf("input1 %d, input2 %d, output %d", input1, input2, output)

	var result int

	if opcode == 1 {
		result = input1 + input2
	} else if opcode == 2 {
		result = input1 * input2
	}

	log.Printf("result %d", result)

	outputArr[output] = result

	log.Printf("intArr %v", outputArr)
}

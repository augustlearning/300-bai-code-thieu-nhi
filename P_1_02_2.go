package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func arrayToString(array [30]int, length int) string {
	var retstring, substring string
	substring = ""
	retstring = "{"
	for i := 0; i < length; i++ {
		substring = strconv.Itoa(array[i])
		retstring = retstring + substring
		if i < length-1 {
			retstring = retstring + ", "
		} else {
			retstring = retstring + "}"
		}
	}

	return retstring
}

func writeLineAppend(output string) {
	f, err := os.OpenFile("SUBSET.OUT", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	output = output + "\n"
	if _, err = f.WriteString(output); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var string_output string
	var i int
	var j int
	var n, k int

	input, err := os.ReadFile("SUBSET.INP")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(input))

	n = 0
	k = 0

	SplitResult := strings.Split(string(input), " ")

	n, err = strconv.Atoi(string(SplitResult[0]))
	k, err = strconv.Atoi(string(SplitResult[1]))

	string_output = ""
	i = 0
	j = 0
	arrayX := [30]int{}

	for i = 0; i < k; i++ {
		arrayX[i] = i + 1
	}

	for {
		string_output = arrayToString(arrayX, k)
		writeLineAppend(string_output)

		i = k - 1

		for {
			if i < 0 {
				break
			}

			if arrayX[i] != n-k+i+1 {
				break
			}

			i = i - 1
		}

		if i < 0 {
			break
		}

		arrayX[i] = arrayX[i] + 1

		for j = i + 1; j < k; j++ {
			arrayX[j] = arrayX[j-1] + 1
		}
	}
}

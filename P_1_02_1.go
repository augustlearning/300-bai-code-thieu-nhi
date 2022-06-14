package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func arrayToString(array [30]int, length int) string {
	var retstring, substring string
	substring = ""
	retstring = ""
	for i := 0; i < length; i++ {
		substring = strconv.Itoa(array[i])
		retstring = retstring + substring
	}

	return retstring
}

func writeLineAppend(output string) {
	f, err := os.OpenFile("BSTR.OUT", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	output = output + "\n"
	if _, err = f.WriteString(output); err != nil {
		log.Fatal(err)
	}
}

func originalWay() {
	var string_output string
	var i int
	var j int
	input, err := os.ReadFile("BSTR.INP")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(input))

	ilength_input, err := strconv.Atoi(string(input))

	string_output = ""
	i = 0
	arrayX := [30]int{}

	for {
		string_output = arrayToString(arrayX, ilength_input)
		writeLineAppend(string_output)

		i = ilength_input - 1

		for {
			if i < 0 {
				break
			}

			if arrayX[i] != 1 {
				break
			}

			i = i - 1
		}

		if i < 0 {
			break
		}

		arrayX[i] = 1

		for j = i + 1; j < ilength_input; j++ {
			arrayX[j] = 0
		}
	}
}

// Example: recursiveGenString(3, "")
func recursiveGenString(length int, sub string) {
	if len(sub) == length {
		writeLineAppend(sub)
	} else {
		recursiveGenString(length, sub+"0")
		recursiveGenString(length, sub+"1")
	}
}

func main() {
	originalWay()
}

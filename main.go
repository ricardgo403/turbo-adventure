package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func showElements(mySlice []string) {
	for i, element := range mySlice {
		fmt.Println("i: ", i, "var: ", element)
	}
}

func saveToFile(filename string, mySlice []string) {
	if strings.Contains(filename, ".txt") {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		for i := 0; i < len(mySlice); i++ {
			file.WriteString(mySlice[i] + "\n")
		}
	} else {
		fmt.Println("Filename must include the .txt extension")
	}
}

func sortAscAndSave(mySlice []string) {
	sort.Strings(mySlice)
	showElements(mySlice)
	saveToFile("ascendente.txt", mySlice)
}

func sortDescAndSave(mySlice []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(mySlice)))
	showElements(mySlice)
	saveToFile("descendente.txt", mySlice)
}

func main() {
	var n uint64
	fmt.Print("Cuántos strings va a capturar?\n>>> ")
	fmt.Scan(&n)

	strings := make([]string, 0, n)
	for i := uint64(0); i < n; i++ {
		var myString string
		fmt.Print(">>> ")
		fmt.Scan(&myString)
		strings = append(strings, myString)
	}
	showElements(strings)
	sortAscAndSave(strings)
	sortDescAndSave(strings)
}

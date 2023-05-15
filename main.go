package main

import (
	"fmt"
	"strings"
	"sync"
)

func cleaningString(dataInput string, character ...string) (output string) {
	output = dataInput
	wg := sync.WaitGroup{}
	for _, char := range character {
		wg.Add(1)
		go func(char string) {
			output = strings.Replace(output, char, "", -1)
			wg.Done()
		}(char)

	}
	wg.Wait()
	return
}

func GetSimilarityValue(search string, data string) (value int) {
	cleanData := cleaningString(data, ",", "!", "@", "#")
	arrData := strings.Split(cleanData, " ")
	wg := sync.WaitGroup{}
	for _, dataIn := range arrData {
		wg.Add(1)
		go func(dataIn string) {
			if strings.Contains(dataIn, search) {
				value += 1
			}
			wg.Done()
		}(dataIn)

	}
	wg.Wait()
	return
}

func main() {
	value := GetSimilarityValue("ayam", "pak dono suka ,makan ayam di ayami")
	fmt.Println("Value: ", value)
}

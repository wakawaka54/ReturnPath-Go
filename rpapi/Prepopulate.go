/*
Reads data/data.txt and prepopulates the dataset with Sentence objects

Relies of GO_HOME environment variable and binary must be ran from project directory
or GO_HOME must be set to project path
*/

package main

import (
	"bufio"
	"encoding/json"
	"os"
)

func PrepopulateDataset() {
	homePath := os.Getenv("GO_HOME")
	file, _ := os.Open(homePath + "data/data.txt")
	reader := bufio.NewReader(file)
	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		var sentence Sentence
		parseErr := json.Unmarshal([]byte(line), &sentence)
		if parseErr == nil {
			sentences = append(sentences, sentence)
		} else {
			panic(err)
		}
	}
}

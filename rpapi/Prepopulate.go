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

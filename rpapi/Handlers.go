package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

const maxSentencesLimit = 100

//Index
/*
Route: /
*/
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Why are you here?")
	w.WriteHeader(http.StatusOK)
}

//Options
/*
Route: /
Method: OPTIONS
*/
func Options(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//GetSentences
/*
Route: api/sentences?limit=[LIMIT]&offset=[OFFSET]&id=[ID]&sentence=[SENTENCE]&tags=[TAG],[TAG2]
*/
func GetSentences(w http.ResponseWriter, r *http.Request) {
	//Process query string
	qs := r.URL.Query()

	response := Sentences{}

	//Look for filters in query string
	//Wow, is this really the easiest way to do this in Go?
	filters := SentenceCompare{}
	isFiltering := false
	if idStr, ok := qs["id"]; ok {
		id, _ := strconv.Atoi(idStr[0])
		filters.ID = &id
		isFiltering = true
	}
	if senStr, ok := qs["sentence"]; ok {
		filters.Sentence = &senStr[0]
		isFiltering = true
	}
	if tagStr, ok := qs["tags"]; ok {
		filters.Tags = tagStr
		isFiltering = true
	}

	//Apply filters
	if isFiltering {
		response = sentences.Filter(filters)
	} else {
		response = sentences
	}

	total := len(response)
	limit, offset := pagnationUtil(qs, total)
	if limit > maxSentencesLimit {
		limit = maxSentencesLimit
	}

	response = response[offset:(offset + limit)]

	//limitedSentences := sentences[itemsPerPage*page : itemsPerPage*(page+1)]
	w.Header().Set("X-Total-Count", strconv.Itoa(total))
	w.Header().Set("Access-Control-Expose-Headers", w.Header().Get("Access-Control-Expose-Headers")+",X-Total-Count")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "Error encoding internal sentences to json.")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//AddSentence
/*
Route: api/sentences
*/
func AddSentence(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	sentence := Sentence{}
	err := decoder.Decode(&sentence)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusConflict)
		return
	}

	//Assign the next incremental id to the new sentence and prepend to datastore
	sentence.ID = len(sentences)
	sentence.CreateTags()
	sentences = append(Sentences{sentence}, sentences...)

	//Return 201 - Status Created
	w.WriteHeader(http.StatusCreated)
}

//DeleteSentence
/*
Route: api/sentences/{id}
*/
func DeleteSentence(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, sentence := range sentences {
		if strconv.Itoa(sentence.ID) == id {
			sentences = append(sentences[:i], sentences[i+1:]...)

			//Return 202 - Status was successfully deleted
			w.WriteHeader(http.StatusAccepted)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "ID of %s was not found", id)
}

//SentenceStatistics
/*
Route: api/sentences/statistics
*/
func SentenceStatistics(w http.ResponseWriter, r *http.Request) {
	stats := SortableMap{}
	stats.Map = make(map[string]int)
	for _, sentence := range sentences {
		for _, tag := range sentence.Tags {
			stats.Increment(tag)
		}
	}

	sort.Sort(stats)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(stats.First(15)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}

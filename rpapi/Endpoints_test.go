package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

const sentencesToAdd int = 5000

var router http.Handler

//Tests that router is working and that index path is alive
func TestMain(t *testing.T) {
	router = SetupPipeline()
	req, _ := http.NewRequest("GET", "/api/", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert(res.Code == 200, "Response code should equal 200", t)
}

//Tests that OPTIONS request can be sent and receives OK response
func TestOptions(t *testing.T) {
	req, _ := http.NewRequest("OPTIONS", "/api/sentences", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert(res.Code == 200, "Response code should equal 200", t)
}

//Tests that no sentences are returned before sentences are entered
func TestGetSentencesEmpty(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/sentences", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	var responseSchema Sentences
	json.Unmarshal(res.Body.Bytes(), &responseSchema)

	assert(res.Code == 200, "Response code not equal 200", t)
	assert(res.Header().Get("X-Total-Count") == "0", "Total pages must equal 0", t)
	assert(len(responseSchema) == 0, "Sentences should not of been returned", t)
}

//Tests that sentences are returned
func TestGetSentencesNotEmpty(t *testing.T) {
	sentence := Sentence{
		ID:       0,
		Sentence: "This is default",
		Tags:     []string{"this", "default"},
	}

	for i := 0; i != sentencesToAdd; i++ {
		sentences = append(sentences, sentence)
		sentence.ID = i
	}

	req, _ := http.NewRequest("GET", "/api/sentences", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	var responseSchema Sentences
	json.Unmarshal(res.Body.Bytes(), &responseSchema)

	assert(res.Code == 200, "Response code should equal 200", t)
	assert(res.Header().Get("X-Total-Count") != "0", "Total pages must not equal 0", t)
	assert(len(responseSchema) != 0, "Sentences should of been returned", t)
}

//Tests that current page can be changed
func TestGetSentencesPages(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/sentences?limit=10&offset=10", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	var responseSchema Sentences
	json.Unmarshal(res.Body.Bytes(), &responseSchema)

	assert(res.Code == 200, "Response code should equal 200", t)
	assert(res.Header().Get("X-Total-Count") != "0", "Total pages must not equal 0", t)
	assert(len(responseSchema) != 0, "Sentences should of been returned", t)
}

//Tests that no sentences are returned if page is out of range
func TestGetSentencesFilterByID(t *testing.T) {
	sentence := sentences[2]
	req, _ := http.NewRequest("GET", "/api/sentences?limit=20&offset=0&id="+strconv.Itoa(sentence.ID), nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	var responseSchema Sentences
	json.Unmarshal(res.Body.Bytes(), &responseSchema)

	assert(res.Code == 200, "Response code should equal 200", t)
	assert(res.Header().Get("X-Total-Count") == "1", "Total count must equal 1", t)
	assert(responseSchema[0].ID == sentence.ID, "Sentence IDs should be equal", t)
}

//Tests that no sentences are returned if page is out of range
func TestGetSentencesPagesOutOfRange(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/sentences?limit=20&offset=10000", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	var responseSchema Sentences
	json.Unmarshal(res.Body.Bytes(), &responseSchema)

	assert(res.Code == 200, "Response code should equal 200", t)
	assert(res.Header().Get("X-Total-Count") != "0", "Total pages must not equal 0", t)
	assert(len(responseSchema) == 0, "Sentences should not of been returned", t)
}

//Tests that sentences can be added
func TestAddSentence(t *testing.T) {
	sentence := Sentence{
		Sentence: "This is a new sentence!",
		Tags:     []string{"this", "new", "sentence"},
	}

	dataBytes, _ := json.Marshal(sentence)
	data := bytes.NewReader(dataBytes)

	req, _ := http.NewRequest("POST", "/api/sentences", data)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert(res.Code == http.StatusCreated, "Response code should equal 201", t)
	assert(sentences[0].Sentence == sentence.Sentence, "First sentence should equal the added sentence", t)
}

//Tests that sentences can be added
func TestDeleteSentence(t *testing.T) {
	sentence := sentences[0]

	req, _ := http.NewRequest("DELETE", "/api/sentences/"+strconv.Itoa(sentence.ID), nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert(res.Code == http.StatusAccepted, "Response code should equal 202", t)
	assert(sentences[0].ID != sentence.ID, "First sentence of sentences should not contain the ID of deleted sentence", t)
}

//Tests that statistics can be retrieved
func TestGetStatistics(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/sentences/statistics", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	var responseSchema StatisticsSchemas
	json.Unmarshal(res.Body.Bytes(), &responseSchema)

	assert(res.Code == 200, "Response code should equal 200", t)
	assert(responseSchema[0].Count == sentencesToAdd, "Word count should equal number of sentences added", t)
}

func assert(condition bool, msg string, t *testing.T) {
	if !condition {
		t.Error(msg, condition)
	}
}

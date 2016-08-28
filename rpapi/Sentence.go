package main

import "strings"

//Application datastore
var sentences = Sentences{}

var BoringWords []string

type Sentence struct {
	ID       int      `json:"id"`
	Sentence string   `json:"sentence"`
	Tags     []string `json:"tags"`
}

type SentenceCompare struct {
	ID       *int
	Sentence *string
	Tags     []string
}

type Sentences []Sentence

func (s Sentence) Filter(compare *SentenceCompare) bool {
	if compare == nil {
		return true
	}

	if compare.ID != nil && *compare.ID != s.ID {
		return false
	}

	if compare.Sentence != nil {
		if !strings.Contains(strings.ToLower(s.Sentence), strings.ToLower(*compare.Sentence)) {
			return false
		}
	}

	if compare.Tags != nil {
		if !arrayContains(s.Tags, compare.Tags) {
			return false
		}
	}

	return true
}

func (s Sentences) Filter(comparer SentenceCompare) Sentences {
	filtered := make(Sentences, 0)

	for _, sentence := range s {
		if sentence.Filter(&comparer) {
			filtered = append(filtered, sentence)
		}
	}

	return filtered
}

func (s *Sentence) CreateTags() {
	split := strings.Split(s.Sentence, " ")
	for _, tag := range split {
		alpha, str := isStringAlpha(tag)
		if alpha {
			boring := arrayContains(BoringWords, []string{str})
			if boring == false {
				s.Tags = append(s.Tags, str)
			}
		}
	}
}

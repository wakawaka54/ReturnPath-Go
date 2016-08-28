package main

import (
	"net/url"
	"strconv"
	"strings"
)

type SortableMap struct {
	Map    map[string]int
	Sorted []string
}

func (s *SortableMap) Increment(key string) {
	if _, ok := s.Map[key]; ok == false {
		s.Sorted = append(s.Sorted, key)
	}
	s.Map[key]++
}

func (s SortableMap) Less(i, j int) bool {
	return s.Map[s.Sorted[i]] > s.Map[s.Sorted[j]]
}

func (s SortableMap) Swap(i, j int) {
	toJ := s.Sorted[i]
	toI := s.Sorted[j]

	s.Sorted[i] = toI
	s.Sorted[j] = toJ
}

func (s SortableMap) Len() int {
	return len(s.Sorted)
}

func (s SortableMap) First(n int) StatisticsSchemas {
	m := StatisticsSchemas{}
	for i := 0; i != n && i < len(s.Sorted); i++ {
		m = append(m, StatisticsSchema{Tag: s.Sorted[i], Count: s.Map[s.Sorted[i]]})
	}
	return m
}

func pagnationUtil(vals url.Values, length int) (limit, offset int) {
	limit = 20
	offset = 0

	if limitStr, ok := vals["limit"]; ok {
		limit, _ = strconv.Atoi(limitStr[0])
	}

	if offsetStr, ok := vals["offset"]; ok {
		offset, _ = strconv.Atoi(offsetStr[0])
	}

	if offset > length {
		offset = 0
		limit = 0
	}

	if offset+limit > length {
		limit = length - offset
	}

	return
}

func arrayContains(a1, a2 []string) bool {
	for _, tag := range a2 {
		subcontain := false
		for _, tag2 := range a1 {
			if tag == tag2 {
				subcontain = true
				break
			}
		}
		if !subcontain {
			return false
		}
	}

	return true
}

//Check if string has
func isStringAlpha(s string) (bool, string) {
	var filtered string

	s = strings.TrimSpace(s)
	if s == "" {
		return false, s
	}

	for _, c := range s {
		if c == '.' || c == '!' || c == ',' || c == '?' || c == '(' || c == ')' {
			continue
		} else {
			if ((c >= 'a') && (c <= 'z')) || ((c >= 'A') && (c <= 'Z')) {
				filtered = filtered + string(c)
			} else {
				return false, s
			}
		}
	}

	return true, filtered
}

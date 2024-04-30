package main

import "math"

type Recommendation struct {
	Book  Book
	Score float64
}

func recommend(allReaders []Reader, target Reader, n int) []Recommendation {
	read := newSet(target.Books...)
	recommendations := map[Book]float64{}
	for _, reader := range allReaders {
		if reader.Name == target.Name {
			continue
		}
		var similarity float64 #B
		for _, book := range reader.Books {
			if read.Contains(book) {
				similarity++
			}
			// you could also later extend to liked and dislike score
		}
		if similarity == 0 {
			continue
		}
		score := math.Log(similarity) + 1 #C
		for _, book := range readers.Books {
			if !read.Contains(book) {
				recommendations[book] += score
			}
		}
	}
	// TODO: sort by score
	// TODO: only output a certain amount of recommendations (n)
}

/*
Reference: https://tutorialedge.net/challenges/go/word-frequencies/
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Function for counting word frequencies in a given text string.
It takes a single argument text of type string. This is the input text that will be analyzed. It returns a map[string]int.
This map will store the word frequencies, where the keys are the words (strings) and the values are their corresponding counts (integers).
*/
func CountWords(text string) map[string]int {
	// This line uses the strings.Fields function to split the input text into a slice of strings.
	// Each string in the slice represents a single word from the original text.
	// The Fields function uses whitespace as the delimiter for splitting.
	words := strings.Fields(text)
	wordFrequencies := make(map[string]int)

	// Iterate through this words slice to count the occurrences of each word and store the results in the wordFrequencies map.
	for _, word := range words {
		word = strings.ToLower(word)
		wordFrequencies[word]++
	}

	return wordFrequencies
}

type Word struct {
	Text  string
	Count int
}

// This function takes a map of word frequencies (wordmap where keys are words and values are their counts) and returns a Word struct containing the most frequent word and its count.
func MostCommonWord(wordmap map[string]int) Word {
	maxCount := 0
	mostCommonWord := ""

	/*
	   This loop iterates through the input wordmap. In each iteration:
	     word will hold the current word (string).
	     count will hold the corresponding count (integer).
	   if the current word's count is greater than the current maxCount. If it is:
	   Updates maxCount to the new highest count and Updates mostCommonWord to the current word.
	   After iterating through all words in the map, the function returns a Word struct.
	*/
	for word, count := range wordmap {
		if count > maxCount {
			maxCount = count
			mostCommonWord = word
		}
	}

	return Word{Text: mostCommonWord, Count: maxCount}
}

// This function takes a word frequency map (wordmap) and returns a slice of Word structs sorted by frequency in descending order.
func MostCommonWords(wordmap map[string]int) []Word {
	/*
	   Create a slice of Word structs
	   initializes the slice with a length of 0 and a capacity equal to the length of the wordmap.
	   This is an optimization to avoid unnecessary reallocations as the slice grows.
	*/
	words := make([]Word, 0, len(wordmap))

	/*
	   This loop iterates over the wordmap. For each iteration:
	     word is the key (the word).
	    count is the value (the frequency of the word).
	   Inside the loop, a Word struct is created for each word and its count, and it's appended to the words slice.
	*/
	for word, count := range wordmap {
		words = append(words, Word{Text: word, Count: count})
	}

	// This is where the sorting happens. sort.Slice sorts a slice in-place using a provided less function.
	// The anonymous function compares the Count field of two Word structs at indices i and j.
	// It returns true if the count at i is greater than the count at j.
	// This ensures that the slice is sorted in descending order of frequency.
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})

	return words
}

// This function takes a word frequency map (wordmap) as input and returns a slice of Word structs containing the top 5 most frequent words.
func Top5Words(wordmap map[string]int) []Word {
	// This line calls the MostCommonWords function with the provided wordmap. MostCommonWords returns a slice of Word structs sorted in descending order of frequency.
	top5 := MostCommonWords(wordmap)

	// Checks if the length of the top5 slice is greater than 5.
	if len(top5) > 5 {
		// If true, creates a new slice containing only the elements from index 0 up to (but not including) index 5.
		// Effectively, this takes the first 5 elements of the sorted slice.
		top5 = top5[:5]
	}

	return top5
}

func main() {
	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

	results := CountWords(text)
	MostCommon := Top5Words(results)

	fmt.Println("*********************")
	fmt.Println("Word Frequency Test")
	fmt.Println("*********************")
	fmt.Println("Top 5 Words: ", MostCommon)
	fmt.Println("-----------")
	fmt.Println("Most common Words: ", MostCommonWords(results))
	fmt.Println("-----------")
	fmt.Println("Most common Word: ", MostCommonWord(results))
}

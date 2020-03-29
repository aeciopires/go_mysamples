package main

import (
	"fmt"
)

// Define global constants
const GlobalLimit = 100
const MaxCacheSize = 10 * GlobalLimit
const (
	CacheKeyBook = "book"
	CacheKeyCD   = "cd_"
)

// Define global variable
var cache map[string]string

// Function to return array of string
func cacheGet(key string) string {
	return cache[key]
}

// Function to set value to array of string
func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

// Function to return book ISBN
func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

// Function to set book ISBN
func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

// Function to return CD
func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

// Function to set CD
func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func main() {
	cache = make(map[string]string)
	SetBook("1234-5678", "Get Ready To Go.")
	SetCD("1234-5678", "Get Ready To Go Audio Book.")

	fmt.Println("----- Show inventory -----")
	fmt.Println("Book 1234-5678:", GetBook("1234-5678"))
	fmt.Println("CD 1234-5678:", GetCD("1234-5678"))
	fmt.Println("")
	fmt.Println("===== DEBUG =====")
	fmt.Println("[DEBUG] cache :", cache)
	fmt.Println("[DEBUG] GlobalLimit :", GlobalLimit)
	fmt.Println("[DEBUG] MaxCacheSize :", MaxCacheSize)
	fmt.Println("[DEBUG] CacheKeyBook :", CacheKeyBook)
	fmt.Println("[DEBUG] CacheKeyCD :", CacheKeyCD)
}

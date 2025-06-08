package main

import "fmt"

const GlobalSizelimit = 100
const maxCachesize int = 10 * GlobalSizelimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD = "cd_"
)

var cache map[string]string

func main() {
	cache = make(map[string]string)

	setBook("1234","Carpe diem")

	setCD("1234","Dead Poets Society")

	fmt.Println("Book : ",getBook("1234"))

	fmt.Println("CD : ",getCD("1234"))
	
	setBook("1234","Carpe diem : Sieze the day")

	fmt.Println("Updated Book : ",getBook("1234"))
}

func getItems(key string) string {
	return cache[key]
}

func cacheSet(key,val string){
	if len(cache) + 1 > maxCachesize {
		return
	}

	cache[key] = val
}

func cacheGet(key string) string{
	return cache[key]
}

func getBook(isbn string) string{
	return cacheGet(CacheKeyBook + isbn)
}

func setBook(isbn string, name string){
	cacheSet(CacheKeyBook + isbn, name)
}

func getCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func setCD(sku string, title string){
	cacheSet(CacheKeyCD + sku, title)
}


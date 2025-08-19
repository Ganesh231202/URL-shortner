package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

type URL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURl    string    `json:"short_url"`
	CreatinDate time.Time `json:"create_date"`
}

var urlDB = make(map[string]URL)

func createShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("hasher", hasher)
	return "www.example.com/some/long/url"
}
func main() {
	fmt.Println("Url Shortener Service")
	OrininalURL := "https://www.example.com/some/long/url"
	createShortURL(OrininalURL)
}

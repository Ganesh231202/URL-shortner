package main

import (
	"fmt"
	"time"
)

type URL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURl    string    `json:"short_url"`
	CreatinDate time.Time `json:"create_date"`
}

func main() {
	fmt.Println("Url Shortener Service")
}

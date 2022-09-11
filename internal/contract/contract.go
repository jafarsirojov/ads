package contract

import "time"

type Ad struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Price           int       `json:"price"`
	LinkToMainPhoto []string  `json:"linkToMainPhoto"`
	LinksToPhotos   []string  `json:"linksToPhotos"`
	CreatedAt       time.Time `json:"createdAt"`
}

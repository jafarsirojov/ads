package contract

import "time"

type Ad struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Price         int       `json:"price"`
	LinksToPhotos []string  `json:"linksToPhotos"`
	CreatedAt     time.Time `json:"createdAt"`
}

type AdFromList struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Price       int       `json:"price"`
	LinkToPhoto string    `json:"linkToPhoto"`
	CreatedAt   time.Time `json:"createdAt"`
}

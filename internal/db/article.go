package db

import "time"

type NewsAritcle struct {
	Source      string
	Title       string
	Authors     []string
	Content     string
	AISummary   string
	Url         string
	ImgUrl      string
	PublishedAt time.Time
}

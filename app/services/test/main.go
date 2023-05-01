package main

import (
	"RssReader/bussiness/data/store/rss"
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type PQ struct {
	gorm.Model
	Link        string         `gorm:"link" json:"link"`
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Content     sql.NullString `gorm:"type:text" json:"content"`
	FeedID      uint           `json:"feed_id"`
	PublishedAt sql.NullTime   `json:"published_at"`
}

func (p PQ) Encode() ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(p)
	if err != nil {
		return nil, err
	}
	fmt.Println("uncompressed size (bytes): ", len(b.Bytes()))
	return b.Bytes(), nil
}
func main() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(PQ{
		Link:        "2",
		Title:       "32",
		Author:      "4",
		Content:     sql.NullString{},
		FeedID:      0,
		PublishedAt: sql.NullTime{},
	})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// HERE ARE YOUR BYTES!!!!
	fmt.Println(network.Bytes())

	// Decode (receive) the value.
	var q rss.Post
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println(q.Link)
}

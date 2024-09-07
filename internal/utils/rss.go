package utils

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}

	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}
	return rssFeed, nil
}

func RSSFeedScrapper(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(
			context.Background(),
			int32(concurrency),
		)

		if err != nil {
			log.Printf("error fetching feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched: ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error fetching feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {

		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("error parsing date %v: %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubAt,
			Url:         item.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("error creating post: %v", err)
		}
	}
	log.Printf("feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

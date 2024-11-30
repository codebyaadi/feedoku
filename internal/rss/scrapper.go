package rss

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
	
	"github.com/google/uuid"
	"go.uber.org/zap"
	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/codebyaadi/rss-feed-agg/internal/models"
	"github.com/codebyaadi/rss-feed-agg/internal/utils"
	"github.com/codebyaadi/rss-feed-agg/pkg/logger"
)

func urlToFeed(url string) (models.RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return models.RSSFeed{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return models.RSSFeed{}, err
	}

	rssFeed := models.RSSFeed{}

	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return models.RSSFeed{}, err
	}
	return rssFeed, nil
}

func RSSFeedScrapper(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	logger.Logger.Info(
		fmt.Sprintf("scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest),
	)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(
			context.Background(),
			int32(concurrency),
		)

		if err != nil {
			logger.Logger.Error("error fetching feeds:", zap.Error(err))
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
		logger.Logger.Error("error marking feed as fetched: ", zap.Error(err))
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		logger.Logger.Error("error fetching feed: ", zap.Error(err))
		return
	}

	for _, item := range rssFeed.Channel.Item {

		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := utils.ParseDate(item.PubDate)
		if err != nil {
			logger.Logger.Error(
				fmt.Sprintf("error parsing date %v:", item.PubDate),
				zap.Error(err),
			)
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
			logger.Logger.Error("error creating post: ", zap.Error(err))
		}
	}
	logger.Logger.Info(
		fmt.Sprintf(
			"feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item),
		),
	)
}

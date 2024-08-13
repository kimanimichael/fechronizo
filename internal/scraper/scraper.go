package scraper

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mike-kimani/fechronizo/internal/database"
	"github.com/mike-kimani/fechronizo/internal/rss"
)

func StartScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v go routines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("couldn't get feeds to fetch", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go ScrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}

func ScrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	rssFeed, err := rss.UrlToFeed(feed.Url)
	if err != nil {
		log.Println("couldn't convert feed to url", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		// log.Println("Found post ", item.Title, " on feed ", feed.Name)
		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			log.Printf("couldn't parse date  %v with error %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePosts(context.Background(), database.CreatePostsParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
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
			log.Println("failed to create post", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found \n", feed.Name, len(rssFeed.Channel.Item))

	_, err = db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("couldn't mark feed as fetched", err)
		return
	}
}

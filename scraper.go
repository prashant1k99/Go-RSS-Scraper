package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/prashant1k99/Go-RSS-Scraper/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenReqest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration \n", concurrency, timeBetweenReqest)

	ticker := time.NewTicker(timeBetweenReqest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("Error fetching feeds:", err)
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

	log.Printf("Scraping feed: %s \n", feed.Url)
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		log.Printf("Found post %s on feed %s\n", item.Title, feed.Name)
	}
	log.Printf("Feed %s collected, %v items found \n", feed.Url, len(rssFeed.Channel.Items))
}

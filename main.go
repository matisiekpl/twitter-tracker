package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"strconv"
)

var scraper *twitterscraper.Scraper
var latestScraper *twitterscraper.Scraper
var size int
var requests int

func handleTweetSearch(c *fiber.Ctx) error {
	query := c.Query("q")
	num := c.QueryInt("n", 10)
	logrus.Info("Searching for tweets with query: ", query)
	var results []*twitterscraper.TweetResult
	tweets := scraper.SearchTweets(context.Background(), url.QueryEscape(query), num/2)
	for tweet := range tweets {
		results = append(results, tweet)
	}
	latestTweets := latestScraper.SearchTweets(context.Background(), url.QueryEscape(query), num/2)
	for tweet := range latestTweets {
		results = append(results, tweet)
	}
	size += len(results)
	requests++
	return c.JSON(results)
}

func main() {
	scraper = twitterscraper.New()
	err := scraper.LoginOpenAccount()
	if err != nil {
		logrus.Error(err)
		return
	}
	latestScraper = twitterscraper.New()
	err = scraper.LoginOpenAccount()
	if err != nil {
		logrus.Error(err)
		return
	}
	scraper.SetSearchMode(twitterscraper.SearchLatest)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"size":     strconv.Itoa(size),
			"requests": strconv.Itoa(requests),
			"version":  "v1.0.0",
		})
	})
	app.Get("/search", handleTweetSearch)
	app.Use(logger.New())
	app.Get("/metrics", monitor.New())
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	logrus.Error(app.Listen(":" + port))
}

package rss_feed

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed

	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}

	feed = *unescapeHtml(&feed)

	return &feed, nil
}

func unescapeHtml(feed *RSSFeed) *RSSFeed {
	for i, item := range feed.Channel.Items {
		feed.Channel.Items[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Items[i].Description = html.UnescapeString(item.Description)
	}

	return feed
}

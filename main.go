package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/nlopes/slack"
)

// settings of slack
const (
	WEBHOOKURL = "[my webhookurl]"
)

// GetPage function
func GetPage(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("#topics_list4 li").Each(func(_ int, s *goquery.Selection) {
		imageURL, _ := s.Find("img").Attr("src")
		url, _ := s.Find(".pic a").Attr("href")
		shopName := s.Find(".shop-nm .name").Text()
		text := url + "\n" + shopName

		PostSlack(text, imageURL)
	})
}

// PostSlack function
func PostSlack(text, imageURL string) {
	payload := &slack.WebhookMessage{
		Attachments: []slack.Attachment{
			{
				Text:     text,
				ImageURL: imageURL,
			},
		},
	}

	slack.PostWebhook(WEBHOOKURL, payload)
}

func main() {
	url := "https://pets-kojima.com/small_list/?topics_group_id=4&group=2&order_type=2"
	GetPage(url)
}

package wpress

import "fmt"

func NewClient(baseUrl, versionApi, consumerKey, consumerSecret string) *Client {
	c := &Client{
		BaseUrl: baseUrl,
		ConsumerKey: consumerKey,
		ConsumerSecret: consumerSecret,
	}

	c.apiUrl = fmt.Sprintf("%s/wp-json/wc/%s", baseUrl, versionApi)

	c.Posts = (*PostService)(c)

	return c
}
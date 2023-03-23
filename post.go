package wpress

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type PostService Client

type Category struct {
	Description string `json:"description"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type MetaTag struct {
	Key string `json:"meta_key"`
	Value string `json:"meta_value"`
}

type Post struct {
	ID string `json:"id"`
	Link string `json:"link"`
	Slug string `json:"slug"`
	Status string `json:"status"`
	PostType string `json:"type"`
	GeneratedSlug string `json:"generated_slug"`
	Title string `json:"title"`
	Content string `json:"content"`
	Categories []Category `json:"categories"`
	Meta []MetaTag `json:"meta"`
}

func (s *PostService) List() ([]Post, error) {
	return nil, nil
}

func (s *PostService) ListByPage(page, per_page int) ([]Post, error) {
	if page <= 0 {
		page = 1
	}

	if per_page <= 0 {
		per_page = 1
	}

	apiUrl := fmt.Sprintf("%s/posts", s.apiUrl)

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return []Post{}, err
	}

	req.SetBasicAuth(s.ConsumerKey, s.ConsumerSecret)

	q := url.Values{}
	q.Add("page", fmt.Sprintf("%d", page))
	q.Add("per_page", fmt.Sprintf("%d", per_page))

	req.URL.RawQuery = q.Encode()

	return s.Get(req)
}

func (s *PostService) Get(request *http.Request) ([]Post, error) {
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Post{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []Post{}, err
	}

	var posts []Post
	if err := json.Unmarshal(body, &posts); err != nil {
		return []Post{}, err
	}

	return posts, nil
}
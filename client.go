package wpress

type Client struct {
	BaseUrl        string
	apiUrl         string
	VersionApi     string
	ConsumerKey    string
	ConsumerSecret string

	Posts *PostService
}
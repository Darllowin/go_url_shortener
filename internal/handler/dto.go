package handler

import (
	"errors"
	"net/url"
	"time"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

type StatsResponse struct {
	Original  string    `json:"original"`
	CreatedAt time.Time `json:"createdAt"`
	Visits    int       `json:"visits"`
}

func (r ShortenRequest) Validate() error {
	if r.URL == "" {
		return errors.New("url is required")
	}

	// parse string as URL
	u, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return errors.New("invalid url format")
	}
	// checking for the presence of a scheme (http or https)
	if u.Scheme == "" || u.Host == "" {
		return errors.New("url must include scheme (http/https) and host")
	}

	return nil
}

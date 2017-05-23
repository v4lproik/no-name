package client

import (
	"net/http"
	"net/url"
)

type WebClient interface {
	Scrap() (*http.Response, error)
	ScrapWithParameter(path string, method string, values url.Values) (*http.Response, error)
	ScrapWithNoParameter(path string, method string) (*http.Response, error)
	CraftUrl(path string) (string)
	GetUrl() (*url.URL)
}
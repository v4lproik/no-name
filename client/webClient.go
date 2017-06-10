package client

import (
	"net/http"
	"net/url"
)

type WebClient interface {
	Scrap() (*http.Response, error)
	ScrapWithParameter(path string, method string, values url.Values) (*http.Response, error)
	ScrapWithNoParameter(path string, method string) (*http.Response, error)
	BasicAuth(path string, method string, username string, password string) (*http.Response, error)

	GetDomain() (*url.URL)
	GetDomainHttpCode() (int)

	CraftUrlGet(path string, values url.Values) (string)
	CraftUrlPost(path string) (string)
}
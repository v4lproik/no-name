package client

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/juju/loggo"
	"net/http"
	"net/url"
	"net/http/cookiejar"
)

type Web struct {
	client *http.Client
	Url *url.URL
}

var loggerWeb = loggo.GetLogger("web")

func NewWeb(ip string) (*Web){
	client := http.DefaultClient
	cookieJar, _ := cookiejar.New(nil)
	client.Jar = cookieJar

	url, err := url.Parse(ip)
	if err != nil {
		panic(err)
	}

	if url.Scheme == "" {
		loggerWeb.Warningf("No scheme for url " + ip + ". Setting scheme to http://" + ip)
		url.Scheme = "http"
	}

	return &Web{client, url}
}

func (w *Web) Scrap() (*http.Response, error){
	scheme := w.Url.Scheme
	host := w.Url.Path

	res, err := w.client.Get(scheme + "://" + host)
	if err != nil {
		return nil, err
	}

	//output, err := httputil.DumpResponse(res, true)
	//loggerWeb.Debugf(string(output))

	return res, nil
}

func (w *Web) GetDocument(res *http.Response) (*goquery.Document, error){
	return goquery.NewDocumentFromResponse(res)
}

func (w *Web) ScrapWithParameter(path string, method string, values url.Values) (*goquery.Document, error){
	scheme := w.Url.Scheme
	host := w.Url.Path

	if method == "POST" || method == "post"{
		urlToRequest := scheme + "://" + host + "/" + path
		res, err := w.client.PostForm(urlToRequest, values)
		if err != nil {
			return nil, err
		}

		//output, err := httputil.DumpResponse(res, true)
		//loggerWeb.Debugf(string(output))

		return goquery.NewDocumentFromResponse(res)
	}else {
		if method == "GET" || method == "get" {

			// craft url
			urlToRequest := scheme + "://" + host + "/" + path + "?" + values.Encode()

			// submit form
			res, err := w.client.Get(urlToRequest)
			if err != nil {
				return nil, err
			}
			return goquery.NewDocumentFromResponse(res)
		}
	}

	return nil, nil
}

// setter for client
func (w *Web) SetCookieJar(cookies []*http.Cookie){
	w.client.Jar.SetCookies(w.Url,cookies)
}
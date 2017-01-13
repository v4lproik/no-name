package client

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/juju/loggo"
	"net/http"
	"net/url"
	"net/http/cookiejar"
	"strings"
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

	if !strings.HasPrefix(ip, "http://") || !strings.HasPrefix(ip, "https://") {
		loggerWeb.Warningf("No scheme for url " + ip + ". Setting scheme to http://" + ip)
		ip = "http://" + ip
	}

	url, err := url.Parse(ip)
	if err != nil {
		panic(err)
	}

	return &Web{client, url}
}

func (w *Web) Scrap() (*http.Response, error){
	scheme := w.Url.Scheme
	host := w.Url.Host
	path := w.Url.Path

	res, err := w.client.Get(scheme + "://" + host + path)
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

func (w *Web) ScrapWithParameter(path string, method string, values url.Values) (*http.Response, error){
	scheme := w.Url.Scheme
	host := w.Url.Host

	if method == "POST" || method == "post"{

		urlToRequest := ""
		if strings.HasPrefix(path, "/") {
			urlToRequest = scheme + "://" + host + path
		}else{
			urlToRequest = scheme + "://" + host + "/" + path
		}

		res, err := w.client.PostForm(urlToRequest, values)
		if err != nil {
			return nil, err
		}
		//print(res.Request)
		//output, err := httputil.DumpRequest(res.Request, true)
		//loggerWeb.Debugf(string(output))

		return res, nil
	}else {
		if method == "GET" || method == "get" {

			// craft url
			urlToRequest := ""
			if strings.HasPrefix("/", path) {
				urlToRequest = scheme + "://" + host + path + "?" + values.Encode()
			}else{
				urlToRequest = scheme + "://" + host + "/" + path + "?" + values.Encode()
			}

			// submit form
			res, err := w.client.Get(urlToRequest)
			if err != nil {
				return nil, err
			}
			return res, nil
		}
	}

	return nil, nil
}

func (w *Web) CraftUrl(path string) (string){
	url, err := url.Parse(path)
	if err != nil {
		loggerWeb.Errorf(err.Error())
	}

	if url.Host == "" {
		scheme := w.Url.Scheme
		host := w.Url.Path

		return scheme + "://" + host + "/" + path
	}

	return path
}

// setter for client
func (w *Web) SetCookieJar(cookies []*http.Cookie){
	w.client.Jar.SetCookies(w.Url,cookies)
}
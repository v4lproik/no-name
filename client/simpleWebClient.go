package client

import (
	"github.com/juju/loggo"
	"net/http"
	"net/url"
	"net/http/cookiejar"
	"strings"
)

type simpleWebClient struct {
	client *http.Client
	url    *url.URL
}

var loggerWeb = loggo.GetLogger("web")

func NewSimpleWebClient(ip string) (*simpleWebClient){
	client := http.DefaultClient
	cookieJar, _ := cookiejar.New(nil)
	client.Jar = cookieJar

	if !strings.HasPrefix(ip, "http://") && !strings.HasPrefix(ip, "https://") {
		loggerWeb.Warningf("No scheme for url " + ip + ". Setting scheme to http://" + ip)
		ip = "http://" + ip
	}

	url, err := url.Parse(ip)
	if err != nil {
		panic(err)
	}

	return &simpleWebClient{client, url}
}

func (w *simpleWebClient) Scrap() (*http.Response, error){
	scheme := w.url.Scheme
	host := w.url.Host
	path := w.url.Path

	res, err := w.client.Get(scheme + "://" + host + path)
	if err != nil {
		return nil, err
	}

	//output, err := httputil.DumpResponse(res, true)
	//loggerWeb.Debugf(string(output))

	return res, nil
}

func (w *simpleWebClient) ScrapWithParameter(path string, method string, values url.Values) (*http.Response, error){
	scheme := w.url.Scheme
	host := w.url.Host

	if method == "POST" || method == "post"{

		urlToRequest := ""
		if strings.HasPrefix(path, "/") {
			urlToRequest = scheme + "://" + host + path
		}else{
			if  strings.HasPrefix(path,"http"){
				urlToRequest = path
			}else{
				urlToRequest = scheme + "://" + host + "/" + path
			}
		}

		res, err := w.client.PostForm(urlToRequest, values)
		if err != nil {
			return nil, err
		}
		//fmt.Println(values)
		//fmt.Println(res.Request)
		//output, err := httputil.DumpRequest(res.Request, true)
		//fmt.Println(string(output))

		return res, nil
	}else {
		if method == "GET" || method == "get" {

			// craft url
			urlToRequest := ""
			if strings.HasPrefix("/", path) {
				urlToRequest = scheme + "://" + host + path + "?" + values.Encode()
			}else{
				if  strings.HasPrefix(path,"http"){
					urlToRequest = path + "?" + values.Encode()
				}else{
					urlToRequest = scheme + "://" + host + "/" + path + "?" + values.Encode()
				}
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

func (w *simpleWebClient) ScrapWithNoParameter(path string, method string) (*http.Response, error){
	scheme := w.url.Scheme
	host := w.url.Host

	urlToRequest := ""
	if strings.HasPrefix(path, "/") {
		urlToRequest = scheme + "://" + host + path
	}else{
		if  strings.HasPrefix(path,"http"){
			urlToRequest = path
		}else{
			urlToRequest = scheme + "://" + host + "/" + path
		}
	}

	if method == "POST" || method == "post"{

		res, err := w.client.PostForm(urlToRequest, url.Values{})
		if err != nil {
			return nil, err
		}
		//print(res.Request)
		//output, err := httputil.DumpRequest(res.Request, true)
		//loggerWeb.Debugf(string(output))

		return res, nil
	}else {
		if method == "GET" || method == "get" {

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

// TODO: refacto
func (w *simpleWebClient) CraftUrl(path string) (string){
	url, err := url.Parse(path)
	if err != nil {
		loggerWeb.Errorf(err.Error())
	}

	if url.Host == "" {
		scheme := w.url.Scheme
		host := w.url.Path

		return scheme + "://" + host + "/" + path
	}

	return path
}

func (w *simpleWebClient) GetUrl() (*url.URL) {
	return w.url
}
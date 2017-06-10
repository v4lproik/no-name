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

	domain    *url.URL
	domainResponseCode int
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

	return &simpleWebClient{client, url, 0}
}

func (w *simpleWebClient) Scrap() (*http.Response, error){
	res, err := w.client.Get(w.CraftUrlGet(w.domain.Path, url.Values{}))
	if res != nil {
		w.domainResponseCode = res.StatusCode
	}
	if err != nil {
		return nil, err
	}

	//output, err := httputil.DumpResponse(res, true)
	//loggerWeb.Debugf(string(output))

	return res, nil
}

func (w *simpleWebClient) ScrapWithParameter(path string, method string, values url.Values) (*http.Response, error){

	switch {
	case method == "POST" || method == "post" :
		res, err := w.client.PostForm(w.CraftUrlPost(path), values)
		if err != nil {
			return nil, err
		}
		return res, nil
	case method == "GET" || method == "get" :
		res, err := w.client.Get(w.CraftUrlGet(path, values))
		if err != nil {
			return nil, err
		}

		return res, nil
	default:
		loggerWeb.Criticalf("Method " + method + "does not exist.")
		return nil, nil
	}
}

func (w *simpleWebClient) ScrapWithNoParameter(path string, method string) (*http.Response, error){

	switch {
	case method == "POST" || method == "post" :
		res, err := w.client.PostForm(w.CraftUrlPost(path), url.Values{})
		if err != nil {
			return nil, err
		}
		return res, nil
	case method == "GET" || method == "get" :
		res, err := w.client.PostForm(w.CraftUrlPost(path), url.Values{})
		if err != nil {
			return nil, err
		}

		return res, nil
	default:
		loggerWeb.Criticalf("Method " + method + "does not exist.")
		return nil, nil
	}
}

func (w *simpleWebClient) BasicAuth(path string, method string, username string, password string) (*http.Response, error){

	if method != "POST" && method != "post" && method != "GET" && method != "get" {
		loggerWeb.Criticalf("Method " + method + "does not exist.")
		return nil, nil
	}

	request, err := http.NewRequest(method, w.CraftUrlPost(path), strings.NewReader(""))
	request.SetBasicAuth(username, password)
	res, err := w.client.Do(request)
	if err != nil {
		loggerWeb.Criticalf("Error occured while submitting basic auth request.", err.Error())
		return nil, err
	}
	return res, nil
}


func (w *simpleWebClient) CraftUrlPost(path string) (string){
	scheme := w.domain.Scheme
	host := w.domain.Host
	urlToRequest := ""

	// host never has "/" in the end
	// therefore always trim and add manually "/"
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}

	if  strings.HasPrefix(path,"http"){
		urlToRequest = path
	}else{
		if path != "" {
			urlToRequest = scheme + "://" + host + "/" + path
		}else{
			urlToRequest = scheme + "://" + host
		}
	}

	return urlToRequest
}

func (w *simpleWebClient) CraftUrlGet(path string, values url.Values) (string){
	scheme := w.domain.Scheme
	host := w.domain.Host
	urlToRequest := ""

	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}

	if len(values) > 0 {
		if strings.HasSuffix(path, "?") {
			path = strings.TrimSuffix(path, "?")
		}

		if strings.HasPrefix("/", path) {
			urlToRequest = scheme + "://" + host + path + "?" + values.Encode()
		}else{
			if  strings.HasPrefix(path,"http"){
				urlToRequest = path + "?" + values.Encode()
			}else{
				urlToRequest = scheme + "://" + host + "/" + path + "?" + values.Encode()
			}
		}
	}else{
		if strings.HasPrefix("/", path) {
			urlToRequest = scheme + "://" + host + path
		}else{
			if  strings.HasPrefix(path,"http"){
				urlToRequest = path
			}else{
				urlToRequest = scheme + "://" + host + "/" + path
			}
		}
	}


	return urlToRequest
}

func (w *simpleWebClient) GetDomain() (*url.URL) {
	return w.domain
}

func (w *simpleWebClient) GetDomainHttpCode() (int) {
	return w.domainResponseCode
}
package module

import (
	"net/url"
	"strings"
	"errors"
	"net/http/httptest"
	"bytes"
	"github.com/v4lproik/no-name/data"
	"net/http"
	"github.com/v4lproik/no-name/client"
	"os"
	"io/ioutil"
)

const HTML_TAGS_NAMES_TEST = "conf/html-detection-tags_test.txt"
const LOGIN_TEST = "conf/login_test.txt"
const PASSWORD_TEST = "conf/password_test.txt"
const DEFAULT_PASSWORD_TEST = "conf/default-password-web-interface_test.txt"
var CWD, _ = os.Getwd()
var _ client.WebClient = (*fakeWebClient)(nil)

var bytesDocWithoutForm, _ = ioutil.ReadAll(strings.NewReader("<html><title>Scientists Stored These Images in DNA—Then Flawlessly Retrieved Them</title></html>"))
var bytesDocWithFormWithCsrf, _ = ioutil.ReadAll(strings.NewReader(`"<html><form action="url_to_submit" method="POST"><input type="text" name="username" /><input type="password" name="password"><input type="text" name="otherinput" value="random"/><input type="hidden" name="user_token" value="csrftoken" /></form></html>"`))
var bytesDocWithFormWithoutCsrf, _ = ioutil.ReadAll(strings.NewReader(`"<html><form action="url_to_submit" method="POST"><input type="text" name="username" /><input type="password" name="password"><input type="text" name="otherinput" value="random"/></form></html>"`))
var bytesDocWithFormWithoutCsrfWithoutGoodCred, _ = ioutil.ReadAll(strings.NewReader(`"<html>ERROR LOGIN ! The password or the login is \nnot valid... Please \ncheck your credentials !\n<form action="url_to_submit" method="POST"><input type="text" name="username" /><input type="password" name="password"><input type="text" name="otherinput" value="random"/></form></html>"`))
var bytesDocWithFormWithoutCsrfWithGoodCred, _ = ioutil.ReadAll(strings.NewReader(`"<html>Welcome admin !\n <div>\nYour are now logged on to the super admin page !<br /> Do not give your credentials to anyone else...\n step mother included :)</div></html>"`))


//// MOCKING ////
// MODULE //
type fakeNextModule struct{
	count int
}
func (m *fakeNextModule) Request(flag bool, wi *data.WebInterface) {
	m.count += 1
}
func (m *fakeNextModule) SetNextModule(next Module){}
// CLIENT //
type fakeWebClient struct{
	client *http.Client
	url    *url.URL

	CountScrapWithParameter int
}
func NewFakeWebClient(ip string) (*fakeWebClient){
	if !strings.HasPrefix(ip, "http://") && !strings.HasPrefix(ip, "https://") {
		ip = "http://" + ip
	}

	url, err := url.Parse(ip)
	if err != nil {
		panic(err)
	}

	return &fakeWebClient{nil, url, 0}
}

func (w *fakeWebClient) Scrap() (*http.Response, error){
	switch w.url.String() {
	case "http://127.0.0.0.0.a":
		return nil, error(errors.New("Bad ip"))
	case "http://127.0.0.1":
		rw := httptest.NewRecorder()
		rw.Header().Set("Content-Type", "text/html")
		rw.Code = 200
		rw.Body = bytes.NewBuffer(bytesDocWithoutForm)
		httpReponse := rw.Result()
		httpReponse.Request = httptest.NewRequest("GET", "http://127.0.0.1", strings.NewReader("request"))

		return httpReponse, nil
	case "http://127.0.0.2":
		rw := httptest.NewRecorder()
		rw.Header().Set("Content-Type", "text/html")
		rw.Code = 200
		rw.Body = bytes.NewBuffer(bytesDocWithFormWithCsrf)
		httpReponse := rw.Result()
		httpReponse.Request = httptest.NewRequest("GET", "http://127.0.0.2", strings.NewReader("request"))

		return httpReponse, nil

	case "http://127.0.0.3":
		rw := httptest.NewRecorder()
		rw.Header().Set("Content-Type", "text/html")
		rw.Code = 200
		rw.Body = bytes.NewBuffer(bytesDocWithFormWithoutCsrf)
		httpReponse := rw.Result()
		httpReponse.Request = httptest.NewRequest("GET", "http://127.0.0.3", strings.NewReader("request"))

		return httpReponse, nil
	default:
		return nil, nil
	}

}

func (w *fakeWebClient) ScrapWithParameter(path string, method string, values url.Values) (*http.Response, error){

	switch {
	case w.url.String() == "http://127.0.0.3" && path == "url_to_submit" && values.Get("username") == "bug" && values.Get("password") == "admin":
		rw := httptest.NewRecorder()
		rw.Header().Set("Content-Type", "text/html")
		rw.Code = 200
		rw.Body = bytes.NewBuffer(bytesDocWithFormWithoutCsrfWithGoodCred)
		httpReponse := rw.Result()
		httpReponse.Request = httptest.NewRequest("GET", "http://127.0.0.3", strings.NewReader("request"))

		return httpReponse, nil
	case w.url.String() == "http://127.0.0.3" && path == "url_to_submit":
		rw := httptest.NewRecorder()
		rw.Header().Set("Content-Type", "text/html")
		rw.Code = 200
		rw.Body = bytes.NewBuffer(bytesDocWithFormWithoutCsrfWithoutGoodCred)
		httpReponse := rw.Result()
		httpReponse.Request = httptest.NewRequest("GET", "http://127.0.0.3", strings.NewReader("request"))

		return httpReponse, nil
	default:
		return nil, nil
	}

	w.CountScrapWithParameter += 1
	return nil, nil
}

func (w *fakeWebClient) ScrapWithNoParameter(path string, method string) (*http.Response, error){return nil, nil}
func (w *fakeWebClient) CraftUrlGet(path string, values url.Values) (string){return ""}
func (w *fakeWebClient) CraftUrlPost(path string) (string){return ""}
func (w *fakeWebClient) GetUrl() (*url.URL) {return w.url}

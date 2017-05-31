package module

import (
	"testing"
	"github.com/v4lproik/no-name/data"
	"github.com/v4lproik/no-name/client"
	"net/http"
	"net/url"
	"strings"
	"errors"
	"net/http/httptest"
	"io/ioutil"
	"bytes"
	"github.com/stretchr/testify/assert"
)

var _ client.WebClient = (*fakeWebClient)(nil)
var _ Module = (*fakeNextModule)(nil)
var bytesDoc, _ = ioutil.ReadAll(strings.NewReader("<html><title>Scientists Stored These Images in DNA—Then Flawlessly Retrieved Them</title></html>"))

func TestNewScrapModule_withNonReachableUrl_shouldReturnError_and_stopTheChain(t *testing.T) {
	t.Log("Call ScrapModule with non-reachable url should return an error & stop the chain")

	//given
	sm := scrapModule{}
	wi := data.NewWebInterface(NewFakeWebClient("127.0.0.0.0.a"))
	fakeNextModule:= &fakeNextModule{0}
	sm.SetNextModule(fakeNextModule)

	// when
	sm.Request(true, wi)

	// then
	assert.Equal(t, fakeNextModule.count, 0)
}

func TestNewScrapModule_withNonParsableResponse_shouldReturnError_and_stopTheChain(t *testing.T) {
	t.Log("Call ScrapModule with non-parsable response should return an error & stop the chain")

	//given
	sm := scrapModule{}
	wi := data.NewWebInterface(NewFakeWebClient("127.0.0.12"))
	fakeNextModule:= &fakeNextModule{0}
	sm.SetNextModule(fakeNextModule)

	// when
	sm.Request(true, wi)

	// then
	assert.Equal(t, fakeNextModule.count, 0)
}

func TestNewScrapModule_withParsableResponse_shouldContinueTheChain(t *testing.T) {
	t.Log("Call ScrapModule with parsable response should continue the chain")

	//given
	sm := scrapModule{}
	wi := data.NewWebInterface(NewFakeWebClient("127.0.0.1"))
	fakeNextModule:= &fakeNextModule{0}
	sm.SetNextModule(fakeNextModule)

	// when
	sm.Request(true, wi)

	// then
	assert.Equal(t, fakeNextModule.count, 1)
}

//// MOCKING ////
// MODULE //
type fakeNextModule struct{
	count int
}
func (m *fakeNextModule) Request(flag bool, wi *data.WebInterface) {
	m.count = 1
}
func (m *fakeNextModule) SetNextModule(next Module){}
// CLIENT //
type fakeWebClient struct{
	client *http.Client
	url    *url.URL
}
func NewFakeWebClient(ip string) (*fakeWebClient){
	if !strings.HasPrefix(ip, "http://") && !strings.HasPrefix(ip, "https://") {
		ip = "http://" + ip
	}

	url, err := url.Parse(ip)
	if err != nil {
		panic(err)
	}

	return &fakeWebClient{nil, url}
}

func (w *fakeWebClient) Scrap() (*http.Response, error){
	switch w.url.String() {
	case "http://127.0.0.0.0.a":
		return nil, error(errors.New("Bad ip"))
	case "http://127.0.0.1":
		rw := httptest.NewRecorder()
		rw.Header().Set("Content-Type", "text/html")
		rw.Code = 200
		rw.Body = bytes.NewBuffer(bytesDoc)
		httpReponse := rw.Result()
		httpReponse.Request = httptest.NewRequest("GET", "http://127.0.0.1", strings.NewReader("<html><title>Scientists Stored These Images in DNA—Then Flawlessly Retrieved Them</title></html>"))

		return httpReponse, nil
	default:
		return nil, nil
	}

}

func (w *fakeWebClient) ScrapWithParameter(path string, method string, values url.Values) (*http.Response, error){return nil, nil}
func (w *fakeWebClient) ScrapWithNoParameter(path string, method string) (*http.Response, error){return nil, nil}
func (w *fakeWebClient) CraftUrlGet(path string, values url.Values) (string){return ""}
func (w *fakeWebClient) CraftUrlPost(path string) (string){return ""}
func (w *fakeWebClient) GetUrl() (*url.URL) {return w.url}
package client

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/url"
	"net/http"
	"github.com/jarcoal/httpmock"
	"errors"
)

var HTTPCLIENT = http.DefaultClient

func TestSimpleWebClientWithoutHttpClientShouldReturnError(t *testing.T) {
	t.Log("Call NewSimpleWebClient without http client should panic")

	//given
	domain := "myurl.com"

	// then
	assert.PanicsWithValue(t, "HttpClient should be initiated", func(){NewSimpleWebClient(domain, nil)})
}

func TestSimpleWebClientWithoutSchemeShouldSetOne(t *testing.T) {
	t.Log("Call NewSimpleWebClient without scheme should set the 'http' scheme")

	//given
	domain := "myurl.com"


	// when
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// then
	assert.Equal(t, "http://myurl.com", simpleWebClient.GetDomain().String(), "The craft url is not the one expected")
}

// domain:myurl.com+path:=http://myurl.com
func TestScrapWithValidDomain(t *testing.T) {
	t.Log("Call scrap with a valid domain should return http response")

	//given
	domain := "myurl.com"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://myurl.com",
		httpmock.NewStringResponder(200, ""))

	res, _ := simpleWebClient.Scrap()

	httpmock.Deactivate()


	// then
	assert.NotNil(t, res)
	assert.Equal(t, 200, simpleWebClient.domainResponseCode)
}

func TestScrapWithNotValidDomain(t *testing.T) {
	t.Log("Call scrap with not a valid domain should return nil http response")

	//given
	domain := "myurl.com"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://myurl.com",
		httpmock.NewErrorResponder(errors.New("connection refused")))

	res, err := simpleWebClient.Scrap()

	httpmock.Deactivate()


	// then
	assert.Nil(t, res)
	assert.Error(t, err)
}

// domain:myurl.com+path:=http://myurl.com
func TestPostNewCraftUrlWithoutSchemeShouldReturnUrlWithScheme(t *testing.T) {
	t.Log("Call CraftUrlPost without scheme should return url with scheme")

	//given
	domain := "myurl.com"
	path := ""
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	craftUrl := simpleWebClient.CraftUrlPost(path)

	// then
	assert.Equal(t, "http://myurl.com" + path, craftUrl, "The craft url is not the one expected")
}

// domain:myurl.com+path:http://myurl.com/=http://myurl.com/
func TestPostNewCraftUrlWithPathEqualsDomainShouldReturnDomainUrl(t *testing.T) {
	t.Log("Call CraftUrlPost with path equals to domain should return domain")

	//given
	domain := "myurl.com"
	path := "http://myurl.com/"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	craftUrl := simpleWebClient.CraftUrlPost(path)

	// then
	assert.Equal(t, path, craftUrl, "The craft url is not the one expected")
}

// domain:myurl.com+path:/=http://myurl.com
func TestPostNewCraftUrlWithSlashShouldReturnUrlWithSchemeAndSlash(t *testing.T) {
	t.Log("Call CraftUrlPost with slash should return domain")

	//given
	domain := "myurl.com"
	path := "/"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	craftUrl := simpleWebClient.CraftUrlPost(path)

	// then
	assert.Equal(t, "http://" + domain, craftUrl, "The craft url is not the one expected")
}

func TestNewCraftUrlWithoutSchemeShouldReturnUrlWithScheme(t *testing.T) {
	t.Log("Call CraftUrlGet without scheme should return url with scheme")

	//given
	domain := "myurl.com"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	craftUrl := simpleWebClient.CraftUrlGet("", url.Values{})

	// then
	assert.Equal(t, "http://myurl.com", craftUrl, "The craft url is not the one expected")
}

func TestNewCraftUrlWithSchemeShouldReturnUrlWithScheme(t *testing.T) {
	t.Log("Call CraftUrlGet with scheme should return url with scheme")

	//given
	domain := "http://myurl.com"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	craftUrl := simpleWebClient.CraftUrlGet("", url.Values{})

	// then
	assert.Equal(t, "http://myurl.com", craftUrl, "The craft url is not the one expected")
}

func TestNewCraftUrlWithPathShouldReturnUrlWithPath(t *testing.T) {
	t.Log("Call CraftUrlGet with path should return url with path")

	//given
	domain := "mydomain.com"
	path := "mypath"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	craftUrl := simpleWebClient.CraftUrlGet(path, url.Values{})

	// then
	assert.Equal(t, "http://mydomain.com/mypath", craftUrl, "The craft url is not the one expected")
}

func TestNewCraftUrlWithPathAndParameterShouldReturnUrlWithPathAndParameter(t *testing.T) {
	t.Log("Call CraftUrlGet with path and parameter should return url with path and parameter")

	//given
	domain := "mydomain.com"
	path := "mypath"
	parameter := url.Values{"parameter": []string{"parametervalue"}}
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	craftUrl := simpleWebClient.CraftUrlGet(path, parameter)

	// then
	assert.Equal(t, "http://mydomain.com/mypath?parameter=parametervalue", craftUrl, "The craft url is not the one expected")
}

func TestNewCraftUrlWithAQuestionMarkShouldReturnUrlWithOneQuestionMark(t *testing.T) {
	t.Log("Call CraftUrlGet with path with a question mark should return url with one question mark")

	//given
	domain := "mydomain.com"
	path := "mypath?"
	parameter := url.Values{"parameter": []string{"parametervalue"}}
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	craftUrl := simpleWebClient.CraftUrlGet(path, parameter)

	// then
	assert.Equal(t, "http://mydomain.com/mypath?parameter=parametervalue", craftUrl, "The craft url is not the one expected")
}


func TestNewCraftUrlWithUrlSlashAndPathSlashShouldReturnUrlWithOneSlash(t *testing.T) {
	t.Log("Call CraftUrlGet with url slash and path slash should return url with one slash")

	//given
	domain := "mydomain.com/"
	path := "/mypath?"
	parameter := url.Values{"parameter": []string{"parametervalue"}}
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	craftUrl := simpleWebClient.CraftUrlGet(path, parameter)

	// then
	assert.Equal(t, "http://mydomain.com/mypath?parameter=parametervalue", craftUrl, "The craft url is not the one expected")
}

func TestNewBasicAuthWithUnknownMethodShouldReturnError(t *testing.T) {
	t.Log("Call basic auth with unknown method should return error")

	//given
	domain := "mydomain.com/"
	path := "/mypath?"
	method := "ZOP"
	username := "username"
	password := "password"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	_, err := simpleWebClient.BasicAuth(path, method, username, password)

	// then
	assert.Error(t, err, "Method " + method + " does not exist.", "Unknown method should return an error")
}

func TestNewBasicAuthWithNotValidDomainShouldReturnResponse(t *testing.T) {
	t.Log("Call basic auth with not valid domain should return null response")

	//given
	domain := "myurl.com"
	path := "/"
	method := "GET"
	username := "username"
	password := "password"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://myurl.com",
		httpmock.NewErrorResponder(errors.New("connection refused")))

	res, err := simpleWebClient.BasicAuth(path, method, username, password)

	httpmock.Deactivate()


	// then
	assert.Nil(t, res)
	assert.Error(t, err, "connection refused", "Expected error")

}

func TestNewBasicAuthWithValidDomainShouldReturnResponse(t *testing.T) {
	t.Log("Call basic auth with valid domain should return response")

	//given
	domain := "myurl.com"
	path := "/"
	method := "GET"
	username := "username"
	password := "password"
	simpleWebClient := NewSimpleWebClient(domain, HTTPCLIENT)

	// when
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://myurl.com",
		httpmock.NewStringResponder(401, ""))

	res, _ := simpleWebClient.BasicAuth(path, method, username, password)

	httpmock.Deactivate()


	// then
	assert.NotNil(t, res)
}
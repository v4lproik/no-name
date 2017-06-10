package client

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/url"
)

func TestSimpleWebClientWithoutSchemeShouldSetOne(t *testing.T) {
	t.Log("Call NewSimpleWebClient without scheme should set the 'http' scheme")

	//given
	domain := "myurl.com"


	// when
	simpleWebClient := NewSimpleWebClient(domain)

	// then
	assert.Equal(t, "http://myurl.com", simpleWebClient.GetDomain().String(), "The craft url is not the one expected")
}

// domain:myurl.com+path:=http://myurl.com
func TestPostNewCraftUrlWithoutSchemeShouldReturnUrlWithScheme(t *testing.T) {
	t.Log("Call CraftUrlPost without scheme should return url with scheme")

	//given
	domain := "myurl.com"
	path := ""
	simpleWebClient := NewSimpleWebClient(domain)

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
	simpleWebClient := NewSimpleWebClient(domain)

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
	simpleWebClient := NewSimpleWebClient(domain)

	// when
	craftUrl := simpleWebClient.CraftUrlPost(path)

	// then
	assert.Equal(t, "http://" + domain, craftUrl, "The craft url is not the one expected")
}

func TestNewCraftUrlWithoutSchemeShouldReturnUrlWithScheme(t *testing.T) {
	t.Log("Call CraftUrlGet without scheme should return url with scheme")

	//given
	domain := "myurl.com"
	simpleWebClient := NewSimpleWebClient(domain)

	// when
	craftUrl := simpleWebClient.CraftUrlGet("", url.Values{})

	// then
	assert.Equal(t, "http://myurl.com", craftUrl, "The craft url is not the one expected")
}

func TestNewCraftUrlWithSchemeShouldReturnUrlWithScheme(t *testing.T) {
	t.Log("Call CraftUrlGet with scheme should return url with scheme")

	//given
	domain := "http://myurl.com"
	simpleWebClient := NewSimpleWebClient(domain)

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
	simpleWebClient := NewSimpleWebClient(domain)

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
	simpleWebClient := NewSimpleWebClient(domain)

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
	simpleWebClient := NewSimpleWebClient(domain)

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
	simpleWebClient := NewSimpleWebClient(domain)

	// when
	craftUrl := simpleWebClient.CraftUrlGet(path, parameter)

	// then
	assert.Equal(t, "http://mydomain.com/mypath?parameter=parametervalue", craftUrl, "The craft url is not the one expected")
}
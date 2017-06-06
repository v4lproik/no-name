package module

import (
	"testing"
	"github.com/v4lproik/no-name/data"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func TestNewFaviconModuleWithNoFavicon(t *testing.T) {
	t.Log("Call favicon module with no favicon should not try to find information from the default web interfaces")

	//given
	ip := "127.0.0.2"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	// when
	NewFaviconModule(credentials.DefaultWebInterfaces).Request(false, wi)

	// //then
	if wi.Form.FaviconPath != "" {
		t.Errorf("Expected faviconPath as empty")
	}
	if webClient.CountScrapWithParameter > 0 {
		t.Errorf("Expected no call to webclient methods are condition is false")
	}
}

func TestNewFaviconModuleWithUnknownFaviconInDatabase(t *testing.T) {
	t.Log("Call favicon module with unknown favicon in database should not find default credentials")

	//given
	ip := "127.0.0.3"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	// when
	NewFaviconModule(credentials.DefaultWebInterfaces).Request(false, wi)

	// //then
	if wi.Form.FaviconPath != "./favicon.ico" {
		t.Errorf("Expected faviconPath as favicon.ico, not <" + wi.Form.FaviconPath + ">")
	}
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected one call to webclient method scrap with parameter, not count <" + string(webClient.CountScrapWithParameter) + ">")
	}
}

func TestNewFaviconModuleWithKnownFaviconInDatabase(t *testing.T) {
	t.Log("Call favicon module with known favicon in database should find default credentials")

	//given
	ip := "127.0.0.3"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	// when
	NewFaviconModule(credentials.DefaultWebInterfaces).Request(false, wi)

	// //then
	if wi.Form.FaviconPath != "./favicon.ico" {
		t.Errorf("Expected faviconPath as favicon.ico, not <" + wi.Form.FaviconPath + ">")
	}
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected one call to webclient method scrap with parameter, not count <" + string(webClient.CountScrapWithParameter) + ">")
	}
	potentialCredentials := cleanSlice(wi.Form.PotentialCredentials)
	if len(potentialCredentials) == 0 {
		t.Errorf("Expected potential username/password, not empty")
	}else{
		if potentialCredentials[0].Username == "" {
			t.Errorf("Expected potential username to be bug, not " + wi.Form.PotentialCredentials[0].Username)
		}
		if potentialCredentials[0].Password == "" {
			t.Errorf("Expected potential password to be admin, not " + wi.Form.PotentialCredentials[0].Password)
		}
	}
}
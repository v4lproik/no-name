package module

import (
	"testing"
	"github.com/v4lproik/no-name/data"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func TestNewBruteforceModuleWithUsernameArgMissingFormValues(t *testing.T) {
	t.Log("Call bruteforce without username argument form values should not start bruteforcing")

	//given
	ip := "127.0.0.2"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	htmlTagsNames := data.NewHtmlSearchValues(CWD[:strings.LastIndex(CWD, "/")] + "/" + HTML_TAGS_NAMES_TEST)
	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	NewFindFormModule(ip, htmlTagsNames).Request(false, wi)

	// when
	wi.Form.UsernameArg = ""
	NewBruteforceModule(credentials, true, htmlTagsNames.LoginPatterns).Request(false, wi)

	 //then
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected no call to webclient methods are condition is false")
	}
}

func TestNewBruteforceModuleWithPasswordMissingFormValues(t *testing.T) {
	t.Log("Call bruteforce without password argument form values should not start bruteforcing")

	//given
	ip := "127.0.0.2"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	htmlTagsNames := data.NewHtmlSearchValues(CWD[:strings.LastIndex(CWD, "/")] + "/" + HTML_TAGS_NAMES_TEST)
	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	NewFindFormModule(ip, htmlTagsNames).Request(false, wi)

	// when
	wi.Form.PasswordArg = ""
	NewBruteforceModule(credentials, true, htmlTagsNames.LoginPatterns).Request(false, wi)

	//then
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected no call to webclient methods are condition is false")
	}
}

func TestNewBruteforceModuleWithUrlToSubmitMissingFormValues(t *testing.T) {
	t.Log("Call bruteforce without an url to submit the form values should not start bruteforcing")

	//given
	ip := "127.0.0.2"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	htmlTagsNames := data.NewHtmlSearchValues(CWD[:strings.LastIndex(CWD, "/")] + "/" + HTML_TAGS_NAMES_TEST)
	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	NewFindFormModule(ip, htmlTagsNames).Request(false, wi)

	// when
	wi.Form.UrlToSubmit = ""
	NewBruteforceModule(credentials, true, htmlTagsNames.LoginPatterns).Request(false, wi)

	//then
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected no call to webclient methods are condition is false")
	}
}

func TestNewBruteforceModuleWithOtherArgWithValueEqualNilFormValues(t *testing.T) {
	t.Log("Call bruteforce with a data structure equals to null for the other arguments of the form values should not start bruteforcing")

	//given
	ip := "127.0.0.2"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	htmlTagsNames := data.NewHtmlSearchValues(CWD[:strings.LastIndex(CWD, "/")] + "/" + HTML_TAGS_NAMES_TEST)
	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	NewFindFormModule(ip, htmlTagsNames).Request(false, wi)

	// when
	wi.Form.OtherArgWithValue = nil
	NewBruteforceModule(credentials, true, htmlTagsNames.LoginPatterns).Request(false, wi)

	//then
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected no call to webclient methods are condition is false")
	}
}

func TestNewBruteforceModuleWithDocWithValueEqualNilFormValues(t *testing.T) {
	t.Log("Call bruteforce with a the document equals to null should not start bruteforcing")

	//given
	ip := "127.0.0.2"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	htmlTagsNames := data.NewHtmlSearchValues(CWD[:strings.LastIndex(CWD, "/")] + "/" + HTML_TAGS_NAMES_TEST)
	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	NewFindFormModule(ip, htmlTagsNames).Request(false, wi)

	// when
	wi.Doc= nil
	NewBruteforceModule(credentials, true, htmlTagsNames.LoginPatterns).Request(false, wi)

	//then
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected no call to webclient methods are condition is false")
	}
}

func TestNewBruteforceModuleWithNoFormFound(t *testing.T) {
	t.Log("Call bruteforce with no form found should not start bruteforcing")

	//given
	ip := "127.0.0.2"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	htmlTagsNames := data.NewHtmlSearchValues(CWD[:strings.LastIndex(CWD, "/")] + "/" + HTML_TAGS_NAMES_TEST)
	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	NewFindFormModule(ip, htmlTagsNames).Request(false, wi)

	// when
	wi.Form = nil
	NewBruteforceModule(credentials, true, htmlTagsNames.LoginPatterns).Request(false, wi)

	//then
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected no call to webclient methods are condition is false")
	}
}

func TestNewBruteforceModuleWithNoWebClient(t *testing.T) {
	t.Log("Call bruteforce with no web client should not start bruteforcing")

	//given
	ip := "127.0.0.2"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	htmlTagsNames := data.NewHtmlSearchValues(CWD[:strings.LastIndex(CWD, "/")] + "/" + HTML_TAGS_NAMES_TEST)
	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	NewFindFormModule(ip, htmlTagsNames).Request(false, wi)

	// when
	wi.ClientWeb = nil
	NewBruteforceModule(credentials, true, htmlTagsNames.LoginPatterns).Request(false, wi)

	//then
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected no call to webclient methods are condition is false")
	}
}

func TestNewBruteforceModuleWithNoMissingValue(t *testing.T) {
	t.Log("Call bruteforce with no form found should not start bruteforcing")

	//given
	ip := "127.0.0.3"
	webClient := NewFakeWebClient(ip)
	wi := data.NewWebInterface(webClient)

	htmlTagsNames := data.NewHtmlSearchValues(CWD[:strings.LastIndex(CWD, "/")] + "/" + HTML_TAGS_NAMES_TEST)
	credentials := data.NewCredentials(CWD[:strings.LastIndex(CWD, "/")] + "/" + DEFAULT_PASSWORD_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + LOGIN_TEST,
		CWD[:strings.LastIndex(CWD, "/")] + "/" + PASSWORD_TEST)

	res, _ := webClient.Scrap()
	wi.Doc, _ = goquery.NewDocumentFromResponse(res)

	NewFindFormModule(ip, htmlTagsNames).Request(false, wi)

	// when
	NewBruteforceModule(credentials, true, htmlTagsNames.LoginPatterns).Request(false, wi)

	//then
	if webClient.CountScrapWithParameter == 1 {
		t.Errorf("Expected no call to webclient methods are condition is false")
	}
	potentialCredentials := cleanSlice(wi.Form.PotentialCredentials)
	if potentialCredentials[0].Username != "bug" {
		t.Errorf("Expected potential username to be bug, not " + wi.Form.PotentialCredentials[0].Username)
	}
	if potentialCredentials[0].Password != "admin" {
		t.Errorf("Expected potential password to be admin, not " + wi.Form.PotentialCredentials[0].Password)
	}
}
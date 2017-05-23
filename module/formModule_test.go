package module

import (
	"testing"
	"github.com/v4lproik/no-name/data"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"os"
	"github.com/v4lproik/no-name/client"
)

const HTML_TAGS_NAMES_TEST = "conf/html-detection-tags_test.txt"
var CWD, _ = os.Getwd()

func TestNewFormModule(t *testing.T) {
	t.Log("Call Form Module with a valid HTML form to analyse should get analysed by the module")

	//given
	ip := "127.0.1.7/formpath?arguments"
	webClient := client.NewSimpleWebClient(ip)
	wi := data.NewWebInterface(webClient)
	wi.Doc, _ = goquery.NewDocumentFromReader(strings.NewReader(`"<html><form action="url_to_submit" method="POST"><input type="text" name="username" /><input type="password" name="password"><input type="text" name="otherinput" value="random"/><input type="hidden" name="user_token" value="csrftoken" /></form></html>"`))
	htmlTagsNames := data.NewHtmlTagsNames(CWD[:strings.LastIndex(CWD, "/")] + "/" + HTML_TAGS_NAMES_TEST)

	// when
	NewFindFormModule(ip, htmlTagsNames).Request(false, wi)

	 //then
	if wi.Form.UrlForm != "/formpath?arguments" {
		t.Errorf("Expected form's url to contain path and arguments, not : <" + wi.Form.UrlForm + ">")
	}
	if wi.Form.UsernameArg != "username" {
		t.Errorf("Expected username input to be found, not : <" + wi.Form.UsernameArg + ">")
	}
	if wi.Form.PasswordArg != "password" {
		t.Errorf("Expected password input to be found, not : <" + wi.Form.PasswordArg + ">")
	}
	if wi.Form.UrlToSubmit != "url_to_submit" {
		t.Errorf("Expected the url to submit the form to be found, not : <" + wi.Form.UrlToSubmit + ">")
	}
	if wi.Form.MethodSubmitArg != "POST" {
		t.Errorf("Expected the method to submit the form to be found, not : <" + wi.Form.MethodSubmitArg + ">")
	}
	if wi.Form.OtherArgWithValue["otherinput"] != "random" && wi.Form.OtherArgWithValue["user_token"] != "csrftoken" {
		t.Errorf("Expected the other values to submit the form input to be found, including <csrf token> and <other>")
	}
	if wi.Form.CsrfArg != "user_token" {
		t.Errorf("Expected the csrf token form input to be found, not : <" + wi.Form.CsrfArg + ">")
	}
}
package util

import (
	"github.com/pmezard/go-difflib/difflib"
	"github.com/PuerkitoBio/goquery"
	"github.com/juju/loggo"
	"strings"
	"net/http"
)

var loggerUtilHTML = loggo.GetLogger("utilHTML")

func GetDocument(res *http.Response) (*goquery.Document, error){
	return goquery.NewDocumentFromResponse(res)
}

func GetDiffBetweenTwoPages(page1 string, page2 string) (float64){
	sequenceMatcher := difflib.SequenceMatcher{}
	sequenceMatcher.SetSeq1(difflib.SplitLines(page1))
	sequenceMatcher.SetSeq2(difflib.SplitLines(page2))

	return sequenceMatcher.Ratio()
}

func IsFaviconInDocument(doc *goquery.Document) (string, error){
	if doc == nil {
		return "", nil
	}

	urlFavicon := ""
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		rel, exists := s.Attr("rel")
		if exists {
			if strings.Contains(rel, "icon") || strings.Contains(rel, "ICON"){
				href, exists := s.Attr("href")
				if exists{
					loggerUtilHTML.Debugf("Favicon found at address " + href)
					urlFavicon = href
					return
				}
			}
		}
	})


	return urlFavicon, nil
}

func GetCsrfInDocument(doc *goquery.Document, csrfName string) (string, error){
	if doc == nil {
		return "", nil
	}

	csrfValue := ""
	doc.Find("form input").Each(func(i int, s *goquery.Selection) {
		value, existsValue := s.Attr("value")
		if existsValue {
			name, existsName := s.Attr("name")
			if existsName {
				if name == csrfName {
					loggerUtilHTML.Debugf("Csrf found with value " + value)
					csrfValue = value
					return
				}
			}
		}
	})

	return csrfValue, nil
}

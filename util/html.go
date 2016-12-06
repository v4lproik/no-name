package util

import (
	"github.com/pmezard/go-difflib/difflib"
	"github.com/PuerkitoBio/goquery"
	"github.com/juju/loggo"
	"strings"
)

var loggerUtilHTML = loggo.GetLogger("utilHTML")

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

	// find url to submit the form
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		rel, exists := s.Attr("rel")
		if exists {
			if strings.Contains(rel, "icon") || strings.Contains(rel, "ICON"){
				href, exists := s.Attr("href")
				if exists{
					loggerUtilHTML.Debugf("Favicon found at address " + href)
					return href, nil
				}
			}
		}
	})


	return "", nil
}

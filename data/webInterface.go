package data

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/v4lproik/no-name/client"
	"github.com/v4lproik/no-name-domain"
)

type WebInterface struct {
	ClientWeb *client.Web

	Doc *goquery.Document
	Form *domain.Form

	ReportPath string
}

func NewWebInterface(webClient *client.Web) *WebInterface{
	return &WebInterface{webClient, nil, domain.NewForm(), ""}
}

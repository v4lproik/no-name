package data

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/yinkozi/no-name/client"
	"github.com/yinkozi/no-name-domain"
)

type WebInterface struct {
	ClientWeb *client.Web

	Doc *goquery.Document
	Form *domain.Form
}

func NewWebInterface(webClient *client.Web) *WebInterface{
	return &WebInterface{webClient, nil, domain.NewForm()}
}

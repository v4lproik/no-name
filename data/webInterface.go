package data

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/v4lproik/wias/client"
)

type WebInterface struct {
	ClientWeb *client.Web

	Doc *goquery.Document
	Form *Form
}

func NewWebInterface(webClient *client.Web) *WebInterface{
	return &WebInterface{webClient, nil, NewForm()}
}

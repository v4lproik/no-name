package data

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/yinkozi/no-name/client"
)

type WebInterface struct {
	ClientWeb *client.Web

	Doc *goquery.Document
	Form *Form
}

func NewWebInterface(webClient *client.Web) *WebInterface{
	return &WebInterface{webClient, nil, NewForm()}
}

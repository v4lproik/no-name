package module

import (
	"github.com/yinkozi/no-name/data"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type formModule struct {
	next Module
	name string
}

func NewFindFormModule(name string) *formModule {
	return &formModule{nil, name}
}

func (m *formModule) Request(flag bool, wi *data.WebInterface) {

	if wi.Doc != nil {

		form := wi.Form

		// find url to submit the form
		wi.Doc.Find("form").Each(func(i int, s *goquery.Selection) {
			action, exists := s.Attr("action")
			if exists {
				form.UrlToSubmit = action
				logger.Debugf("SubmitUrl has been found with action <" + action + ">")
			}

			method, exists := s.Attr("method")
			if exists {
				form.MethodSubmitArg = strings.ToUpper(method)
				logger.Debugf("MethodSubmit has been found with method <" + method + ">")
			}else{
				form.MethodSubmitArg = "GET"
			}
		})

		// find mandatory inputs to submit the form
		wi.Doc.Find("form input").Each(func(i int, s *goquery.Selection) {
			// find by name
			names, exists := s.Attr("name")
			if exists {
				//try to find if the input is a username or a password field
				switch {
				case strings.Contains(names, "username"):
					logger.Debugf("Username input has been found with name <" + names + ">")
					form.UsernameArg = names
				case strings.Contains(names, "password"):
					logger.Debugf("Password has been found with name <" + names + ">")
					form.PasswordArg = names
				}
			}

			// find by type
			types, exists := s.Attr("type")
			if exists {
				//try to find if the input is a username or a password field
				switch {
				case strings.Contains(types, "submit"):
					logger.Debugf("Submit input has been found with type <" + types + ">")
					form.SubmitArg = types

				case strings.Contains(types, "password"):
					logger.Debugf("Password input has been found with type <" + types + ">")
					form.PasswordArg = types
				}
			}

			// find by value
			value, existsValue := s.Attr("value")
			if existsValue {
				name, existsName := s.Attr("name")
				if existsName {
					form.OtherArgWithValue[name] = value
					logger.Debugf("Couple name=value has been found <" + name + "=" + value + ">")
				}
			}

			wi.Form = form
		})

	}

	if flag && m.next != nil{
		m.next.Request(flag, wi)
	}
}

func (m *formModule) SetNextModule(next Module){
	m.next = next
}
package module

import (
	"github.com/v4lproik/no-name/data"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/v4lproik/no-name/util"
)

type formModule struct {
	name string
	htmlTagsNames *data.HtmlSearchValues

	next Module
}

func NewFindFormModule(name string, htmlTagsNames *data.HtmlSearchValues) *formModule {
	return &formModule{name, htmlTagsNames, nil}
}

func (m *formModule) Request(flag bool, wi *data.WebInterface) {

	if wi.Doc != nil {
		form := wi.Form
		formHtml := wi.Doc.Find("form")

		if len(formHtml.Nodes) < 1 {
			logger.Infof("No form has been found for url " + wi.ClientWeb.GetUrl().String())
		} else {
			// set the url of the form
			wi.Form.UrlForm = wi.ClientWeb.GetUrl().RequestURI()

			// find arguments to submit the form
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
				} else {
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
					case util.Contains(m.htmlTagsNames.UsernameNames, names):
						logger.Debugf("Username input has been found with name <" + names + ">")
						form.UsernameArg = names
					case util.Contains(m.htmlTagsNames.PasswordNames, names):
						logger.Debugf("Password has been found with name <" + names + ">")
						form.PasswordArg = names
					case util.ContainsRegex(m.htmlTagsNames.CsrfNames, names):
						logger.Debugf("Csrf has been found with name <" + names + ">")
						form.CsrfArg = names
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

						names, exists := s.Attr("name")
						if exists {
							form.PasswordArg = names
						}

					}
				}

				// find by value and if this value is not already contained as a username, password etc
				value, existsValue := s.Attr("value")
				if existsValue {
					name, existsName := s.Attr("name")
					if existsName && name != form.UsernameArg && name != form.PasswordArg && name != form.CsrfArg {
						form.OtherArgWithValue[name] = value
						logger.Debugf("Couple name=value has been found <" + name + "=" + value + ">")
					}
				}


				//default values for non found input
				if form.UrlToSubmit == "" {
					form.UrlToSubmit = form.UrlForm
				}
			})

		}
	}

	if flag && m.next != nil{
		m.next.Request(flag, wi)
	}
}

func (m *formModule) SetNextModule(next Module){
	m.next = next
}
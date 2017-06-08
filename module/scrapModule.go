package module

import (
	"github.com/juju/loggo"
	"github.com/v4lproik/no-name/data"
	"github.com/v4lproik/no-name/util"
	"github.com/tebeka/selenium"
	"encoding/base64"
	"net/url"
)

var logger = loggo.GetLogger("scrapModule")

type scrapModule struct {
	seleniumServerUrl string

	next Module
}

func NewScrapModule(seleniumServerUrl string) *scrapModule{
	return &scrapModule{seleniumServerUrl, nil}
}

func (m *scrapModule) Request(flag bool, wi *data.WebInterface) {
	urlToScrap := wi.ClientWeb.CraftUrlGet(wi.ClientWeb.GetUrl().RequestURI(), url.Values{})

	res, err := wi.ClientWeb.Scrap()
	if err != nil {
		logger.Errorf("Can't reach url " + urlToScrap, err.Error())
		return
	}

	if len(m.seleniumServerUrl) != 0 {
		caps := selenium.Capabilities{"browserName": "firefox"}
		wd, err := selenium.NewRemote(caps, m.seleniumServerUrl)
		if err != nil {
			panic(err)
		}
		defer wd.Quit()

		logger.Debugf("Selenium is trying to reach " + urlToScrap)
		wd.Get(urlToScrap)
		data, err := wd.Screenshot()
		if data != nil {
			wi.Form.ScreenShot = base64.StdEncoding.EncodeToString(data)
		}
		if err != nil {
			logger.Errorf("Can't take a Screenshot for url " + wi.ClientWeb.GetUrl().RequestURI(), err.Error())
		}
	}

	doc, err := util.GetDocument(res)
	if err != nil {
		logger.Errorf("Can't transform response to document", err.Error())
		return
	}

	wi.Doc = doc
	wi.Form.Domain = wi.ClientWeb.GetUrl().Host

	if flag && m.next != nil{
		m.next.Request(flag, wi)
	}
}

func (m *scrapModule) SetNextModule(next Module){
	m.next = next
}
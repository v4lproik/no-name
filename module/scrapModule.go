package module

import (
	"github.com/juju/loggo"
	"github.com/v4lproik/no-name/data"
	"github.com/v4lproik/no-name/util"
)

var logger = loggo.GetLogger("scrapModule")

type scrapModule struct {
	next Module
}

func NewScrapModule() *scrapModule{
	return &scrapModule{}
}

func (m *scrapModule) Request(flag bool, wi *data.WebInterface) {

	res, err := wi.ClientWeb.Scrap()
	if err != nil {
		logger.Errorf("Can't reach url " + wi.ClientWeb.GetUrl().RequestURI(), err.Error())
		return
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
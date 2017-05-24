package data

import (
	"encoding/json"
	"io/ioutil"
)

type HtmlSearchValues struct {
	UsernameNames []string
	PasswordNames []string
	CsrfNames []string
	LoginPatterns []string
}

func NewHtmlSearchValues(pathDefault string) (*HtmlSearchValues){
	content, err := ioutil.ReadFile(pathDefault)
	if err != nil {
		logger.Errorf(err.Error())
		return nil
	}

	htmlTagsNames := HtmlSearchValues{nil, nil, nil, nil}
	json.Unmarshal(content, &htmlTagsNames)

	return &htmlTagsNames
}
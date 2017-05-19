package data

import (
	"encoding/json"
	"io/ioutil"
)

type HtmlTagsNames struct {
	UsernameNames []string
	PasswordNames []string
}

func NewHtmlTagsNames(pathDefault string) (*HtmlTagsNames){
	content, err := ioutil.ReadFile(pathDefault)
	if err != nil {
		logger.Errorf(err.Error())
		return nil
	}

	htmlTagsNames := HtmlTagsNames{nil, nil}
	json.Unmarshal(content, &htmlTagsNames)

	return &htmlTagsNames
}
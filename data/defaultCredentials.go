package data

import (
	"encoding/json"
	"github.com/juju/loggo"
	"io/ioutil"
	"github.com/v4lproik/wias/util"
)

var logger = loggo.GetLogger("credentials")


type Webinterface struct {
	Favicon string
	Hash string
	Keywords []string
	Credentials []string
}

type Credentials struct {
	Webinterfaces []Webinterface
	Passwords []string
	Logins []string
}

func NewCredentials(pathDefault string, pathPassword string, pathLogin string) (*Credentials){
	//default
	content, err := ioutil.ReadFile(pathDefault)
	if err != nil {
		logger.Errorf(err.Error())
		return nil
	}

	credentials := Credentials{nil, nil, nil}
	json.Unmarshal(content, &credentials.Webinterfaces)

	//passwords
	passwords, err := util.ReadLines(pathPassword)
	if err != nil {
		logger.Errorf(err.Error())
		return nil
	}
	credentials.Passwords = passwords

	//logins
	logins, err := util.ReadLines(pathLogin)
	if err != nil {
		logger.Errorf(err.Error())
		return nil
	}
	credentials.Logins = logins

	return &credentials
}
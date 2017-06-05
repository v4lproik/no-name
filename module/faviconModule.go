package module

import (
	"github.com/v4lproik/no-name/data"
	"net/url"
	"github.com/v4lproik/no-name/util"
	"crypto/md5"
	"encoding/hex"
	"hash"
)

type faviconModule struct {
	defaultWebInterfaces []data.DefaultWebInterface

	hasher hash.Hash

	next Module
}

func NewFaviconModule(defaultWebInterfaces []data.DefaultWebInterface) *faviconModule{
	return &faviconModule{defaultWebInterfaces, md5.New(), nil}
}

func (m *faviconModule) Request(flag bool, wi *data.WebInterface) {

	condition :=  wi.ClientWeb != nil && wi.Doc != nil

	if condition {
		logger.Infof("Start looking for a favicon")

		// find favicon
		faviconPath, err := util.IsFaviconInDocument(wi.Doc)
		if err != nil {
			logger.Errorf("Error during the process of finding a favicon path in the document", err.Error())
		}

		if faviconPath != "" {
			wi.Form.FaviconPath = faviconPath
			logger.Infof("Favicon path found " + faviconPath)


			//search by favicon md5 hash
			res, err := wi.ClientWeb.ScrapWithParameter(wi.Form.FaviconPath, "GET", make(url.Values))
			if err != nil {
				logger.Errorf("Favicon url can't be reached ", err.Error())
			}else{
				fav, err := util.GetDocument(res)
				if err != nil {
					logger.Errorf("Favicon url response can't be transformed into document", err.Error())
				}

				m.hasher.Write([]byte(fav.Text()))
				wi.Form.FaviconMD5Hash = hex.EncodeToString(m.hasher.Sum(nil))
				logger.Infof("MD5 Favicon is : " + wi.Form.FaviconMD5Hash)

				//search if favicon is known in default web interfaces
				for _, value := range m.defaultWebInterfaces {
					if(wi.Form.FaviconMD5Hash == value.Favicon){
						logger.Infof("Favicon is known in the database with the web interface name of : " + value.Title)
						break;
					}
				}
			}

		}

		if flag && m.next != nil{
			m.next.Request(flag, wi)
		}
	}
}

func (m *faviconModule) SetNextModule(next Module){
	m.next = next
}
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
	credentials *data.Credentials

	hasher hash.Hash

	next Module
}

func NewFaviconModule(credentials *data.Credentials) *faviconModule{
	return &faviconModule{credentials, md5.New(), nil}
}

func (m *faviconModule) Request(flag bool, wi *data.WebInterface) {

	condition :=  wi.ClientWeb != nil && wi.Doc != nil

	if condition {
		logger.Infof("Start looking for a favicon")

		// find favicon
		faviconPath, err := util.IsFaviconInDocument(wi.Doc)
		if err != nil {
			logger.Errorf("Error trying to extract favicon from document", err.Error())
		}

		if faviconPath != "" {
			wi.Form.FaviconPath = faviconPath
			logger.Infof("Favicon path found " + faviconPath)
		}

		//search by favicon md5 hash
		urlToFavicon := wi.ClientWeb.CraftUrl(wi.Form.FaviconPath)
		res, err := wi.ClientWeb.ScrapWithParameter(urlToFavicon, "GET", make(url.Values))
		if err != nil {
			logger.Errorf("Favicon can't be reached ", err.Error())
		}else{
			fav, err := wi.ClientWeb.GetDocument(res)
			if err != nil {
				logger.Errorf("Favicon data can't be transformed into document", err.Error())
			}

			m.hasher.Write([]byte(fav.Text()))
			logger.Infof("MD5 Favicon is : " + hex.EncodeToString(m.hasher.Sum(nil)))
		}

		if flag && m.next != nil{
			m.next.Request(flag, wi)
		}
	}
}

func (m *faviconModule) SetNextModule(next Module){
	m.next = next
}
package module

import (
	"github.com/v4lproik/wias/data"
	"net/url"
	"github.com/v4lproik/wias/util"
	"strconv"
	"crypto/md5"
	"hash"
)

type bruteforceModule struct {
	credentials *data.Credentials

	hasher hash.Hash

	stopFirstFound bool
	next Module
}

func NewBruteforceModule(credentials *data.Credentials, stopFirstFound bool) *bruteforceModule{
	return &bruteforceModule{credentials, md5.New(), stopFirstFound, nil}
}

func (m *bruteforceModule) Request(flag bool, wi *data.WebInterface) {

	condition :=  wi.ClientWeb != nil &&
		wi.Doc != nil &&
		wi.Form != nil &&
		wi.Form.OtherArgWithValue != nil &&
		wi.Form.UrlToSubmit != "" &&
		wi.Form.PasswordArg != "" &&
		wi.Form.UsernameArg != ""

	if condition {
		logger.Debugf("Start bruteforcing")

		values := make(url.Values)
		values.Set(wi.Form.UsernameArg, "test")
		values.Set(wi.Form.PasswordArg, "test")

		for k, v := range wi.Form.OtherArgWithValue {
			values.Set(k, v)
		}

		//get page with errors for comparison purpose
		resWithBadCredentials := ""
		res, err := wi.ClientWeb.ScrapWithParameter(wi.Form.UrlToSubmit, wi.Form.MethodSubmitArg, values)
		if err != nil {
			logger.Errorf("Url bruteforce can't be reached ", err.Error())
		}else{
			doc, err := wi.ClientWeb.GetDocument(res)
			if err != nil {
				logger.Errorf("Data bruteforce can't be transformed into document", err.Error())
			}
			resWithBadCredentials = doc.Text()
		}


		//start bruteforcing
		if resWithBadCredentials != "" {
			found  := false
			for _, usernameTry := range m.credentials.Logins {
				for _, passwordTry := range m.credentials.Passwords {
					// bruteforce ON
					values.Set(wi.Form.UsernameArg, usernameTry)
					values.Set(wi.Form.PasswordArg, passwordTry)

					resWithPotentialGoodCredentials := ""
					res, err = wi.ClientWeb.ScrapWithParameter(wi.Form.UrlToSubmit, wi.Form.MethodSubmitArg, values)
					if err != nil {
						logger.Errorf("Url bruteforce can't be reached ", err.Error())
					}else{
						doc, err := wi.ClientWeb.GetDocument(res)
						if err != nil {
							logger.Errorf("Data bruteforce can't be transformed into document", err.Error())
						}
						resWithPotentialGoodCredentials = doc.Text()
					}

					ratioDiff := util.GetDiffBetweenTwoPages(resWithBadCredentials, resWithPotentialGoodCredentials)
					logger.Debugf("Ratio <" + usernameTry + "/" + passwordTry + ">" + strconv.FormatFloat(ratioDiff, 'f', 6, 64))
					if ratioDiff < 0.92 {
						logger.Infof("Potential credentials: <" + usernameTry + "/" + passwordTry + ">")
						found = true;
						if m.stopFirstFound {
							break;
						}
					}
				}

				if found && m.stopFirstFound {
					break
				}
			}
		}

	}

	if flag && m.next != nil{
		m.next.Request(flag, wi)
	}
}

func (m *bruteforceModule) SetNextModule(next Module){
	m.next = next
}
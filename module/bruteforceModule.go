package module

import (
	"github.com/v4lproik/no-name/data"
	"net/url"
	"github.com/v4lproik/no-name/util"
	"strconv"
	"github.com/v4lproik/no-name-domain"
)

type bruteforceModule struct {
	credentials *data.Credentials
	htmlSearchValues []string

	stopFirstFound bool
	next Module
}

func NewBruteforceModule(credentials *data.Credentials, stopFirstFound bool, htmlSearchValues []string) *bruteforceModule{
	return &bruteforceModule{credentials, htmlSearchValues, stopFirstFound, nil}
}

func (m *bruteforceModule) Request(flag bool, wi *data.WebInterface) {

	if wi.ClientWeb == nil {
		logger.Criticalf("No web client has been initialised. Breuteforce can't start")
		return
	}

	//TODO: need to create a retry system in case of timeout
	if wi.ClientWeb.GetDomainHttpCode() == 401 {
		logger.Debugf("Start bruteforcing")

		found := false
		for _, usernameTry := range m.credentials.Logins {
			for _, passwordTry := range m.credentials.Passwords {
				res, err := wi.ClientWeb.BasicAuth(wi.ClientWeb.GetDomain().RequestURI(), "GET", usernameTry, passwordTry)
				if err != nil {
					logger.Errorf("Url bruteforce can't be reached ", err.Error())
				}else{
					logger.Debugf("Ratio <" + usernameTry + "/" + passwordTry + "> with http response " + strconv.Itoa(res.StatusCode))
					if res != nil && res.StatusCode != 401 {
						logger.Infof("Potential credentials: <" + usernameTry + "/" + passwordTry + ">")
						wi.Form.PotentialCredentials = append(
							wi.Form.PotentialCredentials,
							domain.PotentialCredentials{usernameTry,
										    passwordTry,
										    domain.SourceBruteforce})
						found = true
					}
				}

				if found && m.stopFirstFound {
					break;
				}

			}
		}
	}else{
		condition := wi.Doc != nil &&
			wi.Form != nil &&
			wi.Form.OtherArgWithValue != nil &&
			wi.Form.UrlToSubmit != "" &&
			wi.Form.PasswordArg != "" &&
			wi.Form.UsernameArg != ""

		if condition {
			logger.Debugf("Start bruteforcing")

			values := m.getHTTPArguments(wi, "TEST", "TEST")

			//get page with errors for comparison purpose
			resWithBadCredentials, err := m.getErrorCredentialsPage(wi, values)
			if err != nil {
				logger.Errorf(err.Error())
				return
			}

			//start bruteforcing
			if resWithBadCredentials != "" {
				found := false
				for _, usernameTry := range m.credentials.Logins {
					for _, passwordTry := range m.credentials.Passwords {
						// bruteforce ON
						values := m.getHTTPArguments(wi, usernameTry, passwordTry)

						resWithPotentialGoodCredentials := ""
						res, err := wi.ClientWeb.ScrapWithParameter(wi.Form.UrlToSubmit, wi.Form.MethodSubmitArg, values)
						if err != nil {
							logger.Errorf("Url bruteforce can't be reached ", err.Error())
						}else{
							doc, err := util.GetDocument(res)
							if err != nil {
								logger.Errorf("Data bruteforce can't be transformed into document", err.Error())
							}
							resWithPotentialGoodCredentials = doc.Text()
						}

						ratioDiff := util.GetDiffBetweenTwoPages(resWithBadCredentials, resWithPotentialGoodCredentials)
						logger.Debugf("Ratio <" + usernameTry + "/" + passwordTry + ">" + strconv.FormatFloat(ratioDiff, 'f', 6, 64))

						switch {
						case ratioDiff < 0.92:
							logger.Infof("Potential credentials: <" + usernameTry + "/" + passwordTry + ">")
							wi.Form.PotentialCredentials = append(
								wi.Form.PotentialCredentials,
								domain.PotentialCredentials{usernameTry,
											    passwordTry,
											    domain.SourceBruteforce})
							found = true;
							if m.stopFirstFound {
								break;
							}
						case ratioDiff >= 0.92 && ratioDiff <= 99:
							if(util.ContainsRegex(m.htmlSearchValues, resWithPotentialGoodCredentials)) {
								logger.Infof("Potential credentials: <" + usernameTry + "/" + passwordTry + ">")
								wi.Form.PotentialCredentials = append(
									wi.Form.PotentialCredentials,
									domain.PotentialCredentials{usernameTry,
												    passwordTry,
												    domain.SourceBruteforce})
								found = true;
								if m.stopFirstFound {
									break;
								}
							}
						}
					}

					if found && m.stopFirstFound {
						break
					}
				}
			}
		}
	}


	if flag && m.next != nil{
		m.next.Request(flag, wi)
	}
}

func (m *bruteforceModule) getHTTPArguments(wi *data.WebInterface, username string, password string) (url.Values) {
	values := make(url.Values)

	for k, v := range wi.Form.OtherArgWithValue {
		values.Set(k, v)
	}
	values.Set(wi.Form.UsernameArg, username)
	values.Set(wi.Form.PasswordArg, password)

	if wi.Form.CsrfArg != "" {
		res, err := wi.ClientWeb.ScrapWithNoParameter(wi.Form.UrlForm, "GET")

		if err != nil {
			logger.Errorf("Url to get csrf can't be reached out", err.Error())
		}else{
			doc, err := util.GetDocument(res)
			if err != nil {
				logger.Errorf("Document csrf can't be transformed into document", err.Error())
			}

			csrfValue, _ := util.GetCsrfInDocument(doc, wi.Form.CsrfArg)
			if csrfValue != "" {
				values.Set(wi.Form.CsrfArg, csrfValue)
			}
		}
	}

	return values
}

func (m *bruteforceModule) getErrorCredentialsPage(wi *data.WebInterface, values url.Values) (string, error) {
	res, err := wi.ClientWeb.ScrapWithParameter(wi.Form.UrlToSubmit, wi.Form.MethodSubmitArg, values)
	if err != nil {
		logger.Errorf("Url bruteforce can't be reached ", err.Error())
		return "", err
	}

	doc, err := util.GetDocument(res)
	if err != nil {
		logger.Errorf("Data bruteforce can't be transformed into document", err.Error())
		return "", err
	}

	return doc.Text(), nil
}

func (m *bruteforceModule) SetNextModule(next Module){
	m.next = next
}
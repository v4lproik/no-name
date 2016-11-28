package module

import (
	"github.com/v4lproik/wias/data"
	"net/url"
	"github.com/v4lproik/wias/util"
	"strconv"
)

type bruteforceModule struct {
	next Module
}

func NewBruteforceModule() *bruteforceModule{
	return &bruteforceModule{}
}

func (m *bruteforceModule) Request(flag bool, wi *data.WebInterface) {

	condition :=  wi.ClientWeb != nil &&
		wi.Doc != nil &&
		wi.Form != nil &&
		wi.Form.OtherArgWithValue != nil &&
		wi.Form.UrlToSubmit != "" &&
		wi.Form.SubmitArg != "" &&
		wi.Form.PasswordArg != "" &&
		wi.Form.UsernameArg != ""

	if condition {
		logger.Infof("Start bruteforcing")

		//start bruteforcing
		values := make(url.Values)
		values.Set(wi.Form.UsernameArg, "test")
		values.Set(wi.Form.PasswordArg, "test")

		for k, v := range wi.Form.OtherArgWithValue {
			values.Set(k, v)
		}

		//get page with errors for comparison purpose
		res, err := wi.ClientWeb.ScrapWithParameter(wi.Form.UrlToSubmit, wi.Form.MethodSubmitArg, values)
		resWithBadCredentials := ""
		if err != nil {
			logger.Errorf("Error bruteforcing", err)
		}

		resWithBadCredentials = res.Text()


		// bruteforce ON
		usernameTry := "admin"
		PasswordTry := "password"
		values.Set(wi.Form.UsernameArg, usernameTry)
		values.Set(wi.Form.PasswordArg, PasswordTry)

		res, err = wi.ClientWeb.ScrapWithParameter(wi.Form.UrlToSubmit, wi.Form.MethodSubmitArg, values)
		resWithPotentialGoodCredentials := ""
		if err != nil {
			logger.Errorf("Error bruteforcing", err)
		}
		resWithPotentialGoodCredentials = res.Text()

		ratioDiff := util.GetDiffBetweenTwoPages(resWithBadCredentials, resWithPotentialGoodCredentials)
		logger.Debugf("Ratio " + strconv.FormatFloat(ratioDiff, 'f', 6, 64))
		if ratioDiff != 1.0 {
			logger.Infof("Potential credentials: <" + usernameTry + "/" + PasswordTry + ">")
		}
	}

	if flag && m.next != nil{
		m.next.Request(flag, wi)
	}
}

func (m *bruteforceModule) SetNextModule(next Module){
	m.next = next
}
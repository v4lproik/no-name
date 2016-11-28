package data

type Form struct {
	UrlForm string
	UrlToSubmit string
	UsernameArg string
	PasswordArg string
	SubmitArg string
	OtherArgWithValue map[string]string
	MethodSubmitArg string
}

func NewForm() (f *Form) {
	return &Form{"", "", "", "", "", make(map[string]string), ""}
}

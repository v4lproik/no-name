package module

import (
	"github.com/yinkozi/no-name/data"
	"html/template"
	"io"
	"bytes"
	"os"
	"fmt"
	"time"
	"github.com/google/uuid"
	"path/filepath"
	"encoding/json"
)

type reportModule struct {
	io io.Writer
	format string
	templateDir string

	next Module
}

type Format int
const (
	HTML Format = iota
	GREPABLE
)

const REPORT_FOLDER = "./report"

func NewReportModuleWithSource(rootDir string, format Format, io io.Writer) *reportModule{

	formatChoosen := ""
	switch format {
	case HTML:
		formatChoosen = "html"
	case GREPABLE:
		formatChoosen = "txt"
	}

	return &reportModule{io, formatChoosen, rootDir, nil}
}

func NewReportModule(templateDir string, format Format) *reportModule{

	formatChoosen := ""
	switch format {
	case HTML:
		formatChoosen = "html"
	case GREPABLE:
		formatChoosen = "txt"
	}

	var buf bytes.Buffer

	return &reportModule{&buf, formatChoosen, templateDir, nil}
}

func (m *reportModule) Request(flag bool, wi *data.WebInterface) {

	type reportModule struct {
		Domain string
		UrlForm string
		UrlToSubmit string
		FaviconMD5Hash string
		FaviconPath string
		MethodSubmitArg string
		OtherArgWithValue map[string]string
		PasswordArg string
		UsernameArg string
		PotentialPassword string
		PotentialUsername string
		SubmitArg string
	}

	info := reportModule{
		wi.Form.Domain,
		wi.Form.UrlForm,
		wi.Form.UrlToSubmit,
		wi.Form.FaviconMD5Hash,
		wi.Form.FaviconPath,
		wi.Form.MethodSubmitArg,
		wi.Form.OtherArgWithValue,
		wi.Form.PasswordArg,
		wi.Form.UsernameArg,
		wi.Form.PotentialPassword,
		wi.Form.PotentialUsername,
		wi.Form.SubmitArg,
	}


	var templates = template.Must(template.ParseFiles(filepath.Join(m.templateDir, "./template/report" + "." + m.format)))
	err := templates.ExecuteTemplate(m.io, "report" + "." + m.format, info)
	if err != nil {
		logger.Criticalf("Cannot Get View ", err)
	}

	filename := filepath.Join(m.templateDir, REPORT_FOLDER) + "/" + "report-" + time.Now().String() + "_" + uuid.New().String() + "." + m.format
	f, err := os.Create(filename)
	if err != nil {
		logger.Criticalf("Cannot Create File for Report ", err)
	}else{
		logger.Infof("Report has been created at " + filename)
	}

	fmt.Fprintf(f, "%s", m.io)

}

func (m *reportModule) SetNextModule(next Module){
	m.next = next
}
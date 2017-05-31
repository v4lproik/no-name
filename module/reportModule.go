package module

import (
	"github.com/v4lproik/no-name/data"
	"html/template"
	"io"
	"bytes"
	"os"
	"fmt"
	"time"
	"github.com/google/uuid"
	"path/filepath"
	"strings"
)

type reportModule struct {
	io io.Writer
	format string
	templateDir string

	next Module
}

const REPORT_FOLDER = "./report"

func NewReportModuleWithSource(rootDir string, format data.ReportFormat, io io.Writer) *reportModule{

	formatChoosen := ""
	switch format {
	case data.HTML:
		formatChoosen = "html"
	case data.GREPABLE:
		formatChoosen = "txt"
	}

	createReportFolder()

	return &reportModule{io, formatChoosen, rootDir, nil}
}

func NewReportModule(templateDir string, format data.ReportFormat) *reportModule{

	formatChoosen := ""
	switch format {
	case data.HTML:
		formatChoosen = "html"
	case data.GREPABLE:
		formatChoosen = "txt"
	}

	var buf bytes.Buffer

	createReportFolder()

	return &reportModule{&buf, formatChoosen, templateDir, nil}
}

func createReportFolder() {
	if _, err := os.Stat(REPORT_FOLDER); os.IsNotExist(err) {
		err := os.Mkdir(REPORT_FOLDER, os.FileMode(0755))
		if err != nil {
			logger.Criticalf(err.Error())
		}
	}
}

func (m *reportModule) 	Request(flag bool, wi *data.WebInterface) {

	type reportModule struct {
		Domain string
		ScreenShot string
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
		CsrfArg string
	}

	info := reportModule{
		wi.Form.Domain,
		wi.Form.ScreenShot,
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
		wi.Form.CsrfArg,
	}


	var templates = template.Must(template.ParseFiles(filepath.Join(m.templateDir, "./template/report" + "." + m.format)))
	err := templates.ExecuteTemplate(m.io, "report" + "." + m.format, info)
	if err != nil {
		logger.Criticalf("Cannot Get View ", err)
	}

	filename := filepath.Join(m.templateDir, REPORT_FOLDER) + "/" + "report-" + strings.Replace(time.Now().String() + "_" + wi.Form.Domain + "_" + uuid.New().String() + "." + m.format, " ", "_", -1)
	f, err := os.Create(filename)
	if err != nil {
		logger.Criticalf("Cannot Create File for Report ", err)
	}else{
		fmt.Fprintf(f, "%s", m.io)
		logger.Infof("Report has been created at " + filename)

		wi.ReportPath = filename
	}
}

func (m *reportModule) SetNextModule(next Module){
	m.next = next
}
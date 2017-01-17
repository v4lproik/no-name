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

const REPORT_FOLDER = "./public/report"

func NewReportModuleWithSource(rootDir string, format Format, io io.Writer) *reportModule{

	formatChoosen := ""
	switch format {
	case HTML:
		formatChoosen = "html"
	case GREPABLE:
		formatChoosen = "grep"
	}

	return &reportModule{io, formatChoosen, rootDir, nil}
}

func NewReportModule(templateDir string, format Format) *reportModule{

	formatChoosen := ""
	switch format {
	case HTML:
		formatChoosen = "html"
	case GREPABLE:
		formatChoosen = "grep"
	}

	var buf bytes.Buffer

	return &reportModule{&buf, formatChoosen, templateDir, nil}
}

func (m *reportModule) Request(flag bool, wi *data.WebInterface) {

	if m.format == "html" {

		myVars := map[string]string {
			"Title": "Report - " + wi.Form.UrlToSubmit,
			"Content": "Potential username/password find are: " + wi.Form.PotentialUsername + "/" + wi.Form.PotentialPassword,
		}

		var templates = template.Must(template.ParseFiles(filepath.Join(m.templateDir, "./public/template/report.html")))
		err := templates.ExecuteTemplate(m.io, "report.html", myVars)
		if err != nil {
			logger.Criticalf("Cannot Get View ", err)
		}

		filename := filepath.Join(m.templateDir, REPORT_FOLDER) + "/" + "report-" + time.Now().String() + "_" + uuid.New().String() + "." + m.format
		f, err := os.Create(filename)
		if err != nil {
			logger.Criticalf("Cannot Create File for Report ", err)
		}

		fmt.Fprintf(f, "%s", m.io)
	}
}

func (m *reportModule) SetNextModule(next Module){
	m.next = next
}
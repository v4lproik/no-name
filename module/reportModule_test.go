package module

import (
	"testing"
	"github.com/v4lproik/no-name/data"
	"bytes"
	"strings"
	"os"
)

func TestNewHtmlReport(t *testing.T) {
	t.Log("Call ReportModule with existing data and html output should generate report")

	//given
	wi := data.NewWebInterface(nil)
	wi.Form.PotentialPassword = "password"
	wi.Form.PotentialUsername = "username"
	wi.Form.FaviconPath = "/favicon.icon"
	wi.Form.CsrfArg = "csrf"
	wi.Form.UrlForm = "path-to-form.js"
	wi.Form.UrlToSubmit = "path-to-submit-form.js"
	wi.Form.Domain = "domain.com"
	wi.Form.FaviconMD5Hash = "ab416c39d509e72c5a0a7451a45bc65e"
	wi.Form.MethodSubmitArg = "POST"
	wi.Form.SubmitArg = "toSubmit"
	var buf bytes.Buffer
	cwd, _ := os.Getwd()
	rootDir := cwd[:strings.LastIndex(cwd, "/")]

	// when
	NewReportModuleWithSource(rootDir, 0, &buf).Request(false, wi)

	// then
	if !strings.Contains(buf.String(), wi.Form.PotentialPassword) || !strings.Contains(buf.String(), wi.Form.PotentialUsername) {
		t.Errorf("Expected PotentialPassword and PotentialUsername to be in the report")
	}
	//if !strings.Contains(buf.String(), wi.Form.CsrfArg) || !strings.Contains(buf.String(), wi.Form.MethodSubmitArg) || strings.Contains(buf.String(), wi.Form.SubmitArg) {
	//	t.Errorf("Expected other aguments to be in the report")
	//}
	if !strings.Contains(buf.String(), wi.Form.UrlToSubmit) || !strings.Contains(buf.String(), wi.Form.UrlForm) {
		t.Errorf("Expected url links to be in the report")
	}
	if !strings.Contains(buf.String(), wi.Form.FaviconPath) || !strings.Contains(buf.String(), wi.Form.FaviconMD5Hash) {
		t.Errorf("Expected favicon args to be in the report")
	}
	if !strings.Contains(buf.String(), wi.Form.Domain) {
		t.Errorf("Expected domain name to be in the report")
	}
}

func TestNewTxtReport(t *testing.T) {
	t.Log("Call ReportModule with existing data and txt output should generate report")

	//given
	wi := data.NewWebInterface(nil)
	wi.Form.PotentialPassword = "password"
	wi.Form.PotentialUsername = "username"
	var buf bytes.Buffer
	cwd, _ := os.Getwd()
	rootDir := cwd[:strings.LastIndex(cwd, "/")]

	// when
	NewReportModuleWithSource(rootDir, 1, &buf).Request(false, wi)

	// then
	if !strings.Contains(buf.String(), wi.Form.PotentialPassword) || !strings.Contains(buf.String(), wi.Form.PotentialUsername) {
		t.Errorf("Expected PotentialPassword and PotentialUsername to be in the report")
	}
}

func TestNewTxtReportNameWithoutSpaces(t *testing.T) {
	t.Log("Call ReportModule with existing data and txt output should generate report with a name without spaces")

	//given
	wi := data.NewWebInterface(nil)
	wi.Form.PotentialPassword = "password"
	wi.Form.PotentialUsername = "username"
	var buf bytes.Buffer
	cwd, _ := os.Getwd()
	rootDir := cwd[:strings.LastIndex(cwd, "/")]

	// when
	NewReportModuleWithSource(rootDir, 1, &buf).Request(false, wi)

	// then
	if strings.Contains(wi.ReportPath, " ") {
		t.Errorf("The report's name should not contain any spaces")
	}
}
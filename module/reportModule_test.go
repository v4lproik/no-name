package module

import (
	"testing"
	"github.com/yinkozi/no-name/data"
	"bytes"
	"strings"
	"os"
)

func TestNewHtmlReport(t *testing.T) {
	t.Log("Call sReportModule with existing data and html output should generate report")

	//given
	wi := data.NewWebInterface(nil)
	wi.Form.PotentialPassword = "password"
	wi.Form.PotentialUsername = "username"
	var buf bytes.Buffer
	cwd, _ := os.Getwd()
	rootDir := cwd[:strings.LastIndex(cwd, "/")]

	// when
	NewReportModuleWithSource(rootDir, 0, &buf).Request(false, wi)

	// then
	if !strings.Contains(buf.String(), wi.Form.PotentialPassword) || !strings.Contains(buf.String(), wi.Form.PotentialUsername) {
		t.Errorf("Expected PotentialPassword and PotentialUsername to be in the report")
	}
}

//func TestNewTxtReport(t *testing.T) {
//	t.Log("Call ReportModule with existing data and txt output should generate report")
//
//	//given
//	wi := data.NewWebInterface(nil)
//	wi.Form.PotentialPassword = "password"
//	wi.Form.PotentialUsername = "username"
//	var buf bytes.Buffer
//
//	// when
//	NewReportModuleWithSource(1, &buf).Request(false, wi)
//
//	// then
//	if !strings.Contains(buf.String(), wi.Form.PotentialPassword) || !strings.Contains(buf.String(), wi.Form.PotentialUsername) {
//		t.Errorf("Expected PotentialPassword and PotentialUsername to be in the report")
//	}
//}
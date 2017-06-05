package main

import (
	"testing"
	"github.com/v4lproik/no-name/data"
	"github.com/v4lproik/no-name/util"
	"strings"
)

// TODO: This is poorly written ! It is a Temporary solution per port...
// Need to use the same ip for every environment
func TestNewCredentials(t *testing.T) {
	//given
	ips, channels, chains := setUp("ip_test.txt", "", data.GREPABLE,
		"conf/default-password-web-interface_test.txt", "conf/password_test.txt",
		"conf/login_test.txt", "conf/html-detection-tags_test.txt")

	//when
	launchChains(ips, channels, chains)
	reportPaths := waitForResponse(channels)

	//then
	for _, reportPath := range reportPaths {
		content, _ := util.ReadLines(strings.Split(reportPath, " => ")[1])

		if len(content) < 1 {
			t.Errorf("The report has not been written, impossible to analyze " + reportPath)
			continue
		}
		value := content[0]

		//if its port webgoat 8080
		if strings.Contains(value, ":8080/") {
			if !strings.Contains(value, "admintest/admintest") {
				t.Error("The credentials has not been found for the vulnerable box on port 8080")
			}
		}else{
			if strings.Contains(value, ":80/") {
				if !strings.Contains(value, "admin/password") {
					t.Error("The credentials cannot be found for the vulnerable box on port 80")
				}
			} else {
				if strings.Contains(value, ":8081/") {
					if !strings.Contains(value, "//") {
						t.Error("The credentials cannot be found for the vulnerable box on port 8081")
					}
				}else{
					if strings.Contains(value, ":8899/") {
						if !strings.Contains(value, "user1/user1_pass") {
							t.Error("The credentials cannot be found for the vulnerable box on port 8899")
						}
					}else {
						if strings.Contains(value, ":8088/") {
							if !strings.Contains(value, "test/test") {
								t.Error("The credentials cannot be found for the vulnerable box on port 8088")
							}
						}else{
							t.Errorf("A new vulnerable box has been added to the test without being tested " + value)
						}
					}
				}
			}
		}
	}
}
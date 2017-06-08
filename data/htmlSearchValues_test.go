package data

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestNewHtmlSearchValues(t *testing.T) {
	t.Log("Call NewCredentials with existing file")

	//given
	const HTML = "../conf/html-detection-tags_test.txt"

	// when
	htmlTags := NewHtmlSearchValues(HTML)

	// then
	assert.Equal(t, htmlTags.CsrfNames[0], ".*csrf.*")
	assert.Equal(t, htmlTags.CsrfNames[1], ".*_token")
	assert.Equal(t, htmlTags.PasswordNames[0], ".*password.*")
	assert.Equal(t, htmlTags.UsernameNames[0], ".*username.*")
	assert.Equal(t, htmlTags.LoginPatterns[0], "Login Successful")
}
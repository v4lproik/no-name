package data

import (
	"testing"
)

func TestNewCredentials(t *testing.T) {
	t.Log("Call NewCredentials with existing file")

	//given
	path := "../conf/default-password-web-interface_test.txt"

	// when
	credentials := NewCredentials(path)

	// then
	if credentials.Webinterfaces != nil && len(credentials.Webinterfaces) != 2 &&
		credentials.Webinterfaces[0].Hash == "hash" &&
		credentials.Webinterfaces[1].Hash == "hash2" {
		t.Errorf("Expected credentials.Webinterfaces to be length 1")
	}

}
package data

import (
	"testing"
)


func TestNewCredentials(t *testing.T) {
	t.Log("Call NewCredentials with existing file")

	//given
	const LOGIN = "../conf/login_test.txt"
	const PASSWORD = "../conf/password_test.txt"
	const DEFAULT_PASSWORD = "../conf/default-password-web-interface_test.txt"

	// when
	credentials := NewCredentials(DEFAULT_PASSWORD, PASSWORD, LOGIN)

	// then
	if credentials.DefaultWebInterfaces == nil || len(credentials.DefaultWebInterfaces) != 2 ||
		credentials.DefaultWebInterfaces[0].Hash != "hash" || credentials.DefaultWebInterfaces[1].Hash != "94ded9c7e35f95e7b3fe36a0269c1e00" {
		t.Errorf("Expected credentials Webinterfaces to be length 2")
	}
}
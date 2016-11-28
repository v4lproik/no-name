package service

import (
	"testing"
)

func TestArgumentsService_CheckArguments_WithArguments(t *testing.T) {
	t.Log("Call CheckArguments program with appropriate arguments... (Expecting no error)")

	//given
	chars := []string{0:"--filename=ok"}
	a := new(ArgumentsService)

	//when
	a.CheckArguments(chars)
}

func TestArgumentsService_CheckArguments_WithoutArguments(t *testing.T) {
	t.Log("Call CheckArguments program without appropriate arguments... (Expecting error)")

	//given
	chars := []string{0:"test"}
	a := new(ArgumentsService)

	//when
	_ , err := a.CheckArguments(chars)

	//then
	if err == nil {
		t.Errorf("Expected error due to missing arguments, but no error instead", err)
	}
}



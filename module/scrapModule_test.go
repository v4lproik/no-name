package module

import (
	"testing"
	"github.com/v4lproik/no-name/data"
	"github.com/stretchr/testify/assert"
)

//var _ Module = (*fakeNextModule)(nil)

func TestNewScrapModule_withNonReachableUrl_shouldReturnError_and_stopTheChain(t *testing.T) {
	t.Log("Call ScrapModule with non-reachable url should return an error & stop the chain")

	//given
	sm := scrapModule{}
	wi := data.NewWebInterface(NewFakeWebClient("127.0.0.0.0.a"))
	fakeNextModule:= &fakeNextModule{0}
	sm.SetNextModule(fakeNextModule)

	// when
	sm.Request(true, wi)

	// then
	assert.Equal(t, fakeNextModule.count, 0)
}

func TestNewScrapModule_withNonParsableResponse_shouldReturnError_and_stopTheChain(t *testing.T) {
	t.Log("Call ScrapModule with non-parsable response should return an error & stop the chain")

	//given
	sm := scrapModule{}
	wi := data.NewWebInterface(NewFakeWebClient("127.0.0.12"))
	fakeNextModule:= &fakeNextModule{0}
	sm.SetNextModule(fakeNextModule)

	// when
	sm.Request(true, wi)

	// then
	assert.Equal(t, fakeNextModule.count, 0)
}

func TestNewScrapModule_withParsableResponse_shouldContinueTheChain(t *testing.T) {
	t.Log("Call ScrapModule with parsable response should continue the chain")

	//given
	sm := scrapModule{}
	wi := data.NewWebInterface(NewFakeWebClient("127.0.0.1"))
	fakeNextModule:= &fakeNextModule{0}
	sm.SetNextModule(fakeNextModule)

	// when
	sm.Request(true, wi)

	// then
	assert.Equal(t, fakeNextModule.count, 1)
}
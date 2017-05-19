package module

import "github.com/v4lproik/no-name/data"

type Module interface {
	Request(flag bool, wi *data.WebInterface)
	SetNextModule(next Module)
}
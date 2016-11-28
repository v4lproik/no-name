package module

import "github.com/v4lproik/wias/data"

type Module interface {
	Request(flag bool, wi *data.WebInterface)
	SetNextModule(next Module)
}
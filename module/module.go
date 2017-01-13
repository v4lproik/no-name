package module

import "github.com/yinkozi/no-name/data"

type Module interface {
	Request(flag bool, wi *data.WebInterface)
	SetNextModule(next Module)
}
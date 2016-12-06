package data

type favicon struct {
	Hash string
	NameInterface string
}

func NewFavicon(hash string, nameInterface string) (*favicon) {
	return &favicon{hash, nameInterface}
}
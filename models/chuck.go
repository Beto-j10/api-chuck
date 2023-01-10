package models

type Chuck struct {
	Id string `json:"id"`
}

type Chucks struct {
	Chucks []Chuck `json:"chucks"`
}

package entity

type Certificate struct {
	Id         string
	OwnerId    string
	Type       string
	Parameters map[string]string
}

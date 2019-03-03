package model

type Project struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	State       string `json:"state"`
	Revision    int `json:"revision"`
	Visibility  string `json:"visibility"`
}


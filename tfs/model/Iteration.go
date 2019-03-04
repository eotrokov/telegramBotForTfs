package model

import "time"


type Iteration struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	IterationPath string `json:"path"`
	Attributes    struct {
		StartDate  time.Time `json:"startDate"`
		FinishDate time.Time `json:"finishDate"`
		TimeFrame  string    `json:"timeFrame"`
	} `json:"attributes"`
}
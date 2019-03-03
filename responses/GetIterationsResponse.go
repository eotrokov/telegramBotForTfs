package responses

import "../model"

type GetIterationsResponse struct {
	Count int               `json:"count"`
	Items []model.Iteration `json:"value"`
}
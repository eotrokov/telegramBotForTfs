package responses

import "../model"

type GetTeamsResponse struct {
	Count int `json:"count"`
	Items []model.Team `json:"value"`
}
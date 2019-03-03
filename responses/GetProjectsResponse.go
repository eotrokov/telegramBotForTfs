package responses

import "../model"

type GetProjectsResponse struct {
	Count int `json:"count"`
	Items []model.Project `json:"value"`
}

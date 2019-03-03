package responses

import "../model"

type GetCollectionsResponse struct {
	Items []model.Collection `json:"__wrappedArray"`
}
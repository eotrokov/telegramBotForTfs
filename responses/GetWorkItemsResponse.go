package responses

type GetWorkItemsResponse struct {
	QueryType string `json:"queryType"`
	WorkItems []WorkItem `json:"workItems"`
}

type WorkItem struct {
	Id int `json:"id"`
	AssignedTo string `json:"AssignedTo"`
}

type GetWorks struct {
	Count int `json:"count"`
	Items []Item `json:"value"`
}

type Item struct {
	Id     int `json:"id"`
	Fields struct {
		Id            int     `json:"System.Id"`
		AssignedTo    string  `json:"System.AssignedTo"`
		CompletedWork float32 `json:"Microsoft.VSTS.Scheduling.CompletedWork"`
	} `json:"fields"`
}
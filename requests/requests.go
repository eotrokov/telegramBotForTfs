package requests

import (
	. "../model"
	. "../requestHelper"
	. "../responses"
	"fmt"
	"github.com/ahmetb/go-linq"
	. "github.com/ahmetb/go-linq"
	"regexp"
	"strconv"
	"strings"
)


func GetCollections(headers map[string][]string) []Collection {
	var keys= GetCollectionsResponse{}
	uri := "/_api/_common/getCollections?api-version=4.1"
	Get(uri, headers, &keys)
	return keys.Items
}

func GetIterations(collectionUri string,
					projectName string,
					teamId string,
					headers map[string][]string) []Iteration {
	var keys= GetIterationsResponse{}
	uri := fmt.Sprintf("/%s%s/%s/_apis/work/teamsettings/iterations",
		collectionUri, projectName, teamId)
	Get(uri, headers, &keys)
	return keys.Items
}

func GetProjects(collectionUri string,
					headers map[string][]string) []Project {
	uri := fmt.Sprintf("/%s_apis/projects?api-version=4.1", collectionUri)
	var keys= GetProjectsResponse{}
	Get(uri, headers, &keys)
	return keys.Items
}

func GetTeams(collectionUri string,
				projectId string,
				headers map[string][]string) []Team {
	var keys= GetTeamsResponse{}
	uri := fmt.Sprintf("/%s_apis/projects/%s/teams?api-version=4.1",
		collectionUri, projectId)
	Get(uri, headers, &keys)
	return keys.Items
}


func GetWorkItems(iterationPath string,
				  teamProjectName string,
				  collectionUri string,
				  projectName string,
				  teamId string,
				  headers map[string][]string) []Work {
	var workItems= GetWorkItemsResponse{}
	query := fmt.Sprintf("SELECT [System.Id] FROM WorkItems"+
		" WHERE [System.TeamProject] = '%s'"+
		" AND [System.IterationPath] = '%s'"+
		" AND [System.WorkItemType] in ('Task')", teamProjectName, iterationPath)
	uri := fmt.Sprintf("%s%s/%s/_apis/wit/wiql?api-version=4.1",
		collectionUri, projectName, teamId)

	Post(uri, WiqlQuery{Query: query}, headers, &workItems)

	var workItemIds []string
	linq.From(workItems.WorkItems).
		Select(func(c interface{}) interface{} { return strconv.Itoa(c.(WorkItem).Id) }).ToSlice(&workItemIds)
	fields := "System.AssignedTo,Microsoft.VSTS.Scheduling.CompletedWork,System.Id"
	query = fmt.Sprintf("wit/workitems?ids=%s&fields=%s", strings.Join(workItemIds, ","), fields)
	query = fmt.Sprintf("%s%s/_apis/%s&api-version=4.1", collectionUri,
		teamProjectName, query)
	var workItemsInfo = GetWorks{}
	Get(query, headers, &workItemsInfo)

	var groups []Group
	From(workItemsInfo.Items).WhereT(func(w Item) bool { return w.Fields.AssignedTo != "" }).
		GroupByT(ClearAssignedTo, func(p Item) float32 { return p.Fields.CompletedWork },
		).OrderByT(func(g Group) string { return g.Key.(string) }).ToSlice(&groups)

	var works []Work
	for _, assignedTo := range groups {
		var completedWork= float32(0)
		for _, work := range assignedTo.Group {
			completedWork += work.(float32)
		}
		w := Work{AssignedTo: assignedTo.Key.(string), CompletedWork: completedWork}
		works = append(works, w)
	}
	return works
}
//clear
func ClearAssignedTo(p Item) string {
	re := regexp.MustCompile(`([[А-ЯЁ][а-яё]*) ([А-ЯЁ][а-яё]*) (<.*>)`)
	return re.ReplaceAllString(p.Fields.AssignedTo, `$2 $1`)
}
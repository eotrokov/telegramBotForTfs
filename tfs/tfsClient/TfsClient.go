package tfsClient

import (
	"../../config"
	. "../model"
	. "../requests"
	"encoding/base64"
	"fmt"
	. "github.com/ahmetb/go-linq"
	"gopkg.in/ahmetb/go-linq.v3"
	"strings"
	"time"
)

func GetIterationWorks(configuration config.Configuration) string {
	headers := make(map[string][]string)
	headers["authorization"] = []string{"Basic " + base64.StdEncoding.EncodeToString([]byte(configuration.TfsLogin +
		":" + configuration.TfsToken))}
	collections := GetCollections(headers)
	currentCollection := From(collections).
		FirstWithT(func(collection Collection) bool { return collection.Name == "RCUTASUDDC" }).(Collection)

	projects := GetProjects(currentCollection.Uri, headers)
	currentProject := From(projects).First().(Project)

	teams := GetTeams(currentCollection.Uri, currentProject.Id, headers)
	currentTeam := From(teams).First().(Team)

	iterations := GetIterations(currentCollection.Uri, currentProject.Name, currentTeam.Id, headers)
	currentIteration := From(iterations).FirstWithT(func(iteration Iteration) bool {
		return iteration.Attributes.StartDate.Before(time.Now()) &&
			iteration.Attributes.FinishDate.After(time.Now())
	}).(Iteration)

	workItems := GetWorkItems(currentIteration.IterationPath,
		currentProject.Name, currentCollection.Uri, currentProject.Name, currentTeam.Id, headers)

	var result []string
	linq.From(workItems).SelectT(func(c Work) string {
		return fmt.Sprintf("%s: %.2f", c.AssignedTo, c.CompletedWork)
	}).ToSlice(&result)

	return strings.Join(result, "\n")
}
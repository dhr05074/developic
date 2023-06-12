package handler

import (
	"fmt"
	"portfolio/schema"
	"strings"
)

type Parameters struct {
	Job              string
	CareerYears      uint
	TechStacks       []string
	Projects         []schema.Project
	PreferredCompany *string
}

func makePrompt(params Parameters) string {
	original := "I'm working on a service that provides feedback on portfolios submitted by people moving to another development jobs. The feedback will tell the user which parts of their tech stack and projects they don't need to include, which parts they should need, and which parts they should fill in, and will suggest additional projects for them to work on. The input and output formats will be as follows. Give me the result as a JSON format in the following output format.\n\n[input format]\njob: string\ncareer_years: uint (0 for new developers)\ntech_stacks: string array\nprojects: array of projects (a project is an object consisting of a title and a description)\npreferred_company: the look or philosophy of the company the new hire wants to work for (optional)\n\n[output format]\nproject_feedbacks: feedback array (feedback is an object consisting of subject, score(out of 10), and comment)\ntech_stack_feedbacks: array of feedbacks\nproject_recommendations: recommend array (recommend is an object consisting of subject, score(out of 10), and reason)\nAnswer in English.\n\n[input]\njob: {{job}}\ncareer_years: {{career_years}}\ntech_stacks: {{tech_stacks}}\nProjects:\n{{projects}}"

	original = strings.ReplaceAll(original, "{{job}}", params.Job)
	original = strings.ReplaceAll(original, "{{career_years}}", fmt.Sprintf("%d", params.CareerYears))
	original = strings.ReplaceAll(original, "{{tech_stacks}}", makeTechStacks(params.TechStacks))
	original = strings.ReplaceAll(original, "{{projects}}", makeProjects(params.Projects))

	if params.PreferredCompany != nil {
		original = strings.ReplaceAll(original, "{{preferred_company}}", *params.PreferredCompany)
	}

	return original
}

func makeTechStacks(techStacks []string) string {
	return strings.Join(techStacks, ",")
}

func makeProjects(projects []schema.Project) string {
	var (
		projectStrings = make([]string, len(projects))
		projectString  string
	)

	for i, project := range projects {
		projectString = fmt.Sprintf("- {\"%s\",\"%s\"}", project.Title, project.Description)
		projectStrings[i] = projectString
	}

	return strings.Join(projectStrings, "\n")
}

package model

import (
	"bytes"
	"os"
	"text/template"
)

func RenderPrompt(prompt string, data map[string]string) (string, error) {
	tpl, err := template.New("prompt").
		Option("missingkey=zero").
		Parse(prompt)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GetResumeOptimizePrompt() (string, error) {
	file, err := os.ReadFile("/Users/chenjiahan/project/Graduation Project/ice_sparkhire_runtime/prompt/resume_optimize.txt")
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func GetRecommendRecruitmentPrompt() (string, error) {
	file, err := os.ReadFile("/Users/chenjiahan/project/Graduation Project/ice_sparkhire_runtime/prompt/recommend_recruitment.txt")
	if err != nil {
		return "", err
	}
	return string(file), nil
}

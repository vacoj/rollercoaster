package main

import (
	"fmt"
	"strings"

	colorful "github.com/lucasb-eyer/go-colorful"
)

type mermaidNode struct {
	Name  string // friendly name for the object being displayed
	Value string // actual connection string
	Form  string // database, cache, service, Internal API, External API
	Link  string // a link to the item, if it's an application
}

var allContent = map[string][]string{}
var allLinks = []string{}

func buildChartList(friendly string, input map[string]mermaidNode, category string) []string {
	result := []string{}
	tags := []string{}
	links := []string{}

	for _, matcher := range Configuration.Targets {
		for _, x := range matcher.Identifiers {
			if strings.Contains(friendly, x) {
				friendly = matcher.FriendlyName
			}
		}
		for s, mmn := range input {
			mn := mmn
			for _, x := range matcher.Identifiers {
				if strings.Contains(mmn.Value, x) {
					mn.Name = matcher.FriendlyName
					input[s] = mn
				}
			}
		}
	}

	color := colorful.FastHappyColor().Hex()
	color2 := colorful.FastWarmColor().Hex()

	result = append(result, fmt.Sprintf("style %s fill:%s,stroke:%s,stroke-width:2px", friendly, color, color2))
	// result = append(result, fmt.Sprintf("linkStyle %s stroke:%s,stroke-width:2px", friendly, color))
	protoTag := true
	for _, v := range input {
		for _, cm := range Configuration.CategoryMeta {
			if cm.Name == category {
				if stringInSlice(v.Form, cm.Forms) {

					protoTag = false
					entry := lineBuilder(friendly, v.Name, v.Form, v.Value, v.Form, protoTag)
					result = append(result, entry)
					if !stringInSlice(v.Name, tags) {
						tags = append(tags, v.Name)
					}
					if !stringInSlice(entry, allContent[category]) {
						allContent[category] = append(allContent[category], entry)
					}
				}
			}
		}
	}

	// add style here based on what it is
	// 	result = append(result, fmt.Sprintf("style %s fill:#6a7991,stroke:#333,stroke-width:4px", friendly))

	for _, d := range result {
		if !stringInSlice(d, allContent[category]) {
			allContent[category] = append(allContent[category], d)
		}
	}

	for _, link := range links {
		fnd := false
		for _, al := range allLinks {
			if link == al {
				fnd = true
			}
			if !fnd {
				allLinks = append(allLinks, link)
			}
		}
	}

	result = append(result, links...)
	writeContent(friendly, result, tags, category)
	return result
}

func lineBuilder(friendly, name, middle, target, ttype string, protoTag bool) string {

	for _, inp := range Configuration.FlowIngressPoints {
		if friendly == inp {
			friendly = fmt.Sprintf("%s>\\\"fa:fa-arrow-right %s\\\"]", friendly, friendly)
		}
	}

	builder := friendly + " -->"

	if middle != "" && protoTag {
		builder += "|" + middle + "|"
	}
	preset := false
	for _, tc := range Configuration.TargetCategories {
		if ttype == tc.Category && !preset && len(tc.DisplayShape) == 2 {
			preset = true
			builder += fmt.Sprintf("%s%s\\\"%s %s\\\"%s", name, tc.DisplayShape[0], tc.Icon, name, tc.DisplayShape[1])
			// fmt.Println(name, tc.Icon)
			break
		} else if ttype == tc.Category && !preset && len(tc.DisplayShape) != 2 {
			preset = true
			builder += fmt.Sprintf("%s%s\\\"%s %s\\\"%s", name, "(", tc.Icon, name, ")")
			// fmt.Println(name, tc.Icon)
			break
		}
	}
	if !preset {
		builder += fmt.Sprintf("%s[\\\"%s %s\\\"]", name, "fa:fa-question-circle", name)
	}

	return builder
}

func buildMermaidNodes(tokenValueMap map[string]string) map[string]mermaidNode {
	deps := map[string]mermaidNode{}
	for k, v := range tokenValueMap {
		for _, ss := range Configuration.DependencyStrings {
			v = strings.ToLower(v)
			ss = strings.ToLower(ss)
			if strings.Contains(v, ss) {

				for _, mappedItem := range Configuration.Targets {
					if sliceElementInString(mappedItem.Identifiers, v) {
						k = mappedItem.FriendlyName
					}
				}
				parseProtocols := func(fName, fInput string) mermaidNode {
					r := mermaidNode{
						Name: fName,
					}
					for _, ci := range Configuration.TargetCategories {
						if sliceElementInString(ci.Identifiers, fInput) {
							r.Form = ci.Category
							r.Link = ""
						}
					}
					if r.Form == "" {
						r.Form = "???"
						r.Link = ""
					}
					return r
				}
				nk := parseProtocols(k, v)
				if !stringInSliceElement(v, Configuration.ExcludedIdentifiers) {
					deps[k] = nk
				}
			}
		}
	}
	return deps
}

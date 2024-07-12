package main

import (
	"fmt"
	"strings"
)

type EnvironmentFilter func(string) bool

func GetEnvironments(environments []string, filters ...EnvironmentFilter) []string {
	filteredEnvs := []string{}

	for _, env := range environments {
		filtered := false
		for _, f := range filters {
			filtered = f(env)
			if filtered {
				break
			}
		}

		if !filtered {
			filteredEnvs = append(filteredEnvs, strings.ToLower(env))
		}
	}

	return filteredEnvs
}

func ProdEnvFilter(env string) bool {
	return strings.HasPrefix(strings.ToLower(env), "prod")
}

func NonProdEnvFilter(env string) bool {
	return !ProdEnvFilter(env)
}

func main() {
	environments := []string{
		"prod1",
		"prod2",
		"PROD3",
		"staging",
		"dev",
		"local",
	}

	fmt.Println("prod envs:", GetEnvironments(environments, ProdEnvFilter))
	fmt.Println("non-prod envs:", GetEnvironments(environments, NonProdEnvFilter))
}

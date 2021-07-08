package manganatoapi

import "strings"

func getId(s string) string {
	return strings.Split(s, "-")[1]
}

func changeSpaceToUnderscore(s string) string {
	return strings.Join(strings.Fields(s), "_")
}

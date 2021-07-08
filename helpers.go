package manganatoapi

import "strings"

func getId(s string) string {
	tmp := strings.Split(s, "-")
	return tmp[len(tmp)-1]
}

func changeSpaceToUnderscore(s string) string {
	return strings.Join(strings.Fields(s), "_")
}

package manganatoapi

import "strings"

func getID(url, sep string) string {
	tmp := strings.Split(url, sep)
	return tmp[len(tmp)-1]
}

func changeSpaceToUnderscore(s string) string {
	return strings.Join(strings.Fields(s), "_")
}

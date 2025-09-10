package utils

import "fmt"


func FormatRoles(roles []string) string {

	var output []string

	for _,role := range roles {
		output = append(output , fmt.Sprintf("\"%s\"", role))

	}
	return strings.Join(output , ",")
}

func stringSliceToInterface(slice []string) []interface{} {
	res := make([]interface{}, len(slice))
	for i, v := range slice {
		res[i] = v

	}
	return res
}
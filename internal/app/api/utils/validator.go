// Package utils for custom utilits like as custom validator and etc.
package utils

import "regexp"

func ValidUUID(id string) bool {
	re := regexp.MustCompile(`[a-fA-F\d]{8}-[a-fA-F\d]{4}-[a-fA-F\d]{4}-[a-fA-F\d]{4}-[a-fA-F\d]{12}$`)
	return re.MatchString(id)
}

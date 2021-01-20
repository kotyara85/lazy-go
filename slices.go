package lazygo

import (
	"strings"
	"time"
)

// SliceUniq makes sure we have only uniq elements in slice
func SliceUniq(slice *[]string) {
	keys := make(map[string]bool)
	p := make([]string, 0)
	for _, el := range *slice {
		if _, val := keys[strings.ToLower(el)]; !val {
			keys[strings.ToLower(el)] = true
			p = append(p, el)
		}
	}
	*slice = p
}

// SliceRemoveEmpty removes empty elements from slice
func SliceRemoveEmpty(slice *[]string) {
	i := 0
	p := *slice
	for _, entry := range p {
		if strings.Trim(entry, " ") != "" {
			p[i] = strings.Trim(entry, " ")
			i++
		}
	}
	*slice = p[0:i]
}

// SliceInclude checks if a specific values is present in slice
func SliceInclude(s, v interface{}) bool {
	switch s.(type) {
	case []time.Duration:
		for _, kk := range s.([]time.Duration) {
			if kk == v {
				return true
			}
		}
	case []int:
		for _, kk := range s.([]int) {
			if kk == v {
				return true
			}
		}
	case []int64:
		for _, kk := range s.([]int64) {
			if kk == v {
				return true
			}
		}
	case []string:
		for _, kk := range s.([]string) {
			if strings.ToLower(kk) == v {
				return true
			}
		}
	}
	return false
}

// SliceToInterface converts []string{} to []interface{}
func SliceToInterface(in []string) []interface{} {
	out := make([]interface{}, len(in))
	for i, y := range in {
		out[i] = y
	}
	return out
}

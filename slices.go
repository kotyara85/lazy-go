package lazygo

import "strings"

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

// SliceInclude checks if a specific el is present in slice
func SliceInclude(slice []string, str string) bool {
	for _, element := range slice {
		if strings.ToLower(element) == strings.ToLower(str) {
			return true
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

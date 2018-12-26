package surrogate

import "strings"

func Sub(source string, replacement map[string]string) []string {
	result := make([]string, 0)
	source = strings.ToLower(source)
	for k, v := range replacement {
		replacement[strings.ToLower(k)] = strings.ToLower(v)
	}
	for _, c := range source {
		rr := make([]string, 0)
		if v, ok := replacement[string(c)]; ok {
			if len(result) > 0 {
				for _, cx := range result {
					rr = append(rr, cx+string(c))
					rr = append(rr, cx+v)
				}
			} else {
				rr = append(rr, string(c))
				rr = append(rr, v)
			}
		} else {
			for _, cx := range result {
				rr = append(rr, cx+string(c))
			}
		}
		result = make([]string, len(rr))
		copy(result, rr)
	}
	return removeDuplicates(result)
}

func SubMultiMap(source string, replacement map[string][]string) []string {
	result := make([]string, 0)
	source = strings.ToLower(source)
	for _, c := range source {
		rr := make([]string, 0)
		if v, ok := replacement[string(c)]; ok {
			for _, val := range v {
				if len(result) > 0 {
					for _, cx := range result {
						rr = append(rr, cx+string(c))
						rr = append(rr, cx+val)
					}
				} else {
					rr = append(rr, string(c))
					rr = append(rr, val)
				}
			}
		} else {
			if len(result) > 0 {
				for _, cx := range result {
					rr = append(rr, cx+string(c))
				}
			} else {
				rr = append(rr, string(c))
			}
		}
		result = make([]string, len(rr))
		copy(result, rr)
	}
	return removeDuplicates(result)
}
func SubCaseSensitive(source string, replacement map[string]string) []string {
	result := make([]string, 0)
	for _, c := range source {
		rr := make([]string, 0)
		if v, ok := replacement[string(c)]; ok {
			if len(result) > 0 {
				for _, cx := range result {
					rr = append(rr, cx+string(c))
					rr = append(rr, cx+v)
				}
			} else {
				rr = append(rr, string(c))
				rr = append(rr, v)
			}
		} else {
			for _, cx := range result {
				rr = append(rr, cx+string(c))
			}
		}
		result = make([]string, len(rr))
		copy(result, rr)
	}
	return removeDuplicates(result)
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	var result []string

	for v := range elements {
		if encountered[elements[v]] {
			continue
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

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
	return result
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
	return result
}

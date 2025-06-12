//go:build !solution

package hogwarts

func GetCourseList(prereqs map[string][]string) []string {
	result := make([]string, 0)

	prereqsMap := make(map[string]map[string]struct{}, 0)

	for course, reqs := range prereqs {
		if _, ok := prereqsMap[course]; !ok {
			prereqsMap[course] = make(map[string]struct{}, 0)
		}
		for _, req := range reqs {
			if _, ok := prereqsMap[req]; !ok {
				prereqsMap[req] = make(map[string]struct{}, 0)
			}
			prereqsMap[course][req] = struct{}{}
		}
	}

	for len(prereqsMap) > 0 {
		batch := make([]string, 0)
		for course, reqsSet := range prereqsMap {
			if len(reqsSet) == 0 {
				result = append(result, course)
				batch = append(batch, course)
				delete(prereqsMap, course)
			}
		}
		if len(batch) == 0 {
			panic("Cycle found")
		}
		for _, reqsSet := range prereqsMap {
			for _, course := range batch {
				delete(reqsSet, course)
			}
		}
	}

	return result
}

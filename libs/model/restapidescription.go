package model

type RestApiDescription struct {
	Url         string `json:"url"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Location    string `json:"location"`
	IsActive    bool   `json:"active"`
	Id          string `json:"id"`
}

type RestApiDescriptions []RestApiDescription

func Where(list []RestApiDescription, predicate func(RestApiDescription) bool) []RestApiDescription {
	if predicate == nil {
		return list
	}

	var results = []RestApiDescription{}

	for j := 0; j < len(list); j++ {
		if predicate(list[j]) {
			results = append(results, list[j])
		}
	}

	return results
}

func Skip(list []RestApiDescription, number int) []RestApiDescription {
	var results = []RestApiDescription{}
	if number < 0 {
		return results
	}

	for j := number; j < len(list); j++ {
		results = append(results, list[j])
	}

	return results
}

func Take(list []RestApiDescription, number int) []RestApiDescription {
	var results = []RestApiDescription{}
	for j := 0; j < len(list) && j < number; j++ {
		results = append(results, list[j])
	}

	return results
}

func Find(list []RestApiDescription, id string) int {
	for j := 0; j < len(list); j++ {
		if list[j].Id == id {
			return j
		}
	}

	return -1
}

func Replace(list []RestApiDescription, item RestApiDescription) []RestApiDescription {
	for j := 0; j < len(list); j++ {
		if list[j].Id == item.Id {
			list[j] = item
		}
	}

	return list
}

func Remove(list []RestApiDescription, index int) []RestApiDescription {
	var newList = []RestApiDescription{}

	for j := 0; j < len(list); j++ {
		if j != index {
			newList = append(newList, list[j])
		}
	}

	return newList
}

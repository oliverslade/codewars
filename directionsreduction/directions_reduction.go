package directionsreduction

func DirReduc(arr []string) []string {
	var result []string

	oppositesMap := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	for _, direction := range arr {
		if len(result) > 0 && oppositesMap[direction] == result[len(result)-1] {
			result = result[:len(result)-1]
		} else {
			result = append(result, direction)
		}
	}

	return result
}

package sumofintervals

import "sort"

func SumOfIntervals(intervals [][2]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := 0
	mergedIntervals := make([][2]int, 0)

	currentSmaller := intervals[0][0]
	currentLarger := intervals[0][1]

	for _, arrayPairs := range intervals[1:] {
		smallerNumber := arrayPairs[0]
		largerNumber := arrayPairs[1]

		if currentLarger >= smallerNumber {
			if largerNumber > currentLarger {
				currentLarger = largerNumber
			}
		} else {
			mergedIntervals = append(mergedIntervals, [2]int{currentSmaller, currentLarger})
			currentSmaller = smallerNumber
			currentLarger = largerNumber
		}
	}

	mergedIntervals = append(mergedIntervals, [2]int{currentSmaller, currentLarger})

	for _, arrayPairs := range mergedIntervals {
		largerNumber := arrayPairs[1]
		smallerNumber := arrayPairs[0]
		result += (largerNumber - smallerNumber)
	}

	return result
}

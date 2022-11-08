package sentences

import (
	"math"
	"strings"
)

// export of the below minDistance function
func EditDistance(s1, s2 string) int {
	return minDistance(s1, s2)
}

// adapted for words and sentences from the code here:
// https://golangbyexample.com/edit-distance-two-strings-golang/
func minDistance(s1 string, s2 string) int {
	sentenceOneWords := strings.Split(s1, " ")
	sentenceTwoWords := strings.Split(s2, " ")
	lenSent1 := len(sentenceOneWords)
	lenSent2 := len(sentenceTwoWords)

	editDistanceMatrix := make([][]int, lenSent1+1)

	for i := range editDistanceMatrix {
		editDistanceMatrix[i] = make([]int, lenSent2+1)
	}

	for i := 1; i <= lenSent2; i++ {
		editDistanceMatrix[0][i] = i
	}

	for i := 1; i <= lenSent1; i++ {
		editDistanceMatrix[i][0] = i
	}
	for i := 1; i <= lenSent1; i++ {
		for j := 1; j <= lenSent2; j++ {

			if sentenceOneWords[i-1] == sentenceTwoWords[j-1] {
				editDistanceMatrix[i][j] = editDistanceMatrix[i-1][j-1]
			} else {
				// modified here to not count a replace as a distance of 1
				editDistanceMatrix[i][j] = 1 + int(math.Min(float64(editDistanceMatrix[i-1][j]), float64(editDistanceMatrix[i][j-1])))
			}
		}
	}
	return editDistanceMatrix[lenSent1][lenSent2]
}

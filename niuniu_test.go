package niuniu

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNiuniuCartType(t *testing.T) {
	originResults := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}, {26, 27, 28, 29, 30}, {31, 32, 33, 34, 35}, {36, 37, 38, 39, 40}, {41, 42, 43, 44, 45}, {46, 47, 48, 53, 54}}
	// originResults := [][]int{{2, 3, 9, 10, 13}, {41, 42, 48, 49, 51}, {15, 16, 22, 23, 26}, {28, 29, 35, 36, 38}}
	// originResults := [][]int{{2, 3, 9, 10, 12}, {41, 42, 48, 49, 51}, {15, 16, 22, 23, 25}, {28, 29, 35, 36, 38}}
	rankResult := make(map[int][]int)
	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

	n, err := New(cts)
	if err != nil {
		log.Println(err)
	}

	rankResult[1] = []int{5, 1, 3}
	rankResult[2] = []int{7}
	rankResult[4] = []int{2}
	rankResult[6] = []int{6, 0}
	rankResult[8] = []int{4}
	rankResult[11] = []int{9, 8}

	result := n.CompareResult(originResults)
	assert.Equal(t, result, rankResult)
	log.Println(result)
}

func TestCompareType(t *testing.T) {
	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
	scResult := make(map[int]int)
	scResult[0] = 1
	scResult[1] = 2
	scResult[2] = 3
	scResult[3] = 1
	scResult[4] = 2

	cardType := CompareType(scResult, cts)

	assert.Equal(t, "sc", cardType)
}

func TestCompareCow(t *testing.T) {
	scResult := make(map[int]int)
	scResult[0] = 1
	scResult[1] = 9
	scResult[2] = 10
	scResult[3] = 4
	scResult[4] = 2

	cardType := CompareCow(scResult)
	assert.Equal(t, "c6", cardType)
}

func TestGetMaxCard(t *testing.T) {
	mResult := make(map[int]map[int]int)
	oneResult := make(map[int]int)
	twoResult := make(map[int]int)
	oneResult[1] = 1
	oneResult[2] = 2
	oneResult[3] = 3
	oneResult[4] = 4
	oneResult[5] = 5

	twoResult[31] = 5
	twoResult[32] = 6
	twoResult[33] = 7
	twoResult[34] = 8
	twoResult[35] = 9

	mResult[0] = oneResult
	mResult[6] = twoResult

	sMax := GetMaxCard(mResult)

	assert.Equal(t, []int{6, 0}, sMax)
}

func TestChangeCard(t *testing.T) {
	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
	// 1 2 9 8 13
	origin := []int{1, 28, 48, 53, 54}
	mResult := make(map[int]int)
	mResult[1] = 1
	mResult[28] = 2
	mResult[48] = 9
	mResult[53] = 8
	mResult[54] = 13

	result := ChangeCard(origin, cts)

	assert.Equal(t, mResult, result)
}

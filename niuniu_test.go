package niuniu

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNiuniuCartType(t *testing.T) {
	t.Run("case_1", func(t *testing.T) {
		originResults := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}, {26, 27, 28, 29, 30}, {31, 32, 33, 34, 35}, {36, 37, 38, 39, 40}, {41, 42, 43, 44, 45}, {46, 47, 48, 53, 54}}
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
	})
	//同牌型 多副牌兩兩同點
	t.Run("case_2", func(t *testing.T) {
		originResults := [][]int{{2, 3, 9, 10, 13}, {41, 42, 48, 49, 51}, {15, 16, 22, 23, 26}, {28, 29, 35, 36, 38}}
		rankResult := make(map[int][]int)
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

		n, err := New(cts)
		if err != nil {
			log.Println(err)
		}

		rankResult[1] = []int{2, 0, 1, 3}

		result := n.CompareResult(originResults)
		assert.Equal(t, result, rankResult)
	})
	//同牌型 多副牌全同點
	t.Run("case_3", func(t *testing.T) {
		originResults := [][]int{{2, 3, 9, 10, 12}, {41, 42, 48, 49, 51}, {15, 16, 22, 23, 25}, {28, 29, 35, 36, 38}}
		rankResult := make(map[int][]int)
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

		n, err := New(cts)
		if err != nil {
			log.Println(err)
		}

		rankResult[1] = []int{1, 3, 2, 0}

		result := n.CompareResult(originResults)
		assert.Equal(t, result, rankResult)
	})
	t.Run("case_4", func(t *testing.T) {
		//多副牌兩兩同點 有鬼牌
		originResults := [][]int{{2, 3, 9, 10, 53}, {41, 42, 48, 49, 51}, {15, 16, 22, 23, 54}, {28, 29, 35, 36, 38}}
		rankResult := make(map[int][]int)
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

		n, err := New(cts)
		if err != nil {
			log.Println(err)
		}

		rankResult[1] = []int{1, 3}
		rankResult[10] = []int{2, 0}

		result := n.CompareResult(originResults)
		assert.Equal(t, result, rankResult)
	})
	//同牌型 多副牌同副牌兩張相同最大值
	t.Run("case_5", func(t *testing.T) {
		originResults := [][]int{{2, 3, 9, 51, 12}, {15, 16, 22, 38, 25}}
		rankResult := make(map[int][]int)
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

		n, err := New(cts)
		if err != nil {
			log.Println(err)
		}

		rankResult[1] = []int{0, 1}

		result := n.CompareResult(originResults)
		assert.Equal(t, result, rankResult)
	})
	//多副牌都不相同點數
	t.Run("case_6", func(t *testing.T) {
		originResults := [][]int{{2, 3, 9, 10, 13}, {41, 42, 48, 49, 50}, {15, 16, 22, 23, 25}, {28, 29, 35, 36, 38}}
		rankResult := make(map[int][]int)
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

		n, err := New(cts)
		if err != nil {
			log.Println(err)
		}

		rankResult[1] = []int{0, 3, 2, 1}

		result := n.CompareResult(originResults)
		assert.Equal(t, result, rankResult)
	})
	//兩兩同牌型 多副牌兩兩同點
	t.Run("case_7", func(t *testing.T) {
		originResults := [][]int{{2, 7, 9, 10, 14}, {41, 42, 48, 49, 51}, {15, 20, 22, 23, 1}, {28, 29, 35, 36, 38}}
		rankResult := make(map[int][]int)
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

		n, err := New(cts)
		if err != nil {
			log.Println(err)
		}

		rankResult[10] = []int{2, 0}
		rankResult[1] = []int{1, 3}

		result := n.CompareResult(originResults)
		assert.Equal(t, result, rankResult)
	})
	//多副牌都相同點數 有鬼牌
	t.Run("case_8", func(t *testing.T) {
		originResults := [][]int{{2, 7, 9, 10, 14}, {41, 42, 48, 49, 53}, {15, 20, 22, 23, 1}, {28, 29, 35, 36, 54}}
		rankResult := make(map[int][]int)
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

		n, err := New(cts)
		if err != nil {
			log.Println(err)
		}

		rankResult[10] = []int{1, 3, 2, 0}

		result := n.CompareResult(originResults)
		assert.Equal(t, result, rankResult)
	})
	//多副牌 有兩張鬼牌在同一副牌
	t.Run("case_9", func(t *testing.T) {
		originResults := [][]int{{2, 7, 9, 10, 14}, {41, 47, 48, 54, 53}, {15, 20, 22, 23, 1}, {28, 33, 35, 36, 27}}
		rankResult := make(map[int][]int)
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

		n, err := New(cts)
		if err != nil {
			log.Println(err)
		}

		rankResult[11] = []int{1}
		rankResult[10] = []int{3, 2, 0}

		result := n.CompareResult(originResults)
		assert.Equal(t, result, rankResult)
	})
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

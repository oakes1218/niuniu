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
	//多副牌 同牌形同數字鬼牌不是最大數字
	t.Run("case_10", func(t *testing.T) {
		originResults := [][]int{{2, 3, 9, 10, 18}, {41, 42, 48, 49, 53}, {15, 16, 22, 23, 5}, {28, 29, 35, 36, 54}}
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
	//多副牌 同牌形同數字
	t.Run("case_11", func(t *testing.T) {
		originResults := [][]int{{1, 14, 27, 40, 37}, {2, 15, 28, 41, 50}, {3, 16, 29, 42, 24}, {4, 17, 30, 43, 11}}
		rankResult := make(map[int][]int)
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

		n, err := New(cts)
		if err != nil {
			log.Println(err)
		}

		rankResult[14] = []int{1, 0, 2, 3}

		result := n.CompareResult(originResults)
		assert.Equal(t, result, rankResult)
	})
}

func TestCompareType(t *testing.T) {
	t.Run("case_1", func(t *testing.T) {
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
		scResult := make(map[int]int)
		scResult[0] = 1
		scResult[1] = 2
		scResult[2] = 3
		scResult[3] = 1
		scResult[4] = 2

		cardType := CompareType(scResult, cts)

		assert.Equal(t, "sc", cardType)
	})
	t.Run("case_2", func(t *testing.T) {
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
		ffResult := make(map[int]int)
		ffResult[0] = 7
		ffResult[1] = 7
		ffResult[2] = 7
		ffResult[3] = 7
		ffResult[4] = 13

		cardType := CompareType(ffResult, cts)

		assert.Equal(t, "ff", cardType)
	})
	t.Run("case_3", func(t *testing.T) {
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
		f5Result := make(map[int]int)
		f5Result[0] = 11
		f5Result[1] = 11
		f5Result[2] = 12
		f5Result[3] = 13
		f5Result[4] = 13

		cardType := CompareType(f5Result, cts)

		assert.Equal(t, "5f", cardType)
	})
	t.Run("case_4", func(t *testing.T) {
		cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
		f4Result := make(map[int]int)
		f4Result[0] = 10
		f4Result[1] = 11
		f4Result[2] = 12
		f4Result[3] = 13
		f4Result[4] = 11

		cardType := CompareType(f4Result, cts)

		assert.Equal(t, "4f", cardType)
	})
}

func TestCompareCow(t *testing.T) {
	t.Run("case_1", func(t *testing.T) {
		c1Result := make(map[int]int)
		c1Result[0] = 1
		c1Result[1] = 9
		c1Result[2] = 10
		c1Result[3] = 10
		c1Result[4] = 1

		cardType := CompareCow(c1Result)
		assert.Equal(t, "c1", cardType)
	})
	t.Run("case_2", func(t *testing.T) {
		c2Result := make(map[int]int)
		c2Result[0] = 11
		c2Result[1] = 5
		c2Result[2] = 3
		c2Result[3] = 2
		c2Result[4] = 2

		cardType := CompareCow(c2Result)
		assert.Equal(t, "c2", cardType)
	})
	t.Run("case_3", func(t *testing.T) {
		c3Result := make(map[int]int)
		c3Result[0] = 6
		c3Result[1] = 7
		c3Result[2] = 12
		c3Result[3] = 8
		c3Result[4] = 2

		cardType := CompareCow(c3Result)
		assert.Equal(t, "c3", cardType)
	})
	t.Run("case_4", func(t *testing.T) {
		c4Result := make(map[int]int)
		c4Result[0] = 6
		c4Result[1] = 9
		c4Result[2] = 9
		c4Result[3] = 5
		c4Result[4] = 5

		cardType := CompareCow(c4Result)
		assert.Equal(t, "c4", cardType)
	})
	t.Run("case_5", func(t *testing.T) {
		c5Result := make(map[int]int)
		c5Result[0] = 6
		c5Result[1] = 9
		c5Result[2] = 10
		c5Result[3] = 5
		c5Result[4] = 5

		cardType := CompareCow(c5Result)
		assert.Equal(t, "c5", cardType)
	})
	t.Run("case_6", func(t *testing.T) {
		c6Result := make(map[int]int)
		c6Result[0] = 1
		c6Result[1] = 9
		c6Result[2] = 10
		c6Result[3] = 4
		c6Result[4] = 2

		cardType := CompareCow(c6Result)
		assert.Equal(t, "c6", cardType)
	})
	t.Run("case_7", func(t *testing.T) {
		c7Result := make(map[int]int)
		c7Result[0] = 7
		c7Result[1] = 13
		c7Result[2] = 10
		c7Result[3] = 11
		c7Result[4] = 12

		cardType := CompareCow(c7Result)
		assert.Equal(t, "c7", cardType)
	})
	t.Run("case_8", func(t *testing.T) {
		c8Result := make(map[int]int)
		c8Result[0] = 3
		c8Result[1] = 9
		c8Result[2] = 7
		c8Result[3] = 1
		c8Result[4] = 8

		cardType := CompareCow(c8Result)
		assert.Equal(t, "c8", cardType)
	})
	t.Run("case_9", func(t *testing.T) {
		c9Result := make(map[int]int)
		c9Result[0] = 2
		c9Result[1] = 7
		c9Result[2] = 9
		c9Result[3] = 10
		c9Result[4] = 1

		cardType := CompareCow(c9Result)

		assert.Equal(t, "c9", cardType)
	})
	//牛牛
	t.Run("case_10", func(t *testing.T) {
		ccResult := make(map[int]int)
		ccResult[0] = 2
		ccResult[1] = 7
		ccResult[2] = 9
		ccResult[3] = 1
		ccResult[4] = 1

		cardType := CompareCow(ccResult)

		assert.Equal(t, "cc", cardType)
	})
}

func TestGetMaxCard(t *testing.T) {
	t.Run("case_1", func(t *testing.T) {
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
	})
	//多副牌 兩兩相同牌行
	t.Run("case_2", func(t *testing.T) {
		mResult := make(map[int]map[int]int)
		oneResult := make(map[int]int)
		twoResult := make(map[int]int)
		threeResult := make(map[int]int)
		fourResult := make(map[int]int)
		fiveResult := make(map[int]int)
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

		threeResult[18] = 5
		threeResult[19] = 6
		threeResult[20] = 7
		threeResult[21] = 8
		threeResult[22] = 9

		fourResult[14] = 1
		fourResult[15] = 2
		fourResult[16] = 3
		fourResult[17] = 4
		fourResult[44] = 5

		fiveResult[48] = 9
		fiveResult[49] = 10
		fiveResult[50] = 11
		fiveResult[51] = 12
		fiveResult[52] = 13

		mResult[0] = oneResult
		mResult[1] = twoResult
		mResult[2] = threeResult
		mResult[3] = fourResult
		mResult[4] = fiveResult

		sMax := GetMaxCard(mResult)

		assert.Equal(t, []int{4, 1, 2, 3, 0}, sMax)
	})
	//相同牌型 穿插不同花色
	t.Run("case_3", func(t *testing.T) {
		mResult := make(map[int]map[int]int)
		oneResult := make(map[int]int)
		twoResult := make(map[int]int)

		oneResult[1] = 1
		oneResult[14] = 3
		oneResult[17] = 6
		oneResult[25] = 12
		oneResult[26] = 13

		twoResult[14] = 1
		twoResult[3] = 3
		twoResult[6] = 6
		twoResult[12] = 12
		twoResult[13] = 13

		mResult[0] = oneResult
		mResult[1] = twoResult

		sMax := GetMaxCard(mResult)

		assert.Equal(t, []int{0, 1}, sMax)
	})
}

func TestChangeCard(t *testing.T) {
	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
	originResults := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}

	mResult := make(map[int]int)
	mResult[1] = 1
	mResult[2] = 2
	mResult[3] = 3
	mResult[4] = 4
	mResult[5] = 5
	mResult[6] = 6
	mResult[7] = 7
	mResult[8] = 8
	mResult[9] = 9
	mResult[10] = 10
	mResult[11] = 11
	mResult[12] = 12
	mResult[13] = 13
	mResult[14] = 1
	mResult[15] = 2
	mResult[16] = 3
	mResult[17] = 4
	mResult[18] = 5
	mResult[19] = 6
	mResult[20] = 7
	mResult[21] = 8
	mResult[22] = 9
	mResult[23] = 10
	mResult[24] = 11
	mResult[25] = 12
	mResult[26] = 13
	mResult[27] = 1
	mResult[28] = 2
	mResult[29] = 3
	mResult[30] = 4
	mResult[31] = 5
	mResult[32] = 6
	mResult[33] = 7
	mResult[34] = 8
	mResult[35] = 9
	mResult[36] = 10
	mResult[37] = 11
	mResult[38] = 12
	mResult[39] = 13
	mResult[40] = 1
	mResult[41] = 2
	mResult[42] = 3
	mResult[43] = 4
	mResult[44] = 5
	mResult[45] = 6
	mResult[46] = 7
	mResult[47] = 8
	mResult[48] = 9
	mResult[49] = 10
	mResult[50] = 11
	mResult[51] = 12
	mResult[52] = 13

	result := ChangeCard(originResults, cts)

	assert.Equal(t, mResult, result)
}

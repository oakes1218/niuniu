package niuniu

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInlottoNiuNiu(t *testing.T) {
	cts := []string{"cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
	n, err := New(cts)
	if err != nil {
		log.Println(err)
	}

	t.Run("藍牛牛 紅無牛 => 藍贏", func(t *testing.T) {
		cards := [][]int{
			{1, 3, 6, 12, 13},    // ♦A ♦3 ♦6 ♦Q ♦K
			{27, 36, 38, 44, 41}, // ♥A ♥10 ♥Q ♠5 ♠2
		}

		r := n.CompareResult(cards)
		assert.Equal(t, 0, r[11][0])
		assert.Equal(t, 1, r[1][0])
	})

	t.Run("藍牛4 紅牛5 => 紅贏", func(t *testing.T) {
		cards := [][]int{
			{18, 28, 42, 2, 15}, // ♣5 ♥2 ♠3 ♦2 ♣2
			{1, 43, 44, 41, 29}, // ♦A ♠4 ♠5 ♠2 ♥3
		}

		r := n.CompareResult(cards)
		assert.Equal(t, 0, r[5][0])
		assert.Equal(t, 1, r[6][0])
	})

	t.Run("藍牛1(最大牌Q) 紅牛1(最大牌K) => 紅贏", func(t *testing.T) {
		cards := [][]int{
			{2, 3, 5, 12, 14},    // ♦2 ♦3 ♦5 ♦Q ♣A
			{27, 30, 44, 52, 40}, // ♥A ♥4 ♠5 ♠K ♠A
		}

		r := n.CompareResult(cards)
		assert.Equal(t, 1, r[2][0])
		assert.Equal(t, 0, r[2][1])
	})

	t.Run("藍無牛 (最大牌8) 紅無牛 (最大牌A) => 藍贏", func(t *testing.T) {
		cards := [][]int{
			{10, 11, 1, 2, 5},    // ♦10 ♦J ♦A ♦2 ♦5
			{27, 29, 44, 46, 40}, // ♥A ♥3 ♠5 ♠7 ♠A
		}

		r := n.CompareResult(cards)
		assert.Equal(t, 0, r[1][0])
		assert.Equal(t, 1, r[1][1])
	})

	t.Run("藍無牛 (最大牌6) 紅無牛 (最大牌7) => 紅贏", func(t *testing.T) {
		cards := [][]int{
			{15, 6, 1, 4, 27},    // ♣2 ♦6 ♦A ♦4 ♥A
			{31, 29, 44, 46, 40}, // ♥5 ♥3 ♠5 ♠7 ♠A
		}

		r := n.CompareResult(cards)
		assert.Equal(t, 1, r[1][0])
		assert.Equal(t, 0, r[1][1])
	})

	t.Run("藍牛3 (最大牌♠K) 紅牛3 (最大牌♦K) => 藍贏", func(t *testing.T) {
		cards := [][]int{
			{52, 51, 50, 49, 42}, // ♠K ♠Q  ♠J ♠10 ♠3
			{13, 12, 11, 10, 3},  // ♦K ♦10 ♦J ♦Q  ♦3
		}

		r := n.CompareResult(cards)
		assert.Equal(t, 0, r[4][0])
		assert.Equal(t, 1, r[4][1])
	})

	t.Run("藍牛7 (最大牌♣K) 紅牛7 (最大牌♥K) => 紅贏", func(t *testing.T) {
		cards := [][]int{
			{14, 19, 23, 25, 26}, // ♣A ♣6 ♣10 ♣Q ♣K
			{20, 31, 28, 29, 39}, // ♣7 ♥5 ♥2  ♥3 ♥K
		}

		r := n.CompareResult(cards)
		assert.Equal(t, 1, r[8][0])
		assert.Equal(t, 0, r[8][1])
	})

	t.Run("藍牛6 (最大牌♥9) 紅牛6 (最大牌♣9) => 藍贏", func(t *testing.T) {
		cards := [][]int{
			{2, 6, 15, 7, 35},    // ♦2 ♦6 ♣2 ♦7 ♥9
			{22, 33, 43, 42, 29}, // ♣9 ♥7 ♠4 ♠3 ♥3
		}

		r := n.CompareResult(cards)
		assert.Equal(t, 0, r[7][0])
		assert.Equal(t, 1, r[7][1])
	})

	t.Run("藍牛8 (最大牌♥Q) 紅牛8 (最大牌♦Q) => 藍贏", func(t *testing.T) {
		cards := [][]int{
			{3, 5, 24, 36, 38}, // ♦3 ♣5 ♣J ♥10 ♥Q
			{43, 4, 31, 5, 12}, // ♠4 ♦4 ♥5 ♦5 ♦Q
		}

		r := n.CompareResult(cards)
		assert.Equal(t, 0, r[9][0])
		assert.Equal(t, 1, r[9][1])
	})
}

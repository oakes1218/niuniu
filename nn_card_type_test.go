package niuniu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//五小牛
func TestSC(t *testing.T) {
	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
	r := make(map[int]int)
	r[0] = 1
	r[1] = 1
	r[2] = 1
	r[3] = 1
	r[4] = 3

	cardType := CompareType(r, cts)
	assert.Equal(t, "sc", cardType)

	r = make(map[int]int)
	r[0] = 3
	r[1] = 4
	r[2] = 2
	r[3] = 2
	r[4] = 3

	cardType2 := CompareType(r, cts)
	assert.NotEqual(t, "sc", cardType2)
}

//四炸
func TestFF(t *testing.T) {
	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}

	r := make(map[int]int)
	r[0] = 1
	r[1] = 1
	r[2] = 1
	r[3] = 1
	r[4] = 7

	cardType := CompareType(r, cts)
	assert.Equal(t, "ff", cardType)

	r = make(map[int]int)
	r[0] = 11
	r[1] = 11
	r[2] = 11
	r[3] = 11
	r[4] = 13

	cardType2 := CompareType(r, cts)
	assert.Equal(t, "ff", cardType2)

	r = make(map[int]int)
	r[0] = 2
	r[1] = 2
	r[2] = 2
	r[3] = 2
	r[4] = 1

	cardType3 := CompareType(r, cts)
	assert.NotEqual(t, "ff", cardType3)
}

//五花牛
func Test5F(t *testing.T) {
	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
	r := make(map[int]int)
	r[0] = 13
	r[1] = 12
	r[2] = 11
	r[3] = 13
	r[4] = 11

	cardType := CompareType(r, cts)
	assert.Equal(t, "5f", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 12
	r[2] = 11
	r[3] = 13
	r[4] = 11

	cardType2 := CompareType(r, cts)
	assert.NotEqual(t, "5f", cardType2)
}

//四花牛
func Test4F(t *testing.T) {
	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
	r := make(map[int]int)
	r[0] = 13
	r[1] = 12
	r[2] = 11
	r[3] = 13
	r[4] = 10

	cardType := CompareType(r, cts)
	assert.Equal(t, "4f", cardType)

	r = make(map[int]int)
	r[0] = 13
	r[1] = 12
	r[2] = 11
	r[3] = 13
	r[4] = 11

	cardType2 := CompareType(r, cts)
	assert.NotEqual(t, "4f", cardType2)
}

//牛牛
func TestCC(t *testing.T) {
	// 10, 10
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 5
	r[4] = 5

	cardType := CompareCow(r)
	assert.Equal(t, "cc", cardType)

	// 30, 20
	r = make(map[int]int)
	r[0] = 13
	r[1] = 12
	r[2] = 11
	r[3] = 13
	r[4] = 11

	cardType2 := CompareCow(r)
	assert.Equal(t, "cc", cardType2)

	// 30, 10
	r = make(map[int]int)
	r[0] = 13
	r[1] = 12
	r[2] = 11
	r[3] = 4
	r[4] = 6

	cardType3 := CompareCow(r)
	assert.Equal(t, "cc", cardType3)

	// 20, 10
	r = make(map[int]int)
	r[0] = 9
	r[1] = 1
	r[2] = 10
	r[3] = 4
	r[4] = 6

	cardType4 := CompareCow(r)
	assert.Equal(t, "cc", cardType4)
}

//牛9
func TestC9(t *testing.T) {
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 11
	r[4] = 9

	cardType := CompareCow(r)
	assert.Equal(t, "c9", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 13
	r[3] = 13
	r[4] = 9

	cardType2 := CompareCow(r)
	assert.Equal(t, "c9", cardType2)
}

//牛8
func TestC8(t *testing.T) {
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 11
	r[4] = 8

	cardType := CompareCow(r)
	assert.Equal(t, "c8", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 13
	r[3] = 9
	r[4] = 9

	cardType2 := CompareCow(r)
	assert.Equal(t, "c8", cardType2)
}

//牛7
func TestC7(t *testing.T) {
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 11
	r[4] = 7

	cardType := CompareCow(r)
	assert.Equal(t, "c7", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 13
	r[3] = 9
	r[4] = 8

	cardType2 := CompareCow(r)
	assert.Equal(t, "c7", cardType2)
}

//牛6
func TestC6(t *testing.T) {
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 11
	r[4] = 6

	cardType := CompareCow(r)
	assert.Equal(t, "c6", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 13
	r[3] = 9
	r[4] = 7

	cardType2 := CompareCow(r)
	assert.Equal(t, "c6", cardType2)
}

//牛5
func TestC5(t *testing.T) {
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 11
	r[4] = 5

	cardType := CompareCow(r)
	assert.Equal(t, "c5", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 13
	r[3] = 9
	r[4] = 6

	cardType2 := CompareCow(r)
	assert.Equal(t, "c5", cardType2)
}

//牛4
func TestC4(t *testing.T) {
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 11
	r[4] = 4

	cardType := CompareCow(r)
	assert.Equal(t, "c4", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 13
	r[3] = 8
	r[4] = 6

	cardType2 := CompareCow(r)
	assert.Equal(t, "c4", cardType2)

	r = make(map[int]int)
	r[0] = 6
	r[1] = 3
	r[2] = 1
	r[3] = 7
	r[4] = 7

	cardType3 := CompareCow(r)
	assert.Equal(t, "c4", cardType3)
}

//牛3
func TestC3(t *testing.T) {
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 11
	r[4] = 3

	cardType := CompareCow(r)
	assert.Equal(t, "c3", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 13
	r[3] = 7
	r[4] = 6

	cardType2 := CompareCow(r)
	assert.Equal(t, "c3", cardType2)

	r = make(map[int]int)
	r[0] = 6
	r[1] = 3
	r[2] = 1
	r[3] = 4
	r[4] = 9

	cardType3 := CompareCow(r)
	assert.Equal(t, "c3", cardType3)
}

//牛2
func TestC2(t *testing.T) {
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 11
	r[4] = 2

	cardType := CompareCow(r)
	assert.Equal(t, "c2", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 13
	r[3] = 1
	r[4] = 1

	cardType2 := CompareCow(r)
	assert.Equal(t, "c2", cardType2)

	r = make(map[int]int)
	r[0] = 6
	r[1] = 3
	r[2] = 1
	r[3] = 6
	r[4] = 6

	cardType3 := CompareCow(r)
	assert.Equal(t, "c2", cardType3)
}

//牛1
func TestC1(t *testing.T) {
	r := make(map[int]int)
	r[0] = 2
	r[1] = 3
	r[2] = 5
	r[3] = 11
	r[4] = 1

	cardType := CompareCow(r)
	assert.Equal(t, "c1", cardType)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 13
	r[3] = 2
	r[4] = 9

	cardType2 := CompareCow(r)
	assert.Equal(t, "c1", cardType2)

	r = make(map[int]int)
	r[0] = 6
	r[1] = 3
	r[2] = 1
	r[3] = 4
	r[4] = 7

	cardType3 := CompareCow(r)
	assert.Equal(t, "c1", cardType3)
}

//無牛
func TestC0(t *testing.T) {
	r := make(map[int]int)
	r[0] = 13
	r[1] = 3
	r[2] = 1
	r[3] = 1
	r[4] = 2

	cardType := CompareCow(r)
	assert.Equal(t, "c0", cardType)

	r = make(map[int]int)
	r[0] = 1
	r[1] = 1
	r[2] = 1
	r[3] = 2
	r[4] = 2

	cardType2 := CompareCow(r)
	assert.Equal(t, "c0", cardType2)

	r = make(map[int]int)
	r[0] = 10
	r[1] = 11
	r[2] = 7
	r[3] = 1
	r[4] = 1

	cardType3 := CompareCow(r)
	assert.Equal(t, "c0", cardType3)

	r = make(map[int]int)
	r[0] = 1
	r[1] = 2
	r[2] = 3
	r[3] = 12
	r[4] = 13

	cardType4 := CompareCow(r)
	assert.Equal(t, "c0", cardType4)
}
package niuniu

import (
	"errors"
	"sort"
	"strconv"
)

var (
	CardTypeWhite = []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
)

type CardType struct {
	cardTypeS []string
}

func New(CardTypeSlice []string) (c *CardType, err error) {
	if len(CardTypeSlice) == 0 {
		err = errors.New("遊戲type不得為空")
		return
	}

	for _, v1 := range CardTypeSlice {
		var count int
		for _, v2 := range CardTypeWhite {
			if v1 == v2 {
				count++
			}
		}

		if count != 1 {
			err = errors.New(v1 + "遊戲type有誤")
			return
		}
	}

	c = &CardType{
		cardTypeS: CardTypeSlice,
	}

	return c, nil
}

func GhostOne(gcard map[int]int, CardTypeSlice []string) map[int]int {
	var cardType string
	var max []int
	combin := make(map[string][]int)

	for i := 1; i <= 13; i++ {
		var n []int
		for k := range gcard {
			if k > 52 && k <= 54 {
				gcard[k] = i
			}
		}

		n = append(n, i)
		cardType = CompareType(gcard, CardTypeSlice)
		if cardType == "" {
			cardType = CompareCow(gcard)
		}
		combin[cardType] = n
	}

	for _, v := range CardTypeSlice {
		for k, v2 := range combin {
			if v == k {
				max = append(max, v2[0])
			}
		}
	}

	for k := range gcard {
		if k > 52 && k <= 54 {
			gcard[k] = max[0]
		}
	}

	// log.Println(gcard, combin)
	return gcard
}

func GhostTwo(gcard map[int]int, CardTypeSlice []string) map[int]int {
	var cardType string
	var gList, max [][]int
	combin := make(map[string][]int)
	// combin := make(map[string][][]int)

	for i := 1; i <= 13; i++ {
		for j := 1; j <= 13; j++ {
			var g []int
			g = append(g, i)
			g = append(g, j)
			gList = append(gList, g)
		}
	}

	for _, v1 := range gList {
		for k2 := range gcard {
			if k2 == 53 {
				gcard[k2] = v1[0]
			}

			if k2 == 54 {
				gcard[k2] = v1[1]
			}
		}

		cardType = CompareType(gcard, CardTypeSlice)
		if cardType == "" {
			cardType = CompareCow(gcard)
		}

		//全部結果mapping
		// var com [][]int
		// com = append(com, v1)
		// if _, ok := combin[cardType]; ok {
		// 	combin[cardType] = append(combin[cardType], v1)
		// } else {
		// 	combin[cardType] = com
		// }
		combin[cardType] = v1
	}

	// log.Println(combin)
	for _, v := range CardTypeSlice {
		for k, v2 := range combin {
			if v == k {
				max = append(max, v2)
			}
		}
	}

	for k := range gcard {
		if k == 54 {
			gcard[k] = max[0][0]
		}

		if k == 53 {
			gcard[k] = max[0][1]
		}
	}

	// log.Println(gcard)
	return gcard
}

// 轉換牌型
func ChangeCard(origin []int, CardTypeSlice []string) map[int]int {
	var ghost int
	mResult := make(map[int]int)
	for _, v := range origin {
		k := v
		r := v % 13
		if r == 0 {
			r = 13
		}

		if v > 52 && v <= 54 {
			ghost++
			k = v
			r = v
		}

		mResult[k] = r
	}

	if ghost == 1 {
		return GhostOne(mResult, CardTypeSlice)
	} else if ghost == 2 {
		return GhostTwo(mResult, CardTypeSlice)
	}

	return mResult
}

// 牛牛~無牛 牛牛類型的玩法應該不需要白名單
func CompareCow(result map[int]int) (cardType string) {
	var c1, c2, c3, c4, c5, c6, c7, c8, c9, c10 int
	var cFlag bool
	s := make([]int, 0)
	number := make([]int, 0)

	for _, v := range result {
		if v == 11 || v == 12 || v == 13 {
			v = 10
		}

		number = append(number, v)
	}

	if (number[0]+number[1]+number[2])%10 == 0 {
		c1 = (number[3] + number[4]) % 10
	} else {
		c1 = -1
	}
	s = append(s, c1)

	if (number[0]+number[1]+number[3])%10 == 0 {
		c2 = (number[2] + number[4]) % 10
	} else {
		c2 = -1
	}
	s = append(s, c2)

	if (number[0]+number[1]+number[4])%10 == 0 {
		c3 = (number[2] + number[3]) % 10
	} else {
		c3 = -1
	}
	s = append(s, c3)

	if (number[0]+number[2]+number[3])%10 == 0 {
		c4 = (number[1] + number[4]) % 10
	} else {
		c4 = -1
	}
	s = append(s, c4)

	if (number[0]+number[2]+number[4])%10 == 0 {
		c5 = (number[1] + number[3]) % 10
	} else {
		c5 = -1
	}
	s = append(s, c5)

	if (number[0]+number[3]+number[4])%10 == 0 {
		c6 = (number[1] + number[2]) % 10
	} else {
		c6 = -1
	}
	s = append(s, c6)

	if (number[1]+number[2]+number[3])%10 == 0 {
		c7 = (number[0] + number[4]) % 10
	} else {
		c7 = -1
	}
	s = append(s, c7)

	if (number[1]+number[2]+number[4])%10 == 0 {
		c8 = (number[0] + number[3]) % 10
	} else {
		c8 = -1
	}
	s = append(s, c8)

	if (number[1]+number[3]+number[4])%10 == 0 {
		c9 = (number[0] + number[2]) % 10
	} else {
		c9 = -1
	}
	s = append(s, c9)

	if (number[2]+number[3]+number[4])%10 == 0 {
		c10 = (number[0] + number[1]) % 10
	} else {
		c10 = -1
	}

	s = append(s, c10)
	sort.Ints(s)

	// log.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10)
	// log.Println(s)
	if (c1 + c2 + c3 + c4 + c5 + c6 + c7 + c8 + c9 + c10) == -10 {
		return "c0"
	}

	for _, v := range s {
		if v == 0 {
			cFlag = true
			cardType = "cc"
		}
	}

	if !cFlag {
		cardType = "c" + strconv.Itoa(s[9])
	}

	return cardType
}

//五小牛 四炸 五花牛 四花牛
func CompareType(result map[int]int, CardTypeSlice []string) (cardType string) {
	var card, count int
	var iften bool
	CardTypeSwitch := make(map[string]bool)

	for _, v := range CardTypeSlice {
		switch v {
		case "4f":
			CardTypeSwitch[v] = true
		case "5f":
			CardTypeSwitch[v] = true
		case "ff":
			CardTypeSwitch[v] = true
		case "sc":
			CardTypeSwitch[v] = true
		default:
		}
	}

	if CardTypeSwitch["4f"] {
		for _, v := range result {
			if v == 10 {
				iften = true
			}

			if v >= 11 {
				card++
			}

			if card == 4 && iften {
				cardType = "4f"
			}
		}
		card = 0
	}

	if CardTypeSwitch["5f"] {
		for _, v := range result {
			if v >= 11 {
				card++
			}

			if card == 5 {
				cardType = "5f"
			}
		}
		card = 0
	}

	if CardTypeSwitch["ff"] {
		for i := 1; i <= 13; i++ {
			for _, v := range result {
				if v == i {
					card++
				}
			}

			if card == 4 {
				cardType = "ff"
			}
			card = 0
		}
	}

	if CardTypeSwitch["sc"] {
		for _, v := range result {
			if v < 5 {
				card++
				count += v
			}

			if card == 5 && count <= 10 {
				cardType = "sc"
			}
		}
	}

	return cardType
}

func GetMaxCard(result map[int]map[int]int) (maxCard []int) {
	maxColor := make(map[int]int)
	maxNumber := make(map[int]int)

	for menber, card := range result {
		var color, number []int
		for k, v := range card {
			color = append(color, k)
			number = append(number, v)
			sort.Sort(sort.Reverse(sort.IntSlice(color)))
			sort.Sort(sort.Reverse(sort.IntSlice(number)))
		}
		//花色數字排序用
		for _, v := range color {
			nn := v % 13

			if nn == 0 {
				nn = 13
			}

			if nn == number[0] {
				maxColor[v] = menber
				break
			}
		}

		// maxColor[color[0]] = menber
		maxNumber[menber] = number[0]
	}

	var nMax, cMax []int
	//排序數字大玩家 無順序
	for i := 13; i >= 1; i-- {
		for menber, number := range maxNumber {
			if i == number {
				nMax = append(nMax, menber)
			}
		}
	}

	//取重複牌型玩家 無順序
	var duplicate []int
	for k, v := range maxNumber {
		for k2, v2 := range maxNumber {
			if v == v2 && k != k2 {
				duplicate = append(duplicate, k2)
			}
		}
	}

	//golang map 用原生range遍歷不能保證順序輸出
	//降冪排列花色數字
	var keys []int
	for k := range maxColor {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	//若沒有重複牌型 不需找牌型merge 數字大的贏
	if len(duplicate) == 0 {
		cMax = nMax
	} else {
		for _, k := range keys {
			for _, v := range duplicate {
				if maxColor[k] == v {
					cMax = append(cMax, v)
				}
			}
		}
	}

	maxCard = MergeSortSlice(nMax, cMax)

	// log.Println(result)
	// log.Println(cMax, nMax, maxCard, maxColor, maxNumber)
	return maxCard
}

func MergeSortSlice(first []int, second []int) (result []int) {
	var s []int
	f := make(map[int]int)
	for k, v := range first {
		f[k] = v
	}

	for _, v1 := range second {
		for k2, v2 := range f {
			if v1 == v2 {
				s = append(s, k2)
			}
		}
	}

	sort.Ints(s)
	for k1, v1 := range s {
		f[v1] = second[k1]
	}
	//golang map 用原生range遍歷不能保證順序輸出
	var keys []int
	for k := range f {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, v := range keys {
		result = append(result, f[v])
	}

	// log.Println(f, s, result)
	return result
}

func (c *CardType) CompareResult(originResults [][]int) map[int][]int {
	var cardType string
	var level, count1, count2, count3, count4, count5, count6, count7, count8, count9, count10, count11, count12, count13, count14, count15 int
	mSameResult := make(map[int]int)
	mSameCount := make(map[int]int)
	mSameLevelResult := make(map[int][]int)

	for menber, originResult := range originResults {
		result := ChangeCard(originResult, c.cardTypeS)
		cardType = CompareType(result, c.cardTypeS)
		if cardType == "" {
			cardType = CompareCow(result)
		}

		switch cardType {
		case "sc":
			count1++
			level = 15
			mSameCount[level] = count1
		case "ff":
			count2++
			level = 14
			mSameCount[level] = count2
		case "5f":
			count3++
			level = 13
			mSameCount[level] = count3
		case "4f":
			count4++
			level = 12
			mSameCount[level] = count4
		case "cc":
			count5++
			level = 11
			mSameCount[level] = count5
		case "c9":
			count6++
			level = 10
			mSameCount[level] = count6
		case "c8":
			count7++
			level = 9
			mSameCount[level] = count7
		case "c7":
			count8++
			level = 8
			mSameCount[level] = count8
		case "c6":
			count9++
			level = 7
			mSameCount[level] = count9
		case "c5":
			count10++
			level = 6
			mSameCount[level] = count10
		case "c4":
			count11++
			level = 5
			mSameCount[level] = count11
		case "c3":
			count12++
			level = 4
			mSameCount[level] = count12
		case "c2":
			count13++
			level = 3
			mSameCount[level] = count13
		case "c1":
			count14++
			level = 2
			mSameCount[level] = count14
		case "c0":
			count15++
			level = 1
			mSameCount[level] = count15
		default:
			level = 0
		}

		mSameResult[menber] = level
	}
	// log.Println(mSameResult)
	// log.Println(mSameCount)
	for l := 1; l <= 15; l++ {
		var menber []int
		if mSameCount[l] >= 1 {
			for j := 0; j < len(mSameResult); j++ {
				if mSameResult[j] == l {
					menber = append(menber, j)
					mSameLevelResult[l] = menber
				}
			}
		}
	}

	topRank := make(map[int]map[int]map[int]int)
	for level, same := range mSameLevelResult {
		mc := make(map[int]map[int]int)
		for _, menber := range same {
			mc[menber] = ChangeCard(originResults[menber], c.cardTypeS)
		}
		topRank[level] = mc
		//相同牌型在判斷
		if len(mc) > 1 {
			rSort := GetMaxCard(mc)
			mSameLevelResult[level] = rSort
		}
	}

	return mSameLevelResult
}

// func main() {
// 	originResults := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}, {26, 27, 28, 29, 30}, {31, 32, 33, 34, 35}, {36, 37, 38, 39, 40}, {41, 42, 43, 44, 45}, {46, 47, 48, 53, 54}}
// 	// originResults := [][]int{{11, 24, 37, 52, 23}, {12, 25, 38, 50, 36}, {13, 26, 39, 51, 49}, {1, 14, 27, 3, 4}, {10, 9, 8, 7, 5}, {49, 22, 21, 20, 31}}
// 	// originResults := [][]int{{1, 2, 7, 54, 27}, {14, 15, 16, 29, 53}}
// 	//card type 給重複一樣type可以接受
// 	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
// 	n, err := New(cts)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	result := n.CompareResult(originResults)
// 	log.Println(result)
// }

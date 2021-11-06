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

type UserCardType struct {
	Member int
	Color  int
	Number int
}

func GetMaxCard(result map[int]map[int]int) (maxCard []int) {
	for i := len(result); i > 0; i-- {
		var member int
		member, result = cycle(result)
		maxCard = append(maxCard, member)
	}

	return maxCard
}

func cycle(result map[int]map[int]int) (int, map[int]map[int]int) {
	var userCardType UserCardType
	for member, card := range result {
		for k, v := range card {
			// log.Println(k, v, "--------------")
			if userCardType.Number < v {
				userCardType.Member = member
				userCardType.Color = k
				userCardType.Number = v
			}

			if userCardType.Number == v {
				if userCardType.Color < k {
					userCardType.Member = member
					userCardType.Color = k
				}
			}
		}
	}
	// log.Println(userCardType)
	delete(result, userCardType.Member)

	return userCardType.Member, result
}

func (c *CardType) CompareResult(originResults [][]int) map[int][]int {
	var cardType string
	var level int
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
			level = 15
			mSameCount[level]++
		case "ff":
			level = 14
			mSameCount[level]++
		case "5f":
			level = 13
			mSameCount[level]++
		case "4f":
			level = 12
			mSameCount[level]++
		case "cc":
			level = 11
			mSameCount[level]++
		case "c9":
			level = 10
			mSameCount[level]++
		case "c8":
			level = 9
			mSameCount[level]++
		case "c7":
			level = 8
			mSameCount[level]++
		case "c6":
			level = 7
			mSameCount[level]++
		case "c5":
			level = 6
			mSameCount[level]++
		case "c4":
			level = 5
			mSameCount[level]++
		case "c3":
			level = 4
			mSameCount[level]++
		case "c2":
			level = 3
			mSameCount[level]++
		case "c1":
			level = 2
			mSameCount[level]++
		case "c0":
			level = 1
			mSameCount[level]++
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

	// topRank := make(map[int]map[int]map[int]int)
	for level, same := range mSameLevelResult {
		mc := make(map[int]map[int]int)
		for _, menber := range same {
			mc[menber] = ChangeCard(originResults[menber], c.cardTypeS)
		}
		// topRank[level] = mc
		//相同牌型在判斷
		if len(mc) > 1 {
			rSort := GetMaxCard(mc)
			mSameLevelResult[level] = rSort
		}
	}

	return mSameLevelResult
}

// func main() {
// 	// originResults := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}, {26, 27, 28, 29, 30}, {31, 32, 33, 34, 35}, {36, 37, 38, 39, 40}, {41, 42, 43, 44, 45}, {46, 47, 48, 53, 54}}
// 	// originResults := [][]int{{11, 24, 37, 52, 23}, {12, 25, 38, 50, 36}, {13, 26, 39, 51, 49}, {1, 14, 27, 3, 4}, {10, 9, 8, 7, 5}, {49, 22, 21, 20, 31}}
// 	// originResults := [][]int{{16, 28, 12, 22, 49}, {48, 32, 10, 45, 25}}
// 	originResults := [][]int{{2, 3, 9, 10, 13}, {41, 42, 48, 49, 51}, {15, 16, 22, 23, 26}, {28, 29, 35, 36, 38}}
// 	//card type 給重複一樣type可以接受
// 	cts := []string{"sc", "ff", "5f", "4f", "cc", "c9", "c8", "c7", "c6", "c5", "c4", "c3", "c2", "c1", "c0"}
// 	n, err := New(cts)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	for i := 0; i < 500; i++ {
// 		result := n.CompareResult(originResults)
// 		log.Println(result)
// 	}
// }

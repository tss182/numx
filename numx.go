package numx

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

const baseStr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateID() string {
	timeNow := time.Now().UnixNano()
	return convertIntToNumx(timeNow)
}

func GetTimeFromID(str string) time.Time {
	i := convertNumxToInt(str)
	str2 := fmt.Sprintf("%d", i)
	fmt.Println(str2)
	i, _ = strconv.ParseInt(str2, 10, 64)
	return time.Unix(0, i)
}

func convertIntToNumx(n int64) string {
	var result string
	baseNumber := int64(len(baseStr))
	for n >= baseNumber {
		remain := n % baseNumber
		n /= baseNumber
		result = baseStr[remain:remain+1] + result
	}
	var str = baseStr[n : n+1]
	result = str + result
	return result
}

func convertNumxToInt(str string) int64 {
	baseSplit := strings.Split(baseStr, "")
	strSplit := strings.Split(str, "")
	baseNumber := len(baseStr)
	var result int64
	for i, v := range strSplit {
		power := len(strSplit) - i - 1
		num := int64(sort.SearchStrings(baseSplit, v)) * int64(math.Pow(float64(baseNumber), float64(power)))
		result += num
	}
	return result
}

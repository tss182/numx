package numx

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GenerateID() string {
	timeNow := time.Now().UnixNano()
	return encode(timeNow, typeStr)
}

func GenerateNumber() string {
	timeNow := time.Now().UnixNano()
	return encode(timeNow, typeNumber)
}

func GetTimeFromID(str string) time.Time {
	i := decode(str, typeStr)
	str2 := fmt.Sprintf("%d", i)
	fmt.Println(str2)
	i, _ = strconv.ParseInt(str2, 10, 64)
	return time.Unix(0, i)
}

func encode(n int64, typeData string) string {
	baseDefault := getBase(typeData)
	var result string
	baseData := int64(len(baseDefault))
	for n >= baseData {
		remain := n % baseData
		n /= baseData
		result = baseDefault[remain:remain+1] + result
	}
	var str = baseDefault[n : n+1]
	result = str + result
	return result
}

func decode(plaintext, typeData string) int64 {
	baseDefault := getBase(typeData)
	baseSplit := strings.Split(baseDefault, "")
	strSplit := strings.Split(plaintext, "")
	baseData := len(baseDefault)
	var result int64
	for i, v := range strSplit {
		power := len(strSplit) - i - 1
		num := int64(sort.SearchStrings(baseSplit, v)) * int64(math.Pow(float64(baseData), float64(power)))
		result += num
	}
	return result
}

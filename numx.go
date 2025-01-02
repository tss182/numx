package numx

import (
	"crypto/rand"
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
	i, _ = strconv.ParseInt(str2, 10, 64)
	return time.Unix(0, i)
}

func randPrefix(baseString string, number int) string {
	var randomByte = make([]byte, number)
	_, err := rand.Read(randomByte)
	if err != nil {
		panic(err)
	}
	var sb strings.Builder
	for _, b := range randomByte {
		sb.WriteByte(baseString[int(b)%len(baseString)])
	}
	return sb.String()
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
	randomStr := randPrefix(baseDefault,8)
	result = randomStr + result
	var endResult string
	lenSplit := 5
	for i:=0; i<len(result); i+=lenSplit {
		end := i + lenSplit
		if end > len(result) {
			end = len(result)
		}
		endResult += result[i:end]+"-"
	}
	endResult = strings.TrimRight(endResult, "-")
	return endResult
}

func GetOtp(length int) string {
	var randomByte = make([]byte, length)
	_, err := rand.Read(randomByte)
	if err != nil {
		panic(err)
	}
	var sb strings.Builder
	for _, b := range randomByte {
		newChar := int(b)%10
		sb.WriteString(strconv.Itoa(newChar))
	}
	return sb.String()
}

func decode(plaintext, typeData string) int64 {
	baseDefault := getBase(typeData)
	baseSplit := strings.Split(baseDefault, "")
	plaintext = strings.Replace(plaintext,"-","",-1)
	plaintext = plaintext[8:]
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

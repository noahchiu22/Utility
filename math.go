package util

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// unit代表小數點後取到第幾位
func Round(num, unit float64) float64 {
	return math.Round(num*math.Pow(10, unit)) / math.Pow(10, unit)
}

// 找平均值
func Average(data []float64) (avg float64) {
	sum := 0.0
	for i := range data {
		sum += data[i]
	}
	avg = sum / float64(len(data))
	return
}

// 找標準差，會順便一起算平均值
func Standard(data []float64) (SD, avg float64) {
	avg = Average(data)
	sigma := 0.0
	for i := range data {
		sigma += math.Pow((data[i] - avg), 2)
	}
	Variance := sigma / float64(len(data))
	SD = math.Sqrt(Variance)
	return
}

// 找極值
func FindExtremum(data []float64) (max, min float64) {
	if len(data) == 0 {
		return
	}

	max = data[0]
	min = data[0]

	for _, item := range data {
		if item > max {
			max = item
		}
		if item < min {
			min = item
		}
	}

	return
}

// 字串計算(字串+數字、字串+字串)(減法有問題)
// 輸出會依照stringNum的位元數來輸出，多出來的就不輸出
// 字串計算(字串+數字、字串+字串)
// 輸出會依照stringNum的位元數來輸出，多出來的就不輸出
func StringCalculate(stringNum string, inputNum any) (outputNum string) {
	fmt.Printf("stringNum: %v, inputNum: %v\n", stringNum, inputNum)

	num := 0

	// 用於字串變數字
	switch temp := inputNum.(type) {
	case string:
		num, _ = strconv.Atoi(temp)
	case int:
		num = temp
	}

	carry := num
	for i := len(stringNum) - 1; i >= 0; i-- {
		// 數字加進位，加完後進位要清空(第一次的carry就是要加的數字)
		digitNum, _ := strconv.Atoi(string(stringNum[i]))

		if carry < 0 {
			digitNum += (carry % 10)
			carry /= 10
		} else {
			digitNum += carry
			carry = 0
		}
		// 超過9要進位
		if digitNum > 9 {
			carry = (digitNum) / 10
			digitNum = (digitNum) % 10
		}
		// 小於0要退位(-的num)
		if digitNum < 0 {
			digitNum = 10 + digitNum
			carry += -1
		}
		// 由右到左拼起來
		outputNum = fmt.Sprint(digitNum) + outputNum
	}

	return
}

// 自動計算系統維護之ID(加法或減法都可以)
// 沒有輸入prefix會以serialStr前面的prefix為主
// 沒有輸入initNumStr會自動找編號位元大小，有的話會以initNumStr位元大小為主
func AddSerial(prefix, initNumStr, serialStr string, addend int) (output string) {
	if serialStr == "" {
		output = initNumStr
		if prefix != "" {
			output = prefix + output
		}
		return
	}
	// 數字檢查regexp
	numRegexp := regexp.MustCompile("[0-9]")
	carry := 0
	for i := len(serialStr) - 1; i >= 0; i-- {
		// 如果不是數字就把前面的文字加上然後回傳
		if !numRegexp.MatchString(string(serialStr[i])) {
			if prefix != "" {
				output = prefix + output
				return
			}
			output = serialStr[:i+1] + output
			return
		}
		// A的該位數加上B的該位數再加進位
		digitNum, _ := strconv.Atoi(string(serialStr[i]))

		digitNum = digitNum + (addend % 10) + carry

		// 超過9要進位
		if digitNum > 9 {
			carry = (digitNum) / 10
			digitNum = (digitNum) % 10
		}
		// 小於0要退位
		if digitNum < 0 {
			carry = -1
			digitNum = 10 + digitNum
		}

		// 每次只拿個位數計算
		addend /= 10
		// 由右到左拼起來
		output = fmt.Sprint(digitNum) + output

		// 如果輸出長度已經等於initNumStr的長度就把前面的文字加上然後回傳
		if len(serialStr)-i == len(initNumStr) {
			if prefix != "" {
				output = prefix + output
			}
			return
		}
	}
	return
}

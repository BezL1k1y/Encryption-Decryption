package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var (
	Parameter      int64
	counterDeShifr int64 = 1
	Answer         int64 = 1
)
var (
	slicer int64 = 0
	N      int64
	e      int64
	fiN    int64
)
var (
	AnswerDESHIFER   string
	str, xS, command string
	AnswerCode       []int64
)

// Main 0//
func main() {
	for {
		AnswerDESHIFER = ""
		AnswerCode = []int64{}
		fmt.Print(" - (1) Find d\n" +
			" - (2) Decryption\n" +
			" - (3) Encryption\n" +
			" - (4) Find the divisors\n" +
			" - (5) Exit\n")
		fmt.Print("Введите комманду: ")
		fmt.Scan(&command)

		if command == "1" { //Find d
			fmt.Print("Введите ключ\n" +
				"Введите число N:  ")
			fmt.Scan(&N)
			fmt.Print("Введите число e: ")
			fmt.Scan(&e)
			delitel := RhoPollard(N)
			fiN = (delitel - 1) * (N/delitel - 1)
			eNew := big.NewInt(e)
			finNew := big.NewInt(fiN)
			d := new(big.Int).ModInverse(eNew, finNew)
			fmt.Print(d, "\n\n")

		} else if command == "2" { //Decryption
			fmt.Print("Введите ключ\n" +
				"Введите число N:  ")
			fmt.Scan(&N)
			fmt.Print("Введите число d: ")
			fmt.Scan(&e)
			RhoPollard(N)
			u := big.NewInt(e)
			str = Reverse(fmt.Sprintf("%b", u))
			Parameter = -1
			Enter()

		} else if command == "3" { //Encryption
			fmt.Print("Введите ключ\n" +
				"Введите число N:  ")
			fmt.Scan(&N)
			fmt.Print("Введите число e: ")
			fmt.Scan(&e)
			RhoPollard(N)
			u := big.NewInt(e)
			str = Reverse(fmt.Sprintf("%b", u))
			Parameter = 1
			Enter()

		} else if command == "4" { //Find the divisors
			fmt.Print("Введите число для факторизации:  ")
			fmt.Scan(&N)
			ans := RhoPollard(N)
			fmt.Print("1. - ", ans, "\n")
			fmt.Print("2. - ", N/ans, "\n\n")
		} else if command == "5" { //Exit
			os.Exit(1)
		}
	}
}

// RhoPollard 1//
func RhoPollard(N int64) int64 {
	var (
		x           = N / 2
		y     int64 = 1
		i     int64 = 0
		stage int64 = 2
	)
	for NOD(N, int64(math.Abs(float64(x-y)))) == 1 {
		if i == stage {
			y = x
			stage = stage * 2
		}
		x = int64(math.Mod(float64(x*x+1), float64(N)))
		i = i + 1
	}
	return NOD(N, int64(math.Abs(float64(x-y))))
}

// NOD 2//
func NOD(a, b int64) int64 {
	if b == 0 {
		return a
	}
	nod := a
	a = b
	b = nod % a
	return NOD(a, b)
}

// DeShifr 3//
func DeShifr(x int64, e int64, N int64) int64 {
	if counterDeShifr < e {
		if x > N {
			x = int64(math.Mod(float64(x), float64(N)))
		}
		if str[slicer:slicer+1] == "1" {
			Answer *= x
		}
		x = x * x
		counterDeShifr *= 2
		slicer += 1
		return DeShifr(x, e, N)
	}
	return int64(math.Mod(float64(x), float64(N)))
}

// Shifr 4//
func Shifr(x int64, e int64, N int64) int64 {
	if counterDeShifr < e {
		if x > N {
			x = int64(math.Mod(float64(x), float64(N)))
		}
		if str[slicer:slicer+1] == "1" {
			Answer *= x
		}
		x = x * x
		counterDeShifr *= 2
		slicer += 1
		return Shifr(x, e, N)
	}
	return int64(math.Mod(float64(x), float64(N)))
}

// Reverse 5//
func Reverse(t string) string {
	runes := []rune(t)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Enter 6//
func Enter() {
	if Parameter == 1 {
		print("Введите, что нужно Кодировать: ")
		fmt.Scan(&xS)
		byteArray := []byte(xS)
		for o := 0; o < len(byteArray); o++ {
			Shifr(int64(byteArray[o]), e, N)
			asciiNum := int(math.Mod(float64(Answer), float64(N)))
			AnswerCode = append(AnswerCode, int64(asciiNum))
			character := strconv.Itoa(asciiNum)
			AnswerDESHIFER += character + "-"
			counterDeShifr = 1
			slicer = 0
			Answer = 1
		}
		AnswerDESHIFER = strings.TrimRight(AnswerDESHIFER, "-")
		fmt.Print("Кодированное послание на ASCII: ", AnswerDESHIFER, "\n")

	} else if Parameter == -1 {
		print("Введите, что нужно Декодировать: ")
		fmt.Scan(&xS)
		byteArray := strings.Split(xS, "-")
		for o := 0; o < len(byteArray); o++ {
			a := byteArray[o]
			c, _ := strconv.Atoi(a)
			DeShifr(int64(c), e, N)
			asciiNum := int64(math.Mod(float64(Answer), float64(N)))
			character := string(rune(asciiNum))
			AnswerCode = append(AnswerCode, asciiNum)
			AnswerDESHIFER += character
			counterDeShifr = 1
			slicer = 0
			Answer = 1
		}
		fmt.Print("Послание в переводе с ASCII: ", AnswerDESHIFER, "\n")
		fmt.Print("Код после дешифровки: : ", AnswerCode, "\n\n")

	} else {
		os.Exit(1)
	}
}

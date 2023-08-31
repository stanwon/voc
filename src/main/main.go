package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	voc()
}

func voc() {
	path := chooseVoc()
	zh, en := readRaw1(path)
	switch readAction() {
	case "1":
		practice(zh, en)
	case "2":
		zhToEn(zh, en)
	/* case "3":
	enToZh(zh, en) */
	default:
		fmt.Println("unknown action.")
	}
}

func chooseVoc() (path string) {
	for {
		fmt.Println("1: 初中-乱序")
		fmt.Println("2: 高中-乱序")
		fmt.Println("3: 四级-乱序")
		fmt.Println("4: 六级-乱序")
		fmt.Println("5: 考研-乱序")
		fmt.Println("6: 托福-乱序")
		fmt.Println("7: SAT-乱序")
		sIn := bufio.NewScanner(os.Stdin)
		sIn.Scan()
		choose := sIn.Text()
		switch choose {
		case "1":
			path = "english-vocabulary/1.初中-乱序.txt"
			return
		case "2":
			path = "english-vocabulary/2.高中-乱序.txt"
			return
		case "3":
			path = "english-vocabulary/3.四级-乱序.txt"
			return
		case "4":
			path = "english-vocabulary/4.六级-乱序.txt"
			return
		case "5":
			path = "english-vocabulary/5.考研-乱序.txt"
			return
		case "6":
			path = "english-vocabulary/6.托福-乱序.txt"
			return
		case "7":
			path = "english-vocabulary/7.SAT-乱序.txt"
			return
		}
	}
}

func readAction() (action string) {
	for {
		fmt.Println("1: practicing")
		fmt.Println("2: zh to en")
		// fmt.Println("3: en to zh")
		sIn := bufio.NewScanner(os.Stdin)
		sIn.Scan()
		action = sIn.Text()
		if "0" == action || "1" == action || "2" == action {
			break
		}
	}
	return
}

func readRaw1(path string) (zh, en []string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scaner := bufio.NewScanner(f)

  count := 1
	for scaner.Scan() {
		if err = scaner.Err(); err != nil {
			panic(err)
		}
		raw := scaner.Text()
		before, after, ok := strings.Cut(raw, "\t")
		if ok {
			en = append(en, before)
			zh = append(zh, strings.Trim(after, " "))
		}
    count++
	}
  fmt.Println(count)
	return
}

func readRaw() (zh, en []string) {
	enPath := "./en.txt"
	zhPath := "./zh.txt"
	fEn, err := os.Open(enPath)
	if err != nil {
		panic(err)
	}
	fZh, err := os.Open(zhPath)
	if err != nil {
		panic(err)
	}
	defer fEn.Close()
	defer fZh.Close()
	sEn := bufio.NewScanner(fEn)
	sZh := bufio.NewScanner(fZh)

	for sEn.Scan() {
		if err = sEn.Err(); err != nil {
			panic(err)
		}
		en = append(en, sEn.Text())
	}

	for sZh.Scan() {
		if err = sZh.Err(); err != nil {
			panic(err)
		}
		zh = append(zh, sZh.Text())
	}
	return
}

var indexs = make(map[int]int)

func checkIdx(i int) bool {
	if _, ok := indexs[i]; ok {
		return true
	}
	indexs[i] = i
	return false
}

func practice(zh, en []string) {
	for {
		idx := rand.Intn(len(zh))
		if checkIdx(idx) {
			continue
		}
		fmt.Println(en[idx])
		fmt.Println(zh[idx])
		sIn := bufio.NewScanner(os.Stdin)
		sIn.Scan()
		sIn.Text()
	}
}

func zhToEn(zh, en []string) {
	for {
		idx := rand.Intn(len(zh))
		fmt.Println(zh[idx])
		sIn := bufio.NewScanner(os.Stdin)
		for {
			sIn.Scan()
			result := sIn.Text()
			if result == "j" {
				fmt.Println(en[idx])
				break
			}
			if result == en[idx] {
				break
			}
		}
	}
}

func enToZh(zh, en []string) {
	for {
		idx := rand.Intn(len(en))
		fmt.Println(en[idx])
		sIn := bufio.NewScanner(os.Stdin)
		for {
			sIn.Scan()
			result := sIn.Text()
			if result == "j" {
				fmt.Println(zh[idx])
				break
			}
			if result == zh[idx] {
				break
			}
		}
	}
}

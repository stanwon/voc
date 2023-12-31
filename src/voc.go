package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

type voc struct {
	zh   [][]string
	en   [][]string
	nums int
}

func newVoc() *voc {
	return &voc{}
}

func (voc *voc) readVoc(nav *nav) {
	for _, dict := range nav.pages[0].lines {
		en, zh := readRaw(filepath.Join("english-vocabulary", dict))
		voc.en = append(voc.en, en)
		voc.zh = append(voc.zh, zh)
		voc.nums++
	}
	nav.pageNum++
	nav.pages = append(nav.pages, &page{
		init: true,
	})
}

func vocfunc() {
	for {
		path := chooseVoc()
		if path == "" {
			return
		}
		zh, en := readRaw(path)
		switch readAction() {
		case "1":
			practice(zh, en)
		case "2":
			zhToEn(zh, en)
		case "q":
			// continue
		default:
			fmt.Println("unknown action.")
		}
	}
}

func chooseVoc() (path string) {
	for {
		fmt.Println("-----choose dictionary-----")
		fmt.Println("1: 初中-乱序")
		fmt.Println("2: 高中-乱序")
		fmt.Println("3: 四级-乱序")
		fmt.Println("4: 六级-乱序")
		fmt.Println("5: 考研-乱序")
		fmt.Println("6: 托福-乱序")
		fmt.Println("7: SAT-乱序")
		fmt.Println("q: quit")
		fmt.Println("---------------------------")
		fmt.Print("choose: ")
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
		case "q":
			path = ""
			return
		}
	}
}

func readAction() (action string) {
	for {
		fmt.Println("---------------------------")
		fmt.Println("1: show vocabulary")
		fmt.Println("2: zh to en")
		fmt.Println("q: back to choose dictionary")
		fmt.Println("---------------------------")
		fmt.Print("choose: ")
		// fmt.Println("3: en to zh")
		sIn := bufio.NewScanner(os.Stdin)
		sIn.Scan()
		action = sIn.Text()
		if "q" == action || "1" == action || "2" == action {
			break
		}
	}
	return
}

func readRaw(path string) (zh, en []string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scaner := bufio.NewScanner(f)

	count := 1
	for scaner.Scan() {
		if err = scaner.Err(); err != nil {
			log.Fatal(err)
		}
		raw := scaner.Text()
		before, after, ok := strings.Cut(raw, "\t")
		if ok {
			en = append(en, before)
			zh = append(zh, strings.Trim(after, " "))
		}
		count++
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
		sIn := bufio.NewScanner(os.Stdin)
		count := 1
		fmt.Println("-------------------------------------")
		fmt.Println(zh[idx])
		for {
			fmt.Print("answer: ")
			sIn.Scan()
			result := sIn.Text()
			if result == "j" {
				fmt.Println(en[idx])
				break
			}

			if result == en[idx] {
				break
			} else if count <= len(en[idx]) {
				fmt.Print("tips: ")
				fmt.Println(en[idx][:count])
				count++
			} else if count > len(en[idx]) {
				fmt.Println("!!!笨蛋!!!")
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

			fmt.Println(len(en[idx]))
			if result == zh[idx] {
				break
			}
		}
	}
}

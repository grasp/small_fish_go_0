package main

import (
	"fmt"
	//"os"
	"io/ioutil"
	"strings"
)

func scan_dir() {

}

//读取文件到一个map类型中,map["date"]=[signal_value]
func read_single_signal_file(signalName string, percent_num string, buy_by_sell_by string, symbol string) {
	symbol_path := "e:\\single_signal\\" + signalName + "\\" + symbol + ".txt" //could be path split
	dat, _ := ioutil.ReadFile(symbol_path)
	var oneLine []string
	var s = string(dat)
	var lineArray = strings.Split(s, "\r\n")
	var resultMap = make(map[string]string, len(lineArray))
	for _, value := range lineArray {
		oneLine = strings.Split(value, "#")
		if len(oneLine) == 2 {
			resultMap[oneLine[0]] = oneLine[1]
		}
	}
	//fmt.Println(resultMap)
	//fmt.Println(strings.Split("this is a test \n and another", "\n"), symbol_path)

	win_lost_path := "E:\\win_lost\\" + percent_num + "\\" + buy_by_sell_by + "\\" + symbol + ".txt"

	//fmt.Println(win_lost_path)
	dat2, _ := ioutil.ReadFile(win_lost_path)
	s = string(dat2)
	//fmt.Println(s)
	lineArray = strings.Split(s, "\r\n")
	var resultMap2 = make(map[string]string, len(lineArray))

	for _, value := range lineArray {
		oneLine = strings.Split(value, "#")
		if len(oneLine) == 2 {
			resultMap2[oneLine[0]] = oneLine[1]
		}
	}

	type winlost struct {
		win  int
		lost int
	}

	var a, b winlost
	//fmt.Println(resultMap2)
	var statsticMap = make(map[string]winlost, 2)
	//var winStatstic [2]int
	//	var lostStatstic [2]int
	var total_count = 0

	for key, value := range resultMap {
		total_count += 1
		//fmt.Println(key, value)
		if resultMap2[key] == "true" {
			a = statsticMap[value]
			a.win += 1
			statsticMap[value] = a
		} else {
			b = statsticMap[value]
			b.lost += 1
			statsticMap[value] = b
		}
	}

	fmt.Println(statsticMap)

}

//读取文件到一个map类型中,map["date"]=[win_lost]
func read_win_lost_percent(win_percent, duration_days int) {

}

func write_to_single_statistic(win_percent, duration_days int, symbol string) {

}

func main() {
	signalName := "close_equal_low"
	symbol := "000004.sz"
	percent_num := "percent_1_num_1_days"
	buy_by_sell_by := "buy_by_close_sell_by_close"
	//read_single_signal_file(symbol)
	read_single_signal_file(signalName, percent_num, buy_by_sell_by, symbol)
	fmt.Println("hi")
}

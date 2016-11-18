package main

import (
	"fmt"
	"github.com/funkygao/gocli"
	"os"
	"strconv"
)

const (
	theYear = 2015
	theDate = 5
)

func main() {
	ui := cli.BasicUi{
		Writer:      os.Stdout,
		Reader:      os.Stdin,
		ErrorWriter: os.Stderr,
	}
	ui.Output(fmt.Sprintf("%d年5月1日是星期%d", theYear, theDate))
	year, _ := ui.Ask("你要计算哪一年的5月1日是星期几？请输入年份：")
	ui.Output(fmt.Sprintf("好的，宝贝，你输入了 %s", year))
	ui.Output(fmt.Sprintf("现在，use your brains to calculate what is the date of %s-05-01", year))
	guess, _ := ui.Ask(fmt.Sprintf("你计算的 %s-05-01 是星期几？输入数字：", year))
	if strconv.Itoa(dateOf51(&ui, year)) != guess {
		ui.Error("继续努力啊，就差一点点你就做对了")
	} else {
		ui.Info("bingo! 买糖去！")
	}
}

func dateOf51(ui cli.Ui, year string) (answer int) {
	yearN, _ := strconv.Atoi(year)
	switch {
	case yearN == theYear:
		return theDate

	case yearN > theYear:
		days := 0
		yearDays := 0
		for y := theYear + 1; y <= yearN; y++ {
			if runYear(y) {
				yearDays = 366
			} else {
				yearDays = 365
			}

			days += yearDays
			ui.Output(fmt.Sprintf("%d年 增加了 %d天，累计增加了 %d天", y, yearDays, days))
		}

		mod := days % 7
		answer = (mod + theDate) % 7
		ui.Output(fmt.Sprintf("%d天，是 %d 个周期，余数是 %d", days, days/7, mod))
		ui.Output(fmt.Sprintf("所以答案是：%s-05-01 星期%d", year, answer))

	case yearN < theYear:
	}

	return
}

func runYear(year int) bool {
	return year%4 == 0
}

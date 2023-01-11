package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/now"
)

func biweeklyRange(t time.Time) (biweekKey string, start time.Time, end time.Time) {
	year, week := t.ISOWeek()
	if week%2 == 1 {
	}
	fmt.Println(year)
	fmt.Println(week)

	return
}

func weeklyRange(t time.Time) (weekKey int, start time.Time, end time.Time) {
	_, weekKey = t.ISOWeek()
	return weekKey, weekStart(t), weekEnd(t)
}

func weekStart(t time.Time) time.Time {
	offset := (int(time.Monday) - int(t.Weekday()) - 7) % 7
	result := t.Add(time.Duration(offset*24) * time.Hour)
	return result
}

func weekEnd(t time.Time) time.Time {
	offset := (7 - int(t.Weekday())) % 7
	result := t.Add(time.Duration(offset*24) * time.Hour)
	return result
}

func main() {
	var err error
	erri := fmt.Errorf("error: %s", err.Error())
	fmt.Println(erri)
	log.Fatal("done")

	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

	fmt.Println("chanel length: ", len(ch))
	for {
		if len(ch) == 0 {
			break
		}
		fmt.Println("value: ", <-ch)
	}

	log.Fatal("done")

	//for _, item := range apps {
	//	if _, ok := r[item.ApplicationId]; ok {
	//		r[item.ApplicationId] = append(r[item.ApplicationId], item.Verb)
	//	} else {
	//		r[item.ApplicationId] = []int64{item.Verb}
	//	}
	//}

	timeCfg = &now.Config{
		WeekStartDay: time.Monday,
		TimeLocation: getLocation(),
		TimeFormats:  []string{"2006-01-02 15:04:05"},
	}

	t := time.Now()
	fmt.Println(WeeklyRange(t))
	fmt.Println(PreWeeklyRange(t))
	fmt.Println(BiweeklyRange(t))
	fmt.Println(PreBiweeklyRange(t))
	fmt.Println(MonthlyRange(t))
	fmt.Println(PreMonthlyRange(t))

	//s := weekStartDate(t)
	//e := weekEndDate(t)
	//fmt.Println(s.Date())
	//fmt.Println(e.Date())

	//y, w := t.ISOWeek()
	//m := t.Month()
	//d := t.Day()
	//wd := t.Weekday()
	//fmt.Println(y)
	//fmt.Println(m)
	//fmt.Println(d)
	//fmt.Println(w)
	//fmt.Println(wd)
	//fmt.Println(time.Sunday)
	//_ = copy.Copy("demo1.go", "demoo.go")
	//var x = &[]string{}
	//for _, ite := range *x {
	//	fmt.Println(ite)
	//}
	//
	//logger, _ := zap.NewProduction()
	//defer logger.Sync()
	//sugar := logger.Sugar()
	//sugar.With(
	//	"hello", "world",
	//	getCallerInfo(),
	//)
	//url := "this is a url"
	//sugar.Infow("failed to fetch URL",
	//	"url", url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)
	//
	//logger, _ = zap.NewDevelopment()
	//logger.Warn("this is develop log")

	//
	//curPath, _ := os.Getwd()
	//fmt.Println(curPath)
	//
	//dir, _ := os.MkdirTemp("/tmp", "onepunch")
	//fmt.Println(dir)
	//
	//opt := copy.Options{
	//	AddPermission: 0611,
	//	OnDirExists: func(_, _ string) copy.DirExistsAction {
	//		return copy.Replace
	//	},
	//}
	//_ = copy.Copy("nil/", "nil2", opt)
}

//func getCallerInfo() (callerFields []zap.Field) {
//	pc, file, line, ok := runtime.Caller(2)
//	if !ok {
//		return
//	}
//
//	funcName := runtime.FuncForPC(pc).Name()
//	funcName = path.Base(funcName)
//
//	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
//	return
//}

var timeCfg *now.Config

func WeeklyRange(t time.Time) (string, time.Time, time.Time) {
	_, week := t.ISOWeek()
	return fmt.Sprintf("%d", week), timeCfg.With(t).BeginningOfWeek(), timeCfg.With(t).EndOfWeek()
}

func PreWeeklyRange(t time.Time) (string, time.Time, time.Time) {
	preT := t.Add(-time.Duration(7*24) * time.Hour)
	return WeeklyRange(preT)
}

func BiweeklyRange(t time.Time) (string, time.Time, time.Time) {
	_, week := t.ISOWeek()
	if week%2 == 1 {
		return WeeklyRange(t)
	}
	preWeek, preStart, _ := PreWeeklyRange(t)
	curWeek, _, curEnd := WeeklyRange(t)

	return fmt.Sprintf("%s-%s", preWeek, curWeek), preStart, curEnd
}

func PreBiweeklyRange(t time.Time) (string, time.Time, time.Time) {
	_, week := t.ISOWeek()
	preT := t.Add(-time.Duration(14*24) * time.Hour)
	if week%2 == 1 {
		preT = t.Add(-time.Duration(7*24) * time.Hour)
	}

	return BiweeklyRange(preT)
}

func MonthlyRange(t time.Time) (string, time.Time, time.Time) {
	y, m, _ := t.Date()
	return fmt.Sprintf("%d-%d", y, m), timeCfg.With(t).BeginningOfMonth(), timeCfg.With(t).EndOfMonth()
}

func PreMonthlyRange(t time.Time) (string, time.Time, time.Time) {
	y, m, _ := t.Date()
	if m == 1 {
		y = y - 1
		m = 12
	} else {
		m = m - 1
	}
	nTime := timeCfg.With(time.Date(y, m, 15, 12, 0, 0, 0, getLocation()))

	return fmt.Sprintf("%d-%d", y, m), nTime.BeginningOfMonth(), nTime.EndOfMonth()
}

func init() {
	timeCfg = &now.Config{
		WeekStartDay: time.Monday,
		TimeLocation: getLocation(),
		TimeFormats:  []string{"2006-01-02 15:04:05"},
	}
}

func getLocation() *time.Location {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return location
}

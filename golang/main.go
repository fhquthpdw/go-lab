package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/now"
)

// //
type Resource struct {
	Group     string `json:"group"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
type SeDevopsEks struct {
	AppNamespace string     `json:"appNamespace"`
	Name         string     `json:"name"`
	Prune        bool       `json:"prune"`
	Resources    []Resource `json:"resources"`
}

type T struct {
	Id []struct{} `json:"id"`
}

// argocd API doc: https://argocd-se.shared.cdt.thelegogroup.cn/swagger-ui#operation/ApplicationService_Sync
func main() {
	var a []struct{}
	xxx := T{
		Id: a,
	}
	xv, _ := json.Marshal(xxx)
	fmt.Println(string(xv))
	os.Exit(1)

	patchData := SeDevopsEks{
		AppNamespace: "default",
		Name:         "daochun-demo",
		Prune:        false,
		Resources: []Resource{
			{
				Kind:      "Secret",
				Name:      "argocd-demo-secret",
				Namespace: "argo-sbx",
			},
		},
	}

	s, _ := json.Marshal(patchData)
	fmt.Println("=================")
	fmt.Println(string(s))
	fmt.Println("=================")
}

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
func JsonEncode2String(obj interface{}) (string, error) {
	if jsonByte, err := JsonEncode(obj); err != nil {
		return "", err
	} else {
		return string(jsonByte), nil
	}
}

// JsonEncode JSONEncode to []byte
func JsonEncode(obj interface{}) (jsonByte []byte, err error) {
	jsonByte, err = json.Marshal(obj)
	if err != nil {
		return jsonByte, errors.New("marshal failed")
	}
	return jsonByte, nil
}

func FibonacciNumber(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}

type xmlResponse struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Result  struct {
		XMLName xml.Name `xml:"result"`
		Key     string   `xml:"key"`
	} `xml:"result"`
}

//var mx map[string]string

func mainx() {
	// 输入的 UTC 时间字符串
	utcTimeString := "2023-05-02T18:00:00.0+0000"
	ff := "2006-01-02T15:04:05.9+0000"

	// 解析 UTC 时间字符串为 time.Time 类型
	utcTime, err := time.Parse(ff, utcTimeString)
	if err != nil {
		panic(err)
	}

	// 获取上海时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	// 将 UTC 时间转换为上海时间
	shanghaiTime := utcTime.In(loc)
	fmt.Println("*********************")
	fmt.Println(shanghaiTime)
	fmt.Println("*********************")

	// 输出上海时间字符串
	//fmt.Println(shanghaiTime.Format(ff))
	fmt.Println("===============\r\n\r\n\r\n")
	//
	f := "2006-01-02T15:04:05.9+0000"
	ts := "2023-05-02T16:00:00.0+0000"
	//shanghaiLC, _ := time.LoadLocation("Asia/Shanghai")
	shanghaiLC, _ := time.LoadLocation("UTC")
	r, err := time.ParseInLocation(f, ts, shanghaiLC)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("============ %s ============", r)
	//
	//mx["a"] = "b"
	log.Fatal("\r\n")

	d := []byte(`<response status = 'success'><result><key>KeyKey</key></result></response>`)
	var xmlResponse xmlResponse
	if err := xml.Unmarshal(d, &xmlResponse); err != nil {
		fmt.Println(err)
	}
	fmt.Println(xmlResponse.Status)
	fmt.Println(xmlResponse.Result.Key)

	// <response status = 'success'><result><key>asdfasdf</key></result></response>

	fmt.Println("==================")
	log.Fatal("aaa")
a:
	for x := 0; x <= 10; x++ {
		for i := 0; i <= 10; i++ {
			fmt.Println(x, " => ", i)
			if i == 3 {
				break a
			}
		}
		fmt.Println(x)
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

package httputils

import (
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"strings"
	"time"
)

func TimestampToTime(t int64) time.Time {
	timeTemplate1 := "2006-01-02 15:04:05"
	timeStr := time.Unix(t, 0).Format(timeTemplate1)
	t1Time, _ := time.ParseInLocation(timeTemplate1, timeStr, time.Local)
	return t1Time
}

func StringToTime(s string) time.Time {
	//loc, _ := time.LoadLocation("Local")
	ctime, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err == nil {
		//fmt.Println(ctime)
		return ctime
	}
	return time.Now()
}

func TimeAddMinutes(now time.Time, minute int) time.Time {
	m, _ := time.ParseDuration("1m")
	m1 := now.Add(time.Duration(minute) * m)
	return m1
}

// GetLastMonthTime 获取当前时间前一个月的时间
func GetLastMonthTime() time.Time {
	// 获取当前时间
	now := time.Now()

	// 获取当前时间的年份和月份
	year, month, day := now.Date()

	// 计算上个月的时间（处理跨年）
	if month == 1 {
		return time.Date(year-1, 12, day, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	}
	return time.Date(year, month-1, day, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
}

func Utc0Plus8Time(now time.Time) time.Time {
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	return h1
}

// FormaterString 去掉字符串的空格及,
func FormaterString(str string) string {
	str = strings.ReplaceAll(strings.TrimSpace(str), ",", "")
	return str
}

func DecimalValue(str string) decimal.Decimal {
	value, _ := decimal.NewFromString(str)
	//v := value.Div(decimal.NewFromFloat(math.Pow10(18))).Round(4)
	return value
}

func DecimalValueFromFloat(f64 float64) decimal.Decimal {
	value := decimal.NewFromFloat(f64)
	//v := value.Div(decimal.NewFromFloat(math.Pow10(18))).Round(4)
	return value
}

func DecimalDiv18Value(str string) decimal.Decimal {
	value, _ := decimal.NewFromString(str)
	v := value.Div(decimal.NewFromFloat(math.Pow10(18))).Round(8)
	return v
}

func DecimalDivValue(str string, wei int) decimal.Decimal {
	value, _ := decimal.NewFromString(str)
	v := value.Div(decimal.NewFromFloat(math.Pow10(wei))).Round(4)
	return v
}

func DecimalDiv1024x4Value(str string) decimal.Decimal {
	value, _ := decimal.NewFromString(str)
	v := value.Div(decimal.NewFromFloat(math.Pow(1024, 4))).Round(2)
	return v
}

func FloatValue(str string) float64 {
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
		floatValue = 0
	}
	return floatValue
}

func Truncate(f float64, prec int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return ""
	}
	if prec >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || prec >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:prec]
}

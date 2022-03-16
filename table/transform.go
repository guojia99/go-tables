package table

import (
	"fmt"
	"github.com/gookit/color"
	"reflect"
	"runtime"
	"time"
)

/*
	TransformContent shi
	cn:
		当你不想对原有内容作出改变的时候,错误应当返回nil
	en:
		When you do not want to make changes to the original content, the error should return nil
	E.g:
		func t(in interface{}) (interface{}, error) {
			times, ok := in.(time.Time)
			if !ok {
				return in, nil  <<<<-----
			}
			return times.Format(time.RFC822), nil
		}
*/
type TransformContent func(in interface{}) interface{}

type TransformContents []TransformContent

/*
	cn: 将会对每一个变更方法进行一一执行
	en: Convert will implement change method one by one
*/
func (t TransformContents) Convert(in interface{}) interface{} {
	for _, f := range t {
		in = f(in)
	}
	return in
}

/*
	DefaultTransformContentByTime
	cn:
	en:
*/
const (
	second = 1
	minute = second * 60
	hour   = minute * 60
	day    = hour * 24
	month  = day * 30
	year   = month * 12
)

var timeFormal = map[bool]map[string]string{
	false: {
		"second": " second ago",
		"minute": " minute ago",
		"hour":   " hour ago",
		"day":    " days ago",
		"month":  " month ago",
		"year":   " year ago",
	},
	true: {
		"second": " second after",
		"minute": " minute after",
		"hour":   " hour after",
		"day":    " days after",
		"month":  " month after",
		"year":   " year after",
	},
}

func getTimeFormal(isFuture bool, timeF string, stamp int64) string {
	return fmt.Sprint(stamp, timeFormal[isFuture][timeF])
}

func DefaultTransformContentByTime(in interface{}) interface{} {
	t, ok := in.(time.Time)
	if !ok {
		return in
	}
	timeDifference := int64(time.Now().Second() - t.Second())
	isFuture := false
	if timeDifference < 0 {
		isFuture = true
	}
	if timeDifference <= minute {
		return getTimeFormal(isFuture, "second", timeDifference)
	}
	if timeDifference <= hour {
		return getTimeFormal(isFuture, "minute", timeDifference/minute)
	}
	if timeDifference <= day*2 {
		return getTimeFormal(isFuture, "hour", timeDifference/hour)
	}
	if timeDifference <= month {
		return getTimeFormal(isFuture, "day", timeDifference/day)
	}
	if timeDifference <= year {
		return getTimeFormal(isFuture, "month", timeDifference/month)
	}
	return getTimeFormal(isFuture, "year", timeDifference/year)
}

/*
	DefaultTransformContentByColor
	中文：
		依据不同类型进行颜色的区分输出,输出最终格式为string
		目前使用的颜色只支持linux和macOs系统, windows系统需要自己参考定义好一个颜色引擎
		+ 调用优先级: CustomizeTypeTo > typeTo > defaultColor
			- Enable: 启用与不启用
			- typeTo: 是依据对不同的【基础数据类型】进行颜色输出时的处理
			- customizeTypeTo: 依据输入的数据结构反射出的结构进行输出颜色
	en:
		Differentiate the output of colors according to different types, and the final output format is string
		The currently used colors only support linux and `macOS` systems, `windows`system needs refer to define a color engine by itself
		+ call priority: CustomizeTypeTo > typeTo > defaultColor
			- Enable: enable and disable
			- typeTo: it is based on the processing of color output for different [basic data types]
			- customizeTypeTo: output color according to the structure reflected from the input data structure

*/
func DefaultTransformContentByColor(in interface{}) interface{} {
	if !DefaultColorStylesClient.Enable {
		return in
	}
	switch runtime.GOOS {
	case `linux`, `darwin`:
		DefaultColorStylesClient.Enable = true
	default:
		DefaultColorStylesClient.Enable = false
	}
	tp := reflect.TypeOf(in)
	if cor, ok := DefaultColorStylesClient.customizeTypeTo[tp.String()]; ok {
		return cor.Sprint(in)
	}
	if cor, ok := DefaultColorStylesClient.typeTo[tp.Kind()]; ok {
		return cor.Sprint(in)
	}
	if DefaultColorStylesClient.defaultColor != nil {
		return DefaultColorStylesClient.defaultColor.Sprint(in)
	}
	return in
}

type defaultColorStyles struct {
	Enable          bool
	customizeTypeTo map[string]color.Color
	typeTo          map[reflect.Kind]color.Color
	defaultColor    *color.Color
}

var DefaultColorStylesClient = &defaultColorStyles{
	Enable:          true,
	defaultColor:    nil,
	typeTo:          map[reflect.Kind]color.Color{},
	customizeTypeTo: map[string]color.Color{},
}

func DisEnableDefaultColor() { DefaultColorStylesClient.Enable = false }
func EnableDefaultColor()    { DefaultColorStylesClient.Enable = true }
func SetDefaultColor(cor color.Color, in ...interface{}) {
	if len(in) == 0 {
		DefaultColorStylesClient.defaultColor = &cor
	}

	for _, val := range in {
		if kind, ok := val.(reflect.Kind); ok {
			DefaultColorStylesClient.typeTo[kind] = cor
			return
		}
		tp := reflect.TypeOf(in)
		if tp.Kind() <= reflect.UnsafePointer {
			DefaultColorStylesClient.typeTo[tp.Kind()] = cor
		}
		DefaultColorStylesClient.customizeTypeTo[tp.String()] = cor
	}
}
func DeleteDefaultColor(in ...interface{}) {
	if len(in) == 0 {
		DefaultColorStylesClient.defaultColor = nil
	}
	for _, val := range in {
		if kind, ok := val.(reflect.Kind); ok {
			if _, have := DefaultColorStylesClient.typeTo[kind]; have {
				delete(DefaultColorStylesClient.typeTo, kind)
			}
			return
		}
		tp := reflect.TypeOf(in)
		if _, have := DefaultColorStylesClient.typeTo[tp.Kind()]; have {
			delete(DefaultColorStylesClient.typeTo, tp.Kind())
		}
		if _, have := DefaultColorStylesClient.customizeTypeTo[tp.String()]; have {
			delete(DefaultColorStylesClient.customizeTypeTo, tp.String())
		}
	}
}

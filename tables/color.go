package tables

import (
	"fmt"
	"github.com/gookit/color"
	"reflect"
)

// ColorStyles
/*
	中文：
		typeTo 是依据对不同的【基础数据类型】进行颜色输出时的处理
		customizeTypeTo 依据输入的数据结构反射出的结构进行输出颜色
		UseTypeTo 启用与不启用
		调用优先级CustomizeTypeTo > typeTo >Default
	en:

*/
type ColorStyles struct {
	Enable          bool
	customizeTypeTo map[string]color.Color
	typeTo          map[reflect.Kind]color.Color
	Default         *color.Color
}

func NewColorStyles() *ColorStyles {
	return &ColorStyles{
		Enable:          true,
		Default:         nil,
		typeTo:          map[reflect.Kind]color.Color{},
		customizeTypeTo: map[string]color.Color{},
	}
}

func (c *ColorStyles) SetDefaultColor(cor color.Color) {
	c.Default = &cor
}

func (c *ColorStyles) ClearDefaultColor() {
	c.Default = nil
}

func (c *ColorStyles) Set(in interface{}, cor color.Color) {
	if kind, ok := in.(reflect.Kind); ok {
		c.typeTo[kind] = cor
		return
	}
	tp := reflect.TypeOf(in)
	if tp.Kind() <= reflect.UnsafePointer {
		c.typeTo[tp.Kind()] = cor
	}
	c.customizeTypeTo[tp.String()] = cor
}

func (c *ColorStyles) Delete(in interface{}) {
	if kind, ok := in.(reflect.Kind); ok {
		if _, have := c.typeTo[kind]; have {
			delete(c.typeTo, kind)
		}
		return
	}
	tp := reflect.TypeOf(in)
	if _, have := c.typeTo[tp.Kind()]; have {
		delete(c.typeTo, tp.Kind())
	}
	if _, have := c.customizeTypeTo[tp.String()]; have {
		delete(c.customizeTypeTo, tp.String())
	}
}

func (c *ColorStyles) Parse(in interface{}) string {
	if !c.Enable {
		return fmt.Sprintf("%v", in)
	}
	tp := reflect.TypeOf(in)
	if cor, ok := c.customizeTypeTo[tp.String()]; ok {
		return cor.Sprint(in)
	}
	if cor, ok := c.typeTo[tp.Kind()]; ok {
		return cor.Sprint(in)
	}
	if c.Default != nil {
		return c.Default.Sprint(in)
	}
	return fmt.Sprintf("%v", in)
}

// OriginalLen 返回该字符串原始长度
func OriginalLen(in string) int {
	return len(color.ClearCode(in))
}

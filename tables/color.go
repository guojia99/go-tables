package tables

import (
	"fmt"
	"github.com/gookit/color"
	"reflect"
)

type ColorStyles interface {
	// 设置默认颜色值
	SetDefaultColor(color.Color)
	// 清空默认颜色值
	ClearDefaultColor()
	// 设置颜色
	Set(interface{}, color.Color)
	// 删除颜色
	Delete(interface{})
	// 序列化颜色
	Parse(interface{}) string
	// 序列化后的颜色原本的长度
	RealLength(string) int
}

// DefaultColorStyles
/*
	中文：
		typeTo 是依据对不同的【基础数据类型】进行颜色输出时的处理
		customizeTypeTo 依据输入的数据结构反射出的结构进行输出颜色
		UseTypeTo 启用与不启用
		调用优先级CustomizeTypeTo > typeTo >defaultColor
	en:
*/
type DefaultColorStyles struct {
	Enable          bool
	customizeTypeTo map[string]color.Color
	typeTo          map[reflect.Kind]color.Color
	defaultColor    *color.Color
}

func NewDefaultColorStyles() *DefaultColorStyles {
	return &DefaultColorStyles{
		Enable:          true,
		defaultColor:    nil,
		typeTo:          map[reflect.Kind]color.Color{},
		customizeTypeTo: map[string]color.Color{},
	}
}
func (c *DefaultColorStyles) SetDefaultColor(cor color.Color) {
	c.defaultColor = &cor
}
func (c *DefaultColorStyles) ClearDefaultColor() {
	c.defaultColor = nil
}
func (c *DefaultColorStyles) Set(in interface{}, cor color.Color) {
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
func (c *DefaultColorStyles) Delete(in interface{}) {
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
func (c *DefaultColorStyles) Parse(in interface{}) string {
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
	if c.defaultColor != nil {
		return c.defaultColor.Sprint(in)
	}
	return fmt.Sprintf("%v", in)
}
func (c *DefaultColorStyles) RealLength(in string) int {
	return len(color.ClearCode(in))
}

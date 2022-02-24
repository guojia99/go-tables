package tables

import (
	"fmt"
	"github.com/gookit/color"
	"reflect"
	"testing"
)

func TestColorStylesSetAndDelete(t *testing.T) {
	c := NewDefaultColorStyles()

	// 1、存在标准数据结构时删除
	c.Set(reflect.Int, color.Cyan)
	c.Delete(reflect.Int)

	// 2、不存在标准数据结构时删除
	c.Delete(reflect.String)

	// 3、存在自定义数据结构时删除
	type MyStruct struct {
		testF string
	}
	type MyString string
	c.Set(&MyStruct{}, color.Cyan)
	c.Delete(&MyStruct{})
	c.Set(MyString("test"), color.Cyan)
	c.Delete(MyString("a"))

	// 4、不存在标准数据结构时删除
	type MyInt int
	c.Delete(MyInt(1))

	// 5、默认颜色
	c.SetDefaultColor(color.Magenta)
	c.ClearDefaultColor()

	// 6、标准数据结构
	c.Set("", color.Cyan)
	c.Set(int64(1), color.Cyan)
}

func TestColorStylesParse(t *testing.T) {
	c := NewDefaultColorStyles()

	c.Set(1, color.Magenta)
	c.Set("", color.Cyan)

	type MyInt int
	c.Set(MyInt(1), color.Red)
	type MyStruct struct {
		testI int
		testB bool
		testS string
	}
	c.Set(MyStruct{}, color.Yellow)

	var data = []interface{}{
		1, "test", int64(1), MyStruct{
			testB: true,
			testS: "test",
			testI: 123,
		}, MyInt(123),
	}

	for _, val := range data {
		fmt.Println(c.Parse(val), []byte(c.Parse(val)))
	}
}

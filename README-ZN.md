# go-tables 

简体中文 | [English](./README.md) 

> - Table Out 命令行表格输出工具
>
>   - 一个用于命令行输出美观的表格的库
>   - 支持不同的输入表格式
>   - 支持表头自定义
>

### 2、快速使用

> ```bash
> go get -u -d github.com/guojia99/go-tables/table
> ```

- 目前支持以下六种数据结构输出
```
Struct                  struct{}
Map                     map[interface{}]interface{}
MapSlice                []map[interface{}]interface{}
StructSlice             []struct{}
Slice                   []interface{}
Slice2D                 [][]interface{}
```

你可以通过快速调用的方法，直接输出你原有的数据


```go
var opt = &table.Option{
    Contour: table.DefaultContour,
    Align:   table.AlignCenter,
}
tb, _ := table.SimpleTable(data, opt)
fmt.Println(tb)
```

- ##### Struct to table

```go
type structTable struct {
	Str    string
	Val    string `table:"value"`
	Num    int    `json:"number"`
	NoUse  string `json:"-"`
	NoUse2 string `table:"-"`
}
data := structTable{
    Str:    "value",
    Val:    "val",
    Num:    111,
    NoUse:  "nouse",
    NoUse2: "nouse",
}
tb, _ := table.SimpleTable(data, opt)
fmt.Println(tb)

# out:
┌────────┬───────┐
|   #    | value |
├────────┼───────┤
|  Str   | value |
| value  |  val  |
| number |  111  |
└────────┴───────┘
```

![](doc/doc-images/struct-table.png)



- ##### Map to table

```go
data := map[string]string{
    "key1": "value1",
    "key2": "value2",
    "key3": "value3",
    "key4": "value4",
    "key5": "value5",
    "key6": "value6",
}
tb, _ := table.SimpleTable(data, opt)
fmt.Println(tb)

# out:
┌──────┬────────┐
| key  | value  |
├──────┼────────┤
| key1 | value1 |
| key2 | value2 |
| key3 | value3 |
| key4 | value4 |
| key5 | value5 |
| key6 | value6 |
└──────┴────────┘

data2 := map[string]interface{}{
    "number":  1,
    "string":  "guojia",
    "float":   2.4,
    "slide":   []int{1, 2, 3, 4},
    "complex": complex(1, -1),
    "key6": struct {
        a string
    }{a: "123"},
}
tb, _ := table.SimpleTable(data, opt)
fmt.Println(tb)

# out:
┌─────────┬───────────┐
|   key   |   value   |
├─────────┼───────────┤
| complex |  (1-1i)   |
|  key6   |   {123}   |
| number  |     1     |
| string  |  guojia   |
|  float  |    2.4    |
|  slide  | [1 2 3 4] |
└─────────┴───────────┘
```

![](doc/doc-images/map-table.png)

- ##### map slice to table

```go
data := map[string][]string{
		"key1": {"key1-v", "key1-v", "key1-v", "key1-v", "key1-v", "key1-v"},
		"key2": {"key2-v2", "key2-v2", "key2-v2", "key2-v2"},
		"key3": {"key3-v3", "key3-v3", "key3-v3", "key3-v3"},
		"key4": {"key4-v4", "key4-v4", "key4-v4", "key4-v4"},
		"key5": {"key5-v5", "key5-v4", "key5-v4", "key5-v4"},
		"key6": {"key6-v6", "key6-v4", "key6-v4", "key6-v4"},
		"key7": {"key7-v6", "key7-v4", "key7-v4"},
}
tb, _ := table.SimpleTable(data, opt)
fmt.Println(tb)

# out put:
┌─────────┬─────────┬────────┬─────────┬─────────┬─────────┬─────────┐
|  key6   |  key7   |  key1  |  key2   |  key3   |  key4   |  key5   |
├─────────┼─────────┼────────┼─────────┼─────────┼─────────┼─────────┤
| key6-v6 | key7-v6 | key1-v | key2-v2 | key3-v3 | key4-v4 | key5-v5 |
| key6-v4 | key7-v4 | key1-v | key2-v2 | key3-v3 | key4-v4 | key5-v4 |
| key6-v4 | key7-v4 | key1-v | key2-v2 | key3-v3 | key4-v4 | key5-v4 |
| key6-v4 |         | key1-v | key2-v2 | key3-v3 | key4-v4 | key5-v4 |
|         |         | key1-v |         |         |         |         |
|         |         | key1-v |         |         |         |         |
└─────────┴─────────┴────────┴─────────┴─────────┴─────────┴─────────┘
```

![](doc/doc-images/mapslice-table.png)

- ##### Struct Slice to table

```go
type structSliceTable struct {
	Str    string
	Val    string `table:"value"`
	Num    int    `json:"number"`
	NoUse  string `json:"-"`
	NoUse2 string `table:"-"`
}
data := []structSliceTable{
    {"data1", "val1", 1, "no1", "no2"},
    {"data2", "val2", 2, "no2", "no3"},
    {"data3", "val3", 3, "no3", "no4"},
    {"data4", "val4", 4, "no4", "no5"},
    {"data5", "val5", 5, "no5", "no6"},
}
tb, _ := table.SimpleTable(data, opt)
fmt.Println(tb)

# out:
┌───────┬───────┬────────┐
|  Str  | value | number |
├───────┼───────┼────────┤
| data1 | val1  |   1    |
| data2 | val2  |   2    |
| data3 | val3  |   3    |
| data4 | val4  |   4    |
| data5 | val5  |   5    |
└───────┴───────┴────────┘
```

![](doc/doc-images/structslide-table.png)

- ##### Slice to table

```go
data := []string{
	"Each package fulfils a single purpose",
	"Handle errors explicitly",
	"Return early rather than nesting deeply",
	"Leave concurrency to the caller",
	"Before you launch a goroutine, know when it will stop",
	"Avoid package level state",
	"Simplicity matters",
	"Write tests to lock in the behaviour of your package’s API",
	"If you think it’s slow, first prove it with a benchmark",
	"Moderation is a virtue",
}
tb, _ := table.SimpleTable(data, opt)
fmt.Println(tb)

# out:
┌────┬────────────────────────────────────────────────────────────┐
| No |                           value                            |
├────┼────────────────────────────────────────────────────────────┤
| 0  |           Each package fulfils a single purpose            |
| 1  |                  Handle errors explicitly                  |
| 2  |          Return early rather than nesting deeply           |
| 3  |              Leave concurrency to the caller               |
| 4  |   Before you launch a goroutine, know when it will stop    |
| 5  |                 Avoid package level state                  |
| 6  |                     Simplicity matters                     |
| 7  | Write tests to lock in the behaviour of your package’s API |
| 8  |  If you think it’s slow, first prove it with a benchmark   |
| 9  |                   Moderation is a virtue                   |
└────┴────────────────────────────────────────────────────────────┘
```

![](doc/doc-images/slide-table.png)

- ##### Slide2D to table

```go
data := [][]string{
		{"DATA1", "DATA2", "DATA3"},
		{"DATA4", "DATA5", "DATA6"},
		{"DATA7", "DATA8", "DATA9"},
		{"DATA10", "DATA11", "DATA12"},
}
tb, _ := table.SimpleTable(data, opt)
fmt.Println(tb)

# out:
┌────────┬────────┬────────┐
├────────┼────────┼────────┤
| DATA1  | DATA2  | DATA3  |
| DATA4  | DATA5  | DATA6  |
| DATA7  | DATA8  | DATA9  |
| DATA10 | DATA11 | DATA12 |
└────────┴────────┴────────┘
```



![](doc/doc-images/slide2D-table.png)


----



![](doc/doc-images/slide2D-table.png)




----

#  v2 开发中

- 支持可交互
- 支持缓存
- 更多的自定义接口
- 支持并发
- 全函数自测
- 支持颜色绘制
- 更多的数据处理
- todo...


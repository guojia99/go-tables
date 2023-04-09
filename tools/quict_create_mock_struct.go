/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:45.
 * Author:  guojia(https://github.com/guojia99)
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const format = `
/*
 * Copyright (c) 2023 gizwits.com All rights reserved.
 * Created: 2023/4/3 下午6:20.
 * Author: guojia(zjguo@gizwits.com)
 */

package mock

type TestStruct1 struct{
	%s
}
`
	out := ""
	for i := 'A'; i < 'Z'; i++ {
		for j := 0; j < 10; j++ {
			out += fmt.Sprintf("%s%d string `json:\"%s%d\"`\n", string(i), j, strings.ToLower(string(i)), j)
		}
	}

	_ = os.WriteFile("mock/tables_mock_struct.go", []byte(fmt.Sprintf(format, out)), 0644)
}

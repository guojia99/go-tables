/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import "errors"

var (
	ErrNotValidValue = errors.New("the content is not valid")
	ErrAddress       = errors.New("wrong address")
	ErrRange         = errors.New("illegal operating range")
)

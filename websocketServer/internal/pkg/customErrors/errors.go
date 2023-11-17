package customerrors

import "errors"

var ErrRepeat = errors.New("user already exists")
var ErrEmpty = errors.New("this user doesn't exists")

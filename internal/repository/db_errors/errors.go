package dberrors

import "errors"

var ErrNoRows = errors.New("can't find any record with such parameters")

package gofu

import (
	"fmt"
)

type TError struct{
	pos TPos
	message string
}

func Error(pos TPos, spec string, args...interface{}) TError {
	var e TError
	e.Init(pos, spec, args...)
	return e
}

func (self *TError) Init(pos TPos, spec string, args...interface{}) *TError {
	self.pos = pos
	self.message = fmt.Sprintf(spec, args...)
	return self
}

func (self TError) Error() string {
	return fmt.Sprintf("error in %v: %v", self.pos, self.message)}

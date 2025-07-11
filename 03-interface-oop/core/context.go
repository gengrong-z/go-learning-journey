package core

import "go-learning-journey/03-interface-oop/battlelog"

type Context struct {
	log *battlelog.BattleLog
}

func NewContext() *Context {
	return &Context{}
}

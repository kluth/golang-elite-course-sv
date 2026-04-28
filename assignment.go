package assignment

type Handler interface { SetNext(Handler); Handle(string) bool }
type BaseHandler struct { next Handler }
func (b *BaseHandler) Handle(s string) bool { panic("implement me") }

type Command interface { Execute() string }
type LightOn struct{}
func (l LightOn) Execute() string { panic("implement me") }

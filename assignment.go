package assignment

type Visitor interface { VisitCircle() string }
type Circle struct{}
func (c Circle) Accept(v Visitor) string { panic("implement me") }

type TreeType struct { Color string }
var types = make(map[string]*TreeType)
func GetTreeType(c string) *TreeType { panic("implement me") }

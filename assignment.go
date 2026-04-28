package assignment

type Work interface { Step1(); Step2() }
func ExecuteWork(w Work) { panic("implement me") }

type Iterator interface { HasNext() bool; Next() string }
type Group struct { items []string; cur int }
func (g *Group) HasNext() bool { panic("implement me") }
func (g *Group) Next() string { panic("implement me") }

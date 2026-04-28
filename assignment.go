package assignment

type Strategy interface { Route() string }
type Walk struct{}
func (w Walk) Route() string { return "walk" }
type Navigator struct { strategy Strategy }
func (n Navigator) Exec() string { panic("implement me") }

type State interface { Handle() string }
type Pending struct{}
func (p Pending) Handle() string { return "pending" }
type Order struct { state State }
func (o *Order) SetState(s State) { o.state = s }
func (o Order) Process() string { panic("implement me") }

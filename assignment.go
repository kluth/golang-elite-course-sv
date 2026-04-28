package assignment

type Product interface { GetName() string }
type ConcreteProduct struct{}
func (p ConcreteProduct) GetName() string { return "concrete" }
func ProductFactory(t string) Product { panic("implement me") }

type Car struct { Color string }
type CarBuilder struct { car Car }
func (b *CarBuilder) SetColor(c string) *CarBuilder { panic("implement me") }
func (b *CarBuilder) Build() Car { panic("implement me") }

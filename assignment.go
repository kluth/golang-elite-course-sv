package assignment

type Sub1 struct{}
func (s Sub1) Op() string { return "1" }
type Sub2 struct{}
func (s Sub2) Op() string { return "2" }
type Facade struct { s1 Sub1; s2 Sub2 }
func (f Facade) Unified() string { panic("implement me") }

type SecretData interface { Get() string }
type RealData struct{}
func (r RealData) Get() string { return "secret" }
type Proxy struct { real RealData; Role string }
func (p Proxy) Get() string { panic("implement me") }

package assignment

type Node interface { Size() int }
type File struct { val int }
func (f File) Size() int { return f.val }
type Folder struct { children []Node }
func (f *Folder) Add(n Node) { panic("implement me") }
func (f Folder) Size() int { panic("implement me") }

type Service interface { Do() string }
type ConcreteService struct{}
func (s ConcreteService) Do() string { return "do" }
type LoggingDecorator struct { s Service }
func (d LoggingDecorator) Do() string { panic("implement me") }

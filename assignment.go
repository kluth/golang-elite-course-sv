package assignment

type Observer interface { Update(string) }
type Subject struct { obs []Observer }
func (s *Subject) Notify(msg string) { panic("implement me") }

type Mediator interface { Send(msg string, user string) string }
type ChatRoom struct{}
func (c ChatRoom) Send(msg, user string) string { panic("implement me") }

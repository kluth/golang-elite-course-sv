package assignment

type Notifier interface { Notify() string }

type EmailSender struct{}

func (e *EmailSender) Notify() string { panic("implement me") }

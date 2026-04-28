package assignment

type ModernService interface { Call() string }
type LegacySystem struct{}
func (s LegacySystem) LegacyCall() string { return "legacy" }
type LegacyAdapter struct { system LegacySystem }
func (a LegacyAdapter) Call() string { panic("implement me") }

type Printer interface { Print() string }
type Computer struct { printer Printer }
func (c Computer) PrintData() string { panic("implement me") }

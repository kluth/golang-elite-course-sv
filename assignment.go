package assignment

type Database struct { Connection string }
var instance *Database
func GetInstance() *Database { panic("implement me") }

type User struct { Name string }
func (u *User) Clone() *User { panic("implement me") }

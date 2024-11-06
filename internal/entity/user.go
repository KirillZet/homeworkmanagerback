package entity

import "time"

type User struct {
	UserID    UserID
	Name      string
	Surname   *string
	Email     string
	Password  string
	Role      Role
	GroupId  int64
	СreatedAt time.Time
}

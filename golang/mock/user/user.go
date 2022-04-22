package user

import "mock/person"

func NewUser(p person.Man) *User {
    return &User{Person: p}
}

type User struct {
    Person person.Man
}

func (u *User) Walk(road string) string {
    return u.Person.Walk(road)
}

func (u *User) Talk(sentence string) string {
    return u.Person.Talk(sentence)
}

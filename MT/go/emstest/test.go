package emstest

import (
	"fmt"
	"time"
)

func Test() {
	fmt.Println("===========teset===11===333====")
	u1 := User{
		name: "John", dep: "QA", age: 24, createTime: time.Now(),
	}
	fmt.Println("u1.dep=", u1.dep)
	u1.updateDep("BA")
	fmt.Println("u1.new_dep=", u1.dep)

	newdep := showDep(&u1)
	fmt.Println(newdep)
	fmt.Println(showAge(&u1))
	fmt.Println(addAge(&u1))
	fmt.Println((&u1))

}

//CompileDaemon --build="go build -o ems cmd/ems/main.go" --command=./ems

type User struct {
	name, dep  string
	age        int
	createTime time.Time
}

func (u *User) updateDep(newDep string) {
	u.dep = newDep
}

func showDep(u *User) string {
	return u.dep
}

func showAge(u *User) *User {

	u.age++
	return u
}

func addAge(u *User) User {
	u.age++
	return *u
}

func addNewAge(u *User) User {

	u.age++
	return *u
}

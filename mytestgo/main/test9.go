package main

import "fmt"

type User struct {
	Name string
	Age int
}

var userDB  = map[int] User{
	1: User{"Zdz1", 18},
	9: User{"Zdz2", 20},
	8: User{"Zdz3", 27},
}

func QuseryUser(id int) (User, error)  {
	if u, ok := userDB[id]; ok{
		return u, nil
	}
	return User{}, fmt.Errorf("id %d is not in db", id)
}
func main() {
	u , err := QuseryUser(9)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("name: %s, age: %d \n", u.Name, u.Age)
}
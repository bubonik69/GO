package main

import "fmt"

func main() {
	fmt.Println("Hi, its lesson dedicated type of struct")

	type User struct {
		Id     int
		Name   string
		Pass   string
		friend []*User
	}
	var a int = 1
	var b *int = &a
	fmt.Println(a, b)
	*b = *b + 1
	fmt.Printf("%v - %v", a, &b)

	type Base struct {
		Name string
		id int
	}

	c:= Base{
		Name: " vasia ",
		id: 123}




	type  user struct{
		Name string
		Surname string
		Age int
	}
	// дополнительные функции
	//func (p user) String() string {
	//return fmt.Sprintf("%d : %s: %v", p.Age, p.Name, p.Surname)
	//	return "hi"
	//}
	func (p user) getName()  {
		fmt.Println(p.Name)
	}


}


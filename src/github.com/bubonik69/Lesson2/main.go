package main
// импортирование

import (
	"fmt"
)
// глобальные переменные
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
func getName(p user) {
	fmt.Println( p.Name)
}
// сортировка по имени
func main(){


gamers:= []user{
		{"Artem" ,"Nickulin", 22},
		{"Ar" ,"pic", 23},
		{"men" ,"Nun", 24},
		{"mike" ,"Gin", 25},
	}

	for i:=0;i<len(gamers);i++{
		//gamers[i].getName()
		//getName(gamers[i])
		if (gamers[i].Name)!=""{
			fmt.Println(gamers[i])
		}
	}

//	sort.Ints()
}














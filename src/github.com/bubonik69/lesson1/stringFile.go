package main

import (
	"fmt"
	"unicode/utf8"
)
// возвращаем количество символов в руне
func dz1 (s string) (qs int){
qs=utf8.RuneCountInString(s)
return qs
}

func RuneAndString() {
	str := "Хай m3y frien8d!"
	ru := []rune(str)
	var len int = dz1(str)
	var extractStr string

	fmt.Printf("Символов в строке '%s' = %d \n", str, len)
	fmt.Println("вывод руны", ru)
	for i, runeValue := range ru {
		fmt.Printf("%q starts at byte position %d\n", runeValue, i)
		// ищем цифру
		ruToint := int(runeValue) - 48
		if (ruToint >= 0) && (ruToint < 10) {
			fmt.Println(ruToint)
			for k := 0; k < ruToint; k++ {
				extractStr = extractStr + string(ru[i-1])
				//fmt.Printf(string(ru[i-1]))
			}
		}else{
			extractStr= extractStr +string(runeValue)
		}

	}

	println(extractStr)
}
func Slovarik(){
	// книга - автор
	book:=map[string]string{"Пелевин": "Ночь ясна",
		"Tolkien" : "Frodo Baggins",
		"Пушкин" : "Kolobok cok-cok",
		"aaa": "bbb"}
	names:=make([]string,0,len(book))
	for _, name:=range book{
		names=append(names,name)
		fmt.Println(name)
	}
	aftors:=make([]string,0,len(book))

	for aftor,_:=range book{
		aftors=append(aftors,aftor)
		fmt.Println(aftor)
	}

	fmt.Println(names)
	fmt.Println(aftors)
}

//func main(){
//	Slovarik()
//	RuneAndString()
//}


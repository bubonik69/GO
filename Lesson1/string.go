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



func main() {
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
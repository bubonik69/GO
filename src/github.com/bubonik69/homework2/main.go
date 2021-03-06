package main

import (
	"fmt"
	"sort"
	"strings"
)
//func clearSlices ([]string)
type wordbook struct {
	word string
	sumCount  int}


func main() {
	text := "В ГАИ сообщили, что начиная с 1 марта в соответствии с новым административным законодательством предусмотрена возможность привлечения собственников транспортных средств к ответственности за все нарушения против безопасности движения и эксплуатации транспорта на основании их фиксации специальными техническими средствами, работающими в автоматическом режиме. Помимо фотофиксации превышения скорости и правил остановки/стоянки организовано привлечение к ответственности за управление без техосмотра. Камеры фиксируют факт участия транспортного средства в дорожном движении, распознают регистрационный знак и сверяют его с базой данных АИС «Белтехосмотр» на наличие разрешения на допуск к участию в дорожном движении. При отсутствии разрешения в отношении собственника выносится соответствующее постановление. Если нарушение зафиксировано впервые, владелец освобождается от административной ответственности с вынесением ему предупреждения. В следующий раз собственник будет привлечен к ответственности в виде штрафа в размере 0,5 базовой величины. При этом к ответственности привлекается как физическое, так и юридическое лицо — собственник или владелец транспортного средства. В соответствии с положениями нового КоАП, если камеры зафиксируют несколько таких нарушений в течение суток (с 00:00 до 23:59), собственник транспортного средства будет привлечен к административной ответственности только один раз. В случае несогласия с принятым решением оно может быть обжаловано в подразделении ГАИ по месту жительства в течение месяца со дня получения копии постановления. Если вы получили «письмо счастья» за отсутствие ТО — сообщите, пожалуйста, нам (в Telegram или на dl@onliner.by)."
	textSlice:= strings.Split(text," ")
	var mywords string


	// выводим слова / переменные в тексте
	/*for index,letter:=range textSlice{
		fmt.Printf("%v  -,%v \n", letter, index)
	}
	*/

	sumArr:=make([]int,len(textSlice))
	for index:=0;index<len(textSlice);index++{
		if (strings.Count(text,textSlice[index])>1)&& (strings.Count(mywords,textSlice[index])==0) {
			mywords=mywords + " " + textSlice[index] // слайс с часто повторяющимеся словами
			fmt.Printf("слово '%v', с индексом,%v встречаеться в строке %v раз \n",textSlice[index],index,strings.Count(text,textSlice[index]))
			sumArr[index]=strings.Count(text,textSlice[index])
			sort.Ints(sumArr)
		}
	}
	// ищем само часто повторяющиеся слова
	for i:=0;i<5;i++ {
		for index := 0; index < len(textSlice); index++ { // запускаем цикл по всем словам
			if (strings.Count(text, textSlice[index]) == sumArr[len(textSlice)-i-1]){
				fmt.Println(" ------->", i+1,  textSlice[index],"встречается", strings.Count(text, textSlice[index]) )
				break
			}
}
}


	fmt.Println("len - text1" ,len(text)) // количесто байт - длина строки
	fmt.Println("слов в тексте", len(textSlice)) // считаем - выводим сколько слов в тексте , разделенных пробелом
	//fmt.Println("слов в переменной 'text'", strings.Count(text," "))//  считаем-выводим  сколько пробелов в тексте
	fmt.Printf("выборка из часто встречающихся слов - '%v'",mywords) // выборка из часто встречающихся слов
	//fmt.Println(sumArr) массив с количеством часто встречающихся слов
}

package main

import (
	"fmt"
	"regexp"
	"sort"
)

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

// TODO: Handle panic
var delimiter = regexp.MustCompile("[ \\n\\r\\t]")

type WordsByCont struct {
	Count     int
	WordsList []string
}

type WordCounter struct {
	Word  string
	Count int
}

func getSliceFromDictionary(dictionary map[string]int) []WordCounter {
	arr := []WordCounter{}

	for key, value := range dictionary {
		arr = append(arr, *&WordCounter{Word: key, Count: value})
	}
	return arr
}

func groupWordsByFrequency(arr []WordCounter) map[int][]string {
	frequencyDictionary := map[int][]string{}

	for _, value := range arr {
		if _, ok := frequencyDictionary[value.Count]; ok {
			frequencyDictionary[value.Count] = append(frequencyDictionary[value.Count], value.Word)
		} else {
			frequencyDictionary[value.Count] = []string{value.Word}
		}
	}

	return frequencyDictionary
}

func sortWords(arr []WordCounter) []WordCounter {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Count > arr[j].Count
	})

	fmt.Print(arr)

	return arr
}

func Top10(text string) []string {
	dictionary := map[string]int{}

	for _, word := range delimiter.Split(text, -1) {
		if len(word) == 0 || delimiter.MatchString(word) {
			continue
		}
		if number, ok := dictionary[word]; ok {
			dictionary[word] = number + 1
		} else {
			dictionary[word] = 1
		}
	}

	groupWordsByFrequency(getSliceFromDictionary(dictionary))

	return nil
}

func main() {
	Top10(text)
}

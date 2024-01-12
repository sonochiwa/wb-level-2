package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	newAnagrams := make(map[string][]string)

	for _, word := range words {
		wordToLower := strings.ToLower(word) // Приводим слово к нижнему регистру
		sorted := sortWord(wordToLower)
		if _, exists := anagrams[sorted]; exists {
			anagrams[sorted] = append(anagrams[sorted], word)
		} else {
			anagrams[sorted] = []string{word}
		}
	}

	// Чтобы ключом в словаре было первое встретившееся слово из множества
	for key := range anagrams {
		if len(anagrams[key]) > 1 {
			newAnagrams[anagrams[key][0]] = anagrams[key]
		}
	}

	return newAnagrams
}

func sortWord(word string) string {
	runes := []rune(word)
	slices.Sort(runes)
	return string(runes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(words)

	for key, value := range anagrams {
		fmt.Println(key, value)
	}
}

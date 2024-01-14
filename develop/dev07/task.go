package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	// Создаем канал, который будем возвращать
	orChannel := make(chan interface{})

	// Запускаем горутину для каждого входящего канала
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			// Копируем данные из входящего канала в результирующий
			for val := range ch {
				orChannel <- val
			}
		}(ch)
	}

	// Запускаем горутину, которая закроет результирующий канал при закрытии любого из входящих каналов
	go func() {
		// Создаем канал для сигнала о закрытии
		done := make(chan struct{})
		defer close(done)

		// Ожидаем закрытия любого из входящих каналов или завершения работы всех горутин
		select {
		case <-done:
		case <-channels[len(channels)-1]:
		}

		// Закрываем результирующий канал
		close(orChannel)
	}()

	return orChannel
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})

	go func() {
		defer close(c)
		time.Sleep(after)
	}()

	return c
}

func main() {
	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Second),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}

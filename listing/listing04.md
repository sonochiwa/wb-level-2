Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Будет deadlock. For range - бесконечный цикл. Из-за того, что мы не закрываем
канал ch после цикла for i := 0; i <10; i++ { ... } и for range пытается читать
из nil канала
```

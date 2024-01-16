Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
мы получим сообщение "error", потому что интерфейс Error() возвращает !nil значение,
и функция test() возвращает *customError который реализует интерфейс Error(),
поэтому мы попадем в блок err != nil и выведем сообщение об ошибке
```

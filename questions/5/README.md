# Какой размер у структуры `struct{}{}`?

```go
type Struct struct {
	value int
}

func main() {
	first := Struct{}
	second := struct{}{}

	fmt.Printf("First size: %v\n", reflect.TypeOf(first).Size())
	fmt.Printf("Second size: %v\n", reflect.TypeOf(second).Size())
}
```

- Результат работы программы - 8, 0;
- Размер сруктуры `struct{}{}` - 0 байт;
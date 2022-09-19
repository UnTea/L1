# Что выведет данная программа и почему?

```go
func update(p *int) {
    b := 2
    p = &b
}

func main() {
    var (
        a = 1
        p = &a
    )

    fmt.Println(*p)
    update(p)
    fmt.Println(*p)
}
```

- Вывод 1 1
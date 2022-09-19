# Что выведет данная программа и почему?

```go
func main() {
    n := 0

    if true {
        n := 1
        n++
    }

    fmt.Println(n)
}
```

- 0
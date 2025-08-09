# Variables and Data Types

- strings - default value is `""`

```go
 agency := "Fast track"
 fmt.Println("Welcome to", agency)
 name := "Jose Luis"
 fmt.Println(name)

```

- numbers - default value is `0`

```go
age := 33
tall := 1.77
fmt.Printf("My age is: %v\n", age)
fmt.Printf("My tall is: %v\n", tall)

```

- bool - default value is `false`

```go
isTrue := true
fmt.Println("Is true:", isTrue)
```

- more operations with strings [strings](https://pkg.go.dev/strings)

```go
str := "wadafa"
length := len(str)
fmt.Println(length)
str1 := "hello"
str2 := "Hello"
fmt.Println(strings.EqualFold(str1, str2))            //  equals
fmt.Println(strings.Index(str, "w"))                  // index
fmt.Println(str[strings.Index(str, "w"):4])           // slice
fmt.Println(str[:4])                                  // slice
fmt.Println(str[4:])                                  // slice
fmt.Println(strings.Replace(str, "fa", "-omaiga", 1)) // replace
str3 := "Go Programming"
fmt.Println(strings.ToUpper(str3))            // UpperCase
fmt.Println(strings.ToLower(str3))            // LoweCase
fmt.Println(strings.Contains(str3, "Go"))     // contains
fmt.Println(strings.Contains(str3, "wadafa")) // contains
```

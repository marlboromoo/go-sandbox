package main

import (
    "fmt"
)

type human struct {
    name string
}

type student struct {
    human
}

func fp(obj *student) {
    fmt.Printf("\nIn fp -- obj: type=%T, &obj=%p, obj=%p, *obj=%v\n", obj, &obj, obj, *obj)

    obj.name = "John Smith"
    fmt.Printf(`In fp >> obj.name = "John Smith"` + "\n")
    fmt.Printf("In fp -- obj: type=%T, &obj=%p, obj=%p, *obj=%v\n", obj, &obj, obj, *obj)

    (*obj).name = "J0hn Smith"
    fmt.Printf(`In fp >> (*obj).name = "J0hn Smith"` + "\n")
    fmt.Printf("In fp -- obj: type=%T, &obj=%p, obj=%p, *obj=%v\n", obj, &obj, obj, *obj)

    // panic here
    //obj = nil
    //fmt.Printf("In fp -- obj: type=%T, &obj=%p, obj=%p, *obj=%v\n", obj, &obj, obj, *obj)
}

func fv(obj student) {
    fmt.Printf("\nIn fv -- obj: type=%T, &obj=%p, obj=%v\n", obj, &obj, obj)

    obj.name = "John Smith"
    fmt.Printf(`In fv >> obj.name = "John Smith"` + "\n")
    fmt.Printf("In fv -- obj: type=%T, &obj=%p, obj=%v\n", obj, &obj, obj)

    (&obj).name = "John Rock"
    fmt.Printf(`In fv >> (&obj).name = "J0hn Rock"` + "\n")
    fmt.Printf("In fv -- obj: type=%T, &obj=%p, obj=%v\n", obj, &obj, obj)
}

func main() {
    john := student{human{name:"john"}}
    fmt.Printf(`In main >> john := student{human{name:"john"}}` + "\n")
    fmt.Printf("In main -- john: type=%T, &john=%p, john=%v\n", john, &john, john)
    fmt.Printf("In main >> john_ := &john\n")
    john_ := &john
    fmt.Printf("In main -- john_: type=%T, &john_=%p, john_=%p, *john_=%v\n", john_, &john_, john_, *john_)

    // function with pointer
    fmt.Printf("In main >> fp(john_)\n")
    fp(john_)
    fmt.Printf("In main -- john: type=%T, &john=%p, john=%v\n", john, &john, john)

    // function with value
    fmt.Printf("In main >> fv(john)\n")
    fv(john)
    fmt.Printf("In main -- john: type=%T, &john=%p, john=%v\n", john, &john, john)
}

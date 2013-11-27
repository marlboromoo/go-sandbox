// see: http://stackoverflow.com/questions/4938612/how-do-i-print-the-pointer-value-of-a-go-object-what-does-the-pointer-value-mea
package main

import "fmt"

func byval(q *int) {
    fmt.Printf("3. byval -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
    *q = 4143
    fmt.Printf("4. byval -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
    q = nil
}

func main() {
    i := int(42)
    fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
    p := &i
    fmt.Printf("2. main  -- p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
    byval(p)
    fmt.Printf("5. main  -- p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
    fmt.Printf("6. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
}

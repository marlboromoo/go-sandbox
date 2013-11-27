// see: http://golang.org/ref/spec#Calls
// A method call x.m() is valid if the method set of (the type of) x contains m 
// and the argument list can be assigned to the parameter list of m. If x is 
// addressable and &x's method set contains m, x.m() is shorthand for (&x).m()
package main

import (
    "fmt"
    "github.com/daviddengcn/go-colortext"
)

type human struct {
    name string
    fg ct.Color
    bg ct.Color
}

type student struct {
    human
}

type scut struct {
    student
}

type men interface {
    fight()
    punch()
    changeName(name string)
    sayName()
    getName() string
    readBook()
    getColor() ct.Color
}

type killer interface {
    sayName()
    getName() string
    getColor() ct.Color
    kill(m men)
}

func (h human) fight() {
    ct.ChangeColor(ct.Black, false, h.fg, false)
    fmt.Printf("%s's avatar(%T/%p) kick your ass!!\n", h.name, h, &h)
    ct.ResetColor()
}

func (h *human) punch() {
    ct.ChangeColor(h.fg, false, h.bg, false)
    fmt.Printf("%s(%T/%p) punch your face!!\n", h.name, h, h)
    ct.ResetColor()
}

func (h *human) changeName(name string) {
    name_ := h.name
    h.name = name
    ct.ChangeColor(h.fg, false, h.bg, false)
    fmt.Printf("%s(%T/%p) is now know as %s!\n", name_, h, h, h.name)
    ct.ResetColor()
}

func (h *human) sayName() {
    ct.ChangeColor(h.fg, false, h.bg, false)
    fmt.Printf("%s(%T/%p) say: my name is %s.\n", h.name, h, h, h.name)
    ct.ResetColor()
}

func (h *human) getName() string {
    return h.name
}

func (h *human) getColor() ct.Color {
    return h.fg
}

func (s *student) readBook() {
    ct.ChangeColor(s.fg, true, s.bg, false)
    fmt.Printf("%s(%T/%p) is reading the book!\n", s.name, s, s)
    ct.ResetColor()
}

func strip(m men) {
    ct.ChangeColor(m.getColor(), true, ct.None, true)
    fmt.Printf("*** %s(%T/%p) is naked ! ***\n", m.getName(), m, m)
    ct.ResetColor()
}


func (s *scut) kill(m men) {
    ct.ChangeColor(s.getColor(), true, ct.None, true)
    fmt.Printf("%s(%T/%p) kill %s(%T/%p)! !!!\n", s.getName(), s, s, m.getName(), m , m)
    ct.ResetColor()
}

func suicide(k killer) {
    ct.ChangeColor(k.getColor(), true, ct.None, true)
    fmt.Printf("*** %s(%T/%p) was suicide! ***\n", k.getName(), k, k)
    ct.ResetColor()
}

func main() {
    john := student{human{name:"John", fg:ct.Green, bg:ct.None}}
    david := student{human{name:"David", fg:ct.Cyan, bg:ct.None}}

    john_ := &john

    john.sayName()
    strip(&john)
    john.readBook()
    john.fight()
    john.punch()
    (&john).fight()
    (&john).punch()
    (&john).changeName("J0hn")
    john_.sayName()
    john_.fight()
    (*john_).punch()
    john_.readBook()

    mens := []men{&david, &john}
    for _, m := range mens {
        m.sayName()
        m.changeName(fmt.Sprintf("%s Smith", m.getName()))
        m.sayName()
        m.punch()
        m.fight()
        m.readBook()
        strip(m)
    }

    var rick killer
    rick = &scut{student{human{name:"Rick", fg:ct.Red, bg:ct.None}}}
    rick.sayName()
    rick.kill(&david)
    rick.kill(&john)
    suicide(rick)

}


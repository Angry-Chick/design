package main

import "fmt"

type person struct {
	name   string
	sex    string
	height string
	weight string
}

func (p *person) setPersonInfo(name, sex, height, weight string) {
	if p == nil {
		return
	}
	p.name = name
	p.height = height
	p.sex = sex
	p.weight = weight
}

func (p *person) setName(name string) {
	if p == nil {
		return
	}
	p.name = name
}

func (p *person) clone() *person {
	if p == nil {
		return nil
	}
	newObj := *p
	return &newObj
}

func main() {
	p := &person{}
	p.setPersonInfo("张三", "男", "175", "65")
	fmt.Println(p)
	p1 := p.clone()
	p1.setName("李四")
	p2 := p.clone()
	p2.setName("王五")
	fmt.Println(p, p1, p2)
}

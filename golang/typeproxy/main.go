package main

import (
	"fmt"
)

type MyType interface {
	Name() string
	SetName(name string) error
}

type MyTypeRepository interface {
	Create(MyType) (MyStorableType, error)
	Read(name string) (MyStorableType, error)
	Update(MyType) error
	Delete(MyType) error
}

type MyStorableType interface {
	MyType
	Update() error
	Delete() error
}

type MyTypeStruct struct {
	name string
}

type MyTypeProxy struct {
	MyTypeRepository
	MyType
	id uint
}

type Repo struct {
	id uint
}

func (this *MyTypeStruct) Name() string {
	return this.name
}

func (this *MyTypeStruct) SetName(name string) error {
	this.name = name
	return nil
}

func (this *MyTypeProxy) SetName(name string) error {
	return this.MyType.SetName(name)
}

func (this *MyTypeProxy) Update() error {
	return this.MyTypeRepository.Update(this)
}

func (this *MyTypeProxy) Delete() error {
	return this.MyTypeRepository.Delete(this)
}

func (r *Repo) nextID() uint {
	r.id++
	return r.id
}

func (r *Repo) Create(t MyType) (MyStorableType, error) {
	fmt.Printf("%T Create(%s)\n", r, t.Name())
	return &MyTypeProxy{
		MyType: t,
		MyTypeRepository: r,
		id: r.nextID(),
	}, nil
}

func (r *Repo) Read(name string) (MyStorableType, error) {
	mt := &MyTypeStruct{name: name}
	fmt.Printf("%T Read(%s)\n", r, mt.Name())
	return &MyTypeProxy{
		MyType: mt,
		MyTypeRepository: r,
	}, nil
}

func (this *Repo) Delete(m MyType) error {
	fmt.Printf("%T Delete(%s) with id %d\n", this, m.Name(), m.(*MyTypeProxy).id)
	return nil
}

func (this *Repo) Update(m MyType) error {
	fmt.Printf("%T Update(%s) with id %d\n", this, m.Name(), m.(*MyTypeProxy).id)
	return nil
}

func main() {
	var r MyTypeRepository

	r = &Repo{}

	t, _ := r.Create(&MyTypeStruct{name: "John"})
	t.SetName("john")
	t.Update()
	t.Delete()

	t, _ = r.Create(&MyTypeStruct{name: "Jane"})
	t.SetName("jane")
	t.Update()
	t.Delete()
}

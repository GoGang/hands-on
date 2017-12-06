package main

import "fmt"

type Talk struct {
	Name        string
	Description string
	Properties  map[string]interface{}
}

func (t Talk) GetProperty(name string) interface{} { // affectation d'une fonction à une structure
	return t.Properties[name]
}

func (t *Talk) AddProperty(name string, value interface{}) { // affectation d'une fonction à un pointeur de structure
	t.Properties[name] = value
}

func (t Talk) SetNameBogus(name string) { // la fonction utilise une copie de la structure, pas de modification possible
	t.Name = name
}

func (t *Talk) SetName(name string) { // la fonction utilise une copie de la structure, pas de modification possible
	t.Name = name
}

func NewTalk(name, description string) *Talk { // constructeur
	return &Talk{ // & = récupération de l'adresse du pointeur
		Name:        name,
		Description: description,
		Properties:  make(map[string]interface{}),
	}
}

func main() {
	talk := NewTalk("Golang", "A quick tour of golang")
	talk.AddProperty("foo", "bar")
	talk.AddProperty("num", 1)

	fmt.Printf("talk: %+v \n", talk)

	fmt.Println("foo : ", talk.GetProperty("foo"))
	fmt.Println("num : ", talk.GetProperty("num"))

	talkNoPtr := *talk // on récupère le Talk via son pointeur

	fmt.Println("foo noptr : ", talkNoPtr.GetProperty("foo"))
	fmt.Println("num noptr : ", talkNoPtr.GetProperty("num"))

	talkNoPtr.SetNameBogus("name bogus")
	fmt.Println("name bogus : ", talkNoPtr.Name)

	talk.SetNameBogus("name bogus") // une méthode attachée à Talk est appelable sur *Talk (déréférencement automatique)
	fmt.Println("name bogus again : ", talk.Name)

	talk.SetName("name not bogus")
	fmt.Println("name not bogus : ", talk.Name)
}

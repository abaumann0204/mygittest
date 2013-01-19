package controllers

import "github.com/robfig/revel"

type Application struct {
	*rev.Controller
}

func (c Application) Index() rev.Result {
	return c.Render()
}

// Id Feld sollte ausgeliefert werden !!!
type TPerson struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDay  string `json:"birth_day"`
}

type TPersonResponse struct {
	Person TPerson `json:"person"`
}

type TPeople struct {
	People []TPerson `json:"people"`
}

var testPerson TPersonResponse = TPersonResponse{
	Person: TPerson{
		Id:        0,
		FirstName: "Andreas",
		LastName:  "Baumann",
		BirthDay:  "02.04.1960",
	},
}

var testPersonen TPeople = TPeople{
	People: []TPerson{
		TPerson{
			Id:        1,
			FirstName: "Andreas",
			LastName:  "Baumann",
			BirthDay:  "02.04.1960",
		},
		TPerson{
			Id:        2,
			FirstName: "Kerstin",
			LastName:  "Baumann",
			BirthDay:  "10.07.1963",
		},
		TPerson{
			Id:        3,
			FirstName: "Sarah",
			LastName:  "Baumann",
			BirthDay:  "11.10.1990",
		},
		TPerson{
			Id:        4,
			FirstName: "Matthias",
			LastName:  "Baumann",
			BirthDay:  "14.12.1966",
		},
		TPerson{
			Id:        5,
			FirstName: "Ute",
			LastName:  "Petzel",
			BirthDay:  "24.08.1961",
		},
	},
}

func (c Application) SelectedPerson() rev.Result {
	response := c.RenderJson(testPerson)
	rev.INFO.Println("JSon Response:", response)
	return response
}

func (c Application) People() rev.Result {
	return c.RenderJson(testPersonen)
}

func (c Application) Person(id int) rev.Result {
	var p TPerson = testPersonen.People[id]
	return c.RenderJson(p)
}

package main

import (
	"log"
	"os"
	"text/template"
)

type item struct{
	Name, Desc string
	Price float64
}

type meal struct{
	Meal string
	Item []item
}

type menu []meal

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main(){
	m := menu{
		meal{
			Meal: "Dinner",
			Item: []item{
				item{
					Name: "Garbage",
					Desc: "Foodgrade 0",
					Price: 8.69,
				},
				item{
					Name: "Dick",
					Desc: "like sausage",
					Price: 10.01,
				},
			},
		},
		meal{
			Meal: "Breakfast",
			Item: []item{
				item{
					Name: "Boobs",
					Desc: "twin mountains",
					Price: 46.21,
				},
				item{
					Name: "Pussy",
					Desc: "I like it",
					Price: 60.00,
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, m)
	if err != nil{
		log.Fatalln(err)
	}
}
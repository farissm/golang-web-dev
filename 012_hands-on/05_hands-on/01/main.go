package main

import (
	"log"
	"os"
	"text/template"
)

type item struct{
	Name, Desc, Meal string
	Price float64
}

type items []item

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	m := items{
		item{
			Name: "Beef",
			Desc: "Juicy",
			Meal: "Dinner",
			Price: 10.00,
		},
		item{
			Name: "Garbage",
			Desc: "So yummy",
			Meal: "All time",
			Price: 0.00,
		},
	}
	err := tpl.Execute(os.Stdout, m)
	if err != nil{
		log.Fatalln(err)
	}
}
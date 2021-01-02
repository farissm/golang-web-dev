package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct{
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	h := hotels{
		hotel{
			Name: "Kempinski",
			Address: "Bundaran HI",
			City: "Jakarta",
			Zip: "15232",
			Region: "Java",
		},
		hotel{
			Name: "Swissbel",
			Address: "Dago",
			City: "Bandung",
			Zip: "15232",
			Region: "West Java",
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil{
		log.Fatalln(err)
	}
}
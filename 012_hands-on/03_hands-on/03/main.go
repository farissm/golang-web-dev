package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct{
	Name, Address, City, Zip, Region string
}

type region struct{
	Region string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	h := Regions{
		region{
			Region: "Jkt",
			Hotels: []hotel{
				hotel{
					Name: "Kempinski",
					Address: "Bundaran HI",
					City: "Jakarta",
					Zip: "12345",
					Region: "Jakarta",
				},
				hotel{
					Name: "Aston",
					Address: "Kuningan",
					City: "Jakarta",
					Zip: "54321",
					Region: "Jakarta",
				},
			},
		},
		region{
			Region: "Bdo",
			Hotels: []hotel{
				hotel{
					Name: "Padma",
					Address: "Setiabudi",
					City: "Bandung",
					Zip: "09876",
					Region: "Bandung",
				},
				hotel{
					Name: "Swissbel",
					Address: "Dago",
					City: "Bandung",
					Zip: "67890",
					Region: "Bandung",
				},
			},
		},
	} 

	err := tpl.Execute(os.Stdout, h)
	if err != nil{
		log.Fatalln(err)
	}
}
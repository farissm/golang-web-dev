package main

import(
	"log"
	"os"
	"text/template"
)

type hotel struct{
	Name, Address, City, Zip string
}

type region struct{
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	h := region{
		Region: "Jkt",
		Hotels: []hotel{
			hotel{
				Name: "Kempinski",
				Address: "Jakpus",
				City: "Jakarta",
				Zip: "12301",
			},
			hotel{
				Name: "Ritz Carlton",
				Address: "Jaksel",
				City: "Jakarta",
				Zip: "12382",
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil{
		log.Fatalln(err)
	}
}
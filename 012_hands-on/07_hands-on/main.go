package main

import (
	"os"
	"log"
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

type restaurant struct{
	Name string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	m := restaurants{
		restaurant{
			Name: "Federicos",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:    "Oatmeal",
							Desc: "yum yum",
							Price:   4.95,
						},
						item{
							Name:    "Cheerios",
							Desc: "American eating food traditional now",
							Price:   3.95,
						},
						item{
							Name:    "Juice Orange",
							Desc: "Delicious drinking in throat squeezed fresh",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:    "Hamburger",
							Desc: "Delicous good eating for you",
							Price:   6.95,
						},
						item{
							Name:    "Cheese Melted Sandwich",
							Desc: "Make cheese bread melt grease hot",
							Price:   3.95,
						},
						item{
							Name:    "French Fries",
							Desc: "French eat potatoe fingers",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:    "Pasta Bolognese",
							Desc: "From Italy delicious eating",
							Price:   7.95,
						},
						item{
							Name:    "Steak",
							Desc: "Dead cow grilled bloody",
							Price:   13.95,
						},
						item{
							Name:    "Bistro Potatoe",
							Desc: "Bistro bar wood American bacon",
							Price:   6.95,
						},
					},
				},
			},
		},
		restaurant{
			Name: "Domenicos",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:    "Oatmeal",
							Desc: "yum yum",
							Price:   4.95,
						},
						item{
							Name:    "Cheerios",
							Desc: "American eating food traditional now",
							Price:   3.95,
						},
						item{
							Name:    "Juice Orange",
							Desc: "Delicious drinking in throat squeezed fresh",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:    "Hamburger",
							Desc: "Delicous good eating for you",
							Price:   6.95,
						},
						item{
							Name:    "Cheese Melted Sandwich",
							Desc: "Make cheese bread melt grease hot",
							Price:   3.95,
						},
						item{
							Name:    "French Fries",
							Desc: "French eat potatoe fingers",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:    "Pasta Bolognese",
							Desc: "From Italy delicious eating",
							Price:   7.95,
						},
						item{
							Name:    "Steak",
							Desc: "Dead cow grilled bloody",
							Price:   13.95,
						},
						item{
							Name:    "Bistro Potatoe",
							Desc: "Bistro bar wood American bacon",
							Price:   6.95,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil{
		log.Fatalln(err)
	}
}
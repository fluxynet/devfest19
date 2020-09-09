package main

import (
	"net/http"
	"text/template"
)

var (
	tmplHello = template.Must(template.ParseFiles("templates/index.html"))
)

type planet struct {
	Name  string
	Photo string
}

// templateHandler showing templating in golang
func templateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	planets := struct {
		Planets []planet
	}{
		Planets: []planet{
			planet{
				Name:  "Mercury",
				Photo: "https://www.popsci.com/resizer/h_Wb0PwWbHG8pyy-57RGJ563bqo=/1034x1034/arc-anglerfish-arc2-prod-bonnier.s3.amazonaws.com/public/RZODSOHKA6O4VMNPPUW7GUTUBE.jpg",
			},
			planet{
				Name:  "Venus",
				Photo: "https://www.popsci.com/resizer/WsYKweCWKvQec1WYoOSuYxWXyC0=/525x525/arc-anglerfish-arc2-prod-bonnier.s3.amazonaws.com/public/LP5TMWXTF6VV6VQ6YDQGZ3YLQQ.jpg",
			},
			planet{
				Name:  "Earth",
				Photo: "https://www.popsci.com/resizer/JlRmCkKmhQ13t_wvYMA3q5NbcQg=/1034x1034/arc-anglerfish-arc2-prod-bonnier.s3.amazonaws.com/public/57LV7D7YHYJ2BDWMUKBC3J5WG4.jpg",
			},
			planet{
				Name:  "Mars",
				Photo: "https://www.popsci.com/resizer/NYeeNsFatVEMNPneaQoRJCHH_3A=/525x525/arc-anglerfish-arc2-prod-bonnier.s3.amazonaws.com/public/XYNJU7VV7ISTJ4HJLUR5ERUZOE.jpg",
			},
			planet{
				Name:  "Jupiter",
				Photo: "https://www.popsci.com/resizer/QDsh5g2GJttuJ2VX7gLRb4FJ12o=/525x656/arc-anglerfish-arc2-prod-bonnier.s3.amazonaws.com/public/THFNFWTVBVSMJVYCP5BACG3YFQ.jpg",
			},
			planet{
				Name:  "Uranus",
				Photo: "https://www.popsci.com/resizer/vXGP9ACBgERSmLQ-Y7cipa4G7S0=/525x525/arc-anglerfish-arc2-prod-bonnier.s3.amazonaws.com/public/2JK6SFTHA6WHV56FGHXNWWNDSE.jpg",
			},
			planet{
				Name:  "Neptune",
				Photo: "https://www.popsci.com/resizer/qOippOhzEu10l9LaLSU8dnAWRRw=/525x525/arc-anglerfish-arc2-prod-bonnier.s3.amazonaws.com/public/3BR7HFELGONI3J4CB4USPR5GWI.jpg",
			},
		},
	}

	tmplHello.Execute(w, planets)
}

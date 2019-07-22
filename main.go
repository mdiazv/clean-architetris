package main

import (
	"github.com/mdiazv/clean-architetris/domain"
	"github.com/mdiazv/clean-architetris/interfaces"

	"log"
	"net/http"
)

func main() {
	world := domain.MakeTetrisWorld()

	http.HandleFunc("/drop", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Dropping a piece")
		world.Drop()
	})
	http.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Rendering the world")
		out := interfaces.MakeHttpOutputAdapter(w, r)
		out.Render(world)
	})

	log.Println("Starting web server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"log"
	"net/http"
	"os"
	"pemesanan/config"
	"pemesanan/controllers"
	"pemesanan/controllers/auth"
	"pemesanan/repository"
	"pemesanan/usecase"

	"github.com/go-chi/chi"
)

func main() {
	connection := config.Connect()
	repository := repository.NewRepo(connection)
	usecase := usecase.NewUsecase(repository)
	controller := controllers.NewController(usecase)
	routes := chi.NewRouter()

	routes.Group(func(r chi.Router) {
		r.Get("/api/k1/kategori", controller.GetDataKategori)
		r.Post("/api/k2/kategori", controllers.PostKategoris)
		r.Get("/api/k2/kereta", controller.GetDataKereta)
		r.Get("/api/k3/detail", controller.GetDataDetail)
		r.Get("/api/k4/detail/{id}", controller.GetDataDetailById)
		r.Post("/api/u1/user", auth.Register)
		r.Post("/api/k4/detailkereta", controllers.DetailKereta)
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), routes); err != nil {
		log.Fatal(err)
	}
	log.Println("Server Running on port 8002")
}

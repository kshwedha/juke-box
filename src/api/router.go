package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Route(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "POST, GET, PATCH, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           3600,
	}))

	api := app.Group("api")
	api.Post("/album/add", AddAlbum)
	api.Patch("/album/update/:name", UpdateAlbum)
	api.Get("/album/all", RetrieveAlbum)
	api.Get("/album/retrieve/:album", RetrieveAlbumArtists)

	api.Post("/artist/add", AddArtist)
	api.Patch("/artist/update/:name", UpdateArtist)
	api.Get("/artist/retrieve/:artist", RetrieveArtistAlbums)
	api.Get("/artist/all", RetrieveArtist)

	api.Post("/track/add", TunePlaylist)
}

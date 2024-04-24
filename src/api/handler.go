package api

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/kshwedha/juke-box/src/function"
)

type Album struct {
	Name        string `json:"name"`
	Date        string `json:"date"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type Artist struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Playlist struct {
	Song   string `json:"name"`
	Singer string `json:"singer"`
}

func AddAlbum(c *fiber.Ctx) error {
	var album Album
	if err := c.BodyParser(&album); err != nil {
		fmt.Println(err)
		return err
	}
	err := function.Album.AddAlbumF(function.Album(album))
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Album added successfully",
	})
}

func UpdateAlbum(c *fiber.Ctx) error {
	name := c.Params("name")
	decoded_name, err := url.QueryUnescape(name)
	if err != nil {
		return fmt.Errorf("error decoding URL: %v", err)
	}
	DoesExist := function.DoesAlbumExists(decoded_name)

	if !DoesExist {
		return c.JSON(fiber.Map{
			"error": "Album does not exist",
		})
	}

	var album Album
	if err := c.BodyParser(&album); err != nil {
		fmt.Println(err)
		return err
	}
	if album.Date == "" && album.Genre == "" && album.Description == "" && album.Price == "" {
		return c.JSON(fiber.Map{
			"error": "No fields to update",
		})
	}

	album.Name = decoded_name

	err = function.Album.UpdateAlbumF(function.Album(album))
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Album updated successfully",
	})

}

func AddArtist(c *fiber.Ctx) error {
	var artist Artist
	if err := c.BodyParser(&artist); err != nil {
		fmt.Println(err)
		return err
	}
	err := function.Artist.AddArtistF(function.Artist(artist))
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Musician added successfully",
	})
}

func UpdateArtist(c *fiber.Ctx) error {
	name := c.Params("name")
	name, err := url.QueryUnescape(name)
	if err != nil {
		return fmt.Errorf("error decoding URL: %v", err)
	}
	DoesExist := function.DoesMusicianExists(name)
	if !DoesExist {
		return c.JSON(fiber.Map{
			"error": "Musician does not exist",
		})
	}

	var artist Artist
	c.BodyParser(&artist)
	artist.Name = name

	err = function.Artist.UpdateArtistF(function.Artist(artist))
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Musician updated successfully",
	})
}

func RetrieveAlbum(c *fiber.Ctx) error {
	data, err := function.RetrieveAlbum()
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"count":  len(data),
		"data":   data,
	})
}

func RetrieveArtistAlbums(c *fiber.Ctx) error {
	name := c.Params("artist")
	name, err := url.QueryUnescape(name)
	if err != nil {
		return fmt.Errorf("error decoding URL: %v", err)
	}
	DoesExist := function.DoesMusicianExists(name)
	if !DoesExist {
		return c.JSON(fiber.Map{
			"error": "Musician does not exist",
		})
	}

	data, err := function.RetrieveArtistAlbums(name)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"count":  len(data),
		"data":   data,
	})
}

func RetrieveAlbumArtists(c *fiber.Ctx) error {
	name := c.Params("album")
	name, err := url.QueryUnescape(name)
	if err != nil {
		return fmt.Errorf("error decoding URL: %v", err)
	}

	DoesExist := function.DoesAlbumExists(name)
	if !DoesExist {
		return c.JSON(fiber.Map{
			"error": "Album does not exist",
		})
	}
	data, err := function.RetrieveAlbumArtists(name)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"count":  len(data),
		"data":   data,
	})
}

func TunePlaylist(c *fiber.Ctx) error {
	var playlist Playlist
	if err := c.BodyParser(&playlist); err != nil {
		return err
	}
	err := function.Playlist.Tune(function.Playlist(playlist))
	if err != nil {
		c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Playlist created successfully",
	})
}

func RetrieveArtist(c *fiber.Ctx) error {
	data, err := function.RetrieveArtist()
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"count":  len(data),
		"data":   data,
	})
}

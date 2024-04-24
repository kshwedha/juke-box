package function

import (
	"fmt"
	"strconv"

	"github.com/kshwedha/juke-box/src/common/driver"
)

type Album struct {
	Name        string
	Date        string
	Genre       string
	Description string
	Price       string
}

type Artist struct {
	Name string
	Type string
}

type Playlist struct {
	Song   string
	Singer string
}

func DoesAlbumExists(name string) bool {
	db, err := driver.InitDB()
	if err != nil {
		return false
	}
	defer db.Close()

	query := fmt.Sprintf("select * from album where name = '%s';", name)
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()

	return rows.Next()
}

func DoesMusicianExists(name string) bool {
	db, err := driver.InitDB()
	if err != nil {
		return false
	}
	defer db.Close()

	query := fmt.Sprintf("select * from musician where name = '%s';", name)
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()

	return rows.Next()
}

func (album Album) AddAlbumF() error {
	// name string, date string, genre string, price float32, description string
	db, err := driver.InitDB()
	if err != nil {
		return fmt.Errorf("!! DB connection error")
	}
	defer db.Close()

	price, err := strconv.ParseFloat(album.Price, 64)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("insert into album (name, release_date, genre, price, description) values ('%s', '%s', '%s', '%f', '%s');", album.Name, album.Date, album.Genre, price, album.Description)

	results, err := driver.ExecPsqlResult(db, query)
	if err != nil {
		return err
	}
	if results == 1 {
		return nil
	}
	return fmt.Errorf("!! could not register album")
}

func (album Album) UpdateAlbumF() error {
	db, err := driver.InitDB()
	if err != nil {
		return fmt.Errorf("!! DB connection error")
	}
	defer db.Close()

	price, err := strconv.ParseFloat(album.Price, 64)
	if err != nil {
		return err
	}

	query := "UPDATE album SET"
	if album.Date != "" {
		query += fmt.Sprintf(" release_date = '%s'", album.Date)
	}
	if album.Genre != "" {
		query += fmt.Sprintf(", genre = '%s'", album.Genre)
	}
	if album.Description != "" {
		query += fmt.Sprintf(", description = '%s'", album.Description)
	}
	if album.Price != "" {
		query += fmt.Sprintf(", price = '%f'", price)
	}
	query += fmt.Sprintf(" WHERE name = '%s';", album.Name)

	results, err := driver.ExecPsqlResult(db, query)
	if err != nil {
		return err
	}
	if results == 1 {
		return nil
	}
	return fmt.Errorf("!! could not update album")
}

func (artist Artist) AddArtistF() error {
	db, err := driver.InitDB()
	if err != nil {
		return fmt.Errorf("!! DB connection error")
	}
	defer db.Close()

	query := fmt.Sprintf("insert into musician (name, musician_type) values ('%s', '%s');", artist.Name, artist.Type)

	results, err := driver.ExecPsqlResult(db, query)
	if err != nil {
		return err
	}
	if results == 1 {
		return nil
	}
	return fmt.Errorf("!! could not register artist")

}

func (artist Artist) UpdateArtistF() error {
	db, err := driver.InitDB()
	if err != nil {
		return fmt.Errorf("!! DB connection error")
	}
	defer db.Close()

	query := fmt.Sprintf("UPDATE musician SET musician_type = '%s' where name='%s'", artist.Type, artist.Name)
	results, err := driver.ExecPsqlResult(db, query)
	if err != nil {
		return err
	}
	if results == 1 {
		return nil
	}
	return fmt.Errorf("!! could not update musician")
}

func RetrieveAlbum() ([]Album, error) {
	db, err := driver.InitDB()
	if err != nil {
		return []Album{}, fmt.Errorf("!! DB connection error")
	}
	defer db.Close()
	query := "select name, release_date, genre, price from album;"
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()
	var albums []Album
	for rows.Next() {
		var name string
		var release_date string
		var genre string
		var price float32
		if err := rows.Scan(&name, &release_date, &genre, &price); err != nil {
			return []Album{}, err
		}
		albums = append(albums, Album{Name: name, Date: release_date, Genre: genre, Price: fmt.Sprintf("%f", price)})
	}
	return albums, nil
}

func getArtistId(artist string) (int, error) {
	db, err := driver.InitDB()
	if err != nil {
		return 0, fmt.Errorf("!! DB connection error")
	}
	defer db.Close()
	query := "select id from musician where name = '" + artist + "';"
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()
	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
		return id, nil
	}
	return -1, fmt.Errorf("!! could not find artist")
}

func getAlbumId(album string) (int, error) {
	db, err := driver.InitDB()
	if err != nil {
		return 0, fmt.Errorf("!! DB connection error")
	}
	defer db.Close()
	query := "select id from album where name = '" + album + "';"
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()
	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
		return id, nil
	}
	return -1, fmt.Errorf("!! could not find album")
}

func getAlbumName(albumId int) (string, error) {
	db, err := driver.InitDB()
	if err != nil {
		return "", fmt.Errorf("!! DB connection error")
	}
	defer db.Close()
	query := "select name from album where id = " + strconv.Itoa(albumId) + ";"
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return "", err
		}
		return name, nil
	}
	return "", fmt.Errorf("!! could not find album")
}

func getArtistName(artistId int) (string, error) {
	db, err := driver.InitDB()
	if err != nil {
		return "", fmt.Errorf("!! DB connection error")
	}
	defer db.Close()
	query := "select name from musician where id = " + strconv.Itoa(artistId) + ";"
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return "", err
		}
		return name, nil
	}
	return "", fmt.Errorf("!! could not find artist")
}

func RetrieveArtistAlbums(artist string) ([]string, error) {
	db, err := driver.InitDB()
	if err != nil {
		return nil, fmt.Errorf("!! DB connection error")
	}
	defer db.Close()

	artistId, err := getArtistId(artist)
	if err != nil {
		return nil, err
	}
	query := "select album_id from playlist where musician_id = " + strconv.Itoa(artistId) + ";"
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()
	var albums []string

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		name, err = getAlbumName(id)
		if err == nil {
			albums = append(albums, name)
		}
	}
	return albums, nil
}

func RetrieveAlbumArtists(album string) ([]string, error) {
	db, err := driver.InitDB()
	if err != nil {
		return nil, fmt.Errorf("!! DB connection error")
	}
	defer db.Close()

	album_id, err := getAlbumId(album)
	if err != nil {
		return nil, err
	}
	query := "select musician_id from playlist where album_id = " + strconv.Itoa(album_id) + ";"
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()

	var artists []string
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		name, err = getArtistName(id)
		if err == nil {
			artists = append(artists, name)
		}
	}
	return artists, nil
}

func (Playlist Playlist) Tune() error {
	db, err := driver.InitDB()
	if err != nil {
		return fmt.Errorf("!! DB connection error")
	}
	defer db.Close()

	if Playlist.Singer == "" {
		return fmt.Errorf("!! Playlist Singer is empty")
	}
	if Playlist.Song == "" {
		return fmt.Errorf("!! Playlist Song is empty")
	}

	AlbumExists := DoesAlbumExists(Playlist.Song)
	ArtistExists := DoesMusicianExists(Playlist.Singer)
	if AlbumExists && ArtistExists {
		artist_id, err := getArtistId(Playlist.Singer)
		if err != nil {
			return err
		}
		album_id, err := getAlbumId(Playlist.Song)
		if err != nil {
			return err
		}
		query := "insert into playlist (musician_id, album_id) values (" + strconv.Itoa(artist_id) + ", " + strconv.Itoa(album_id) + ");"
		results, err := driver.ExecPsqlResult(db, query)
		if err != nil {
			return err
		}
		if results == 1 {
			return nil
		}
	}
	return fmt.Errorf("!! could not register playlist")
}

func RetrieveArtist() ([]Artist, error) {
	db, err := driver.InitDB()
	if err != nil {
		return []Artist{}, fmt.Errorf("!! DB connection error")
	}
	defer db.Close()

	query := "select name, musician_type from musician;"
	rows := driver.ExecPsqlRows(db, query)
	defer rows.Close()
	var artists []Artist
	for rows.Next() {
		var name string
		var type_ string
		if err := rows.Scan(&name, &type_); err != nil {
			return []Artist{}, err
		}
		artists = append(artists, Artist{Name: name, Type: type_})
	}
	return artists, nil
}

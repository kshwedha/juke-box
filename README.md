
# Juke Box - Go

***Juke-Box***

API's

```POST: /album/add

param:

name, description, genre, release_date, price

PATCH: /album/update/:name

param:

description, genre, release_date, price

GET: /album/all

GET: /album/retrieve/:album

POST: /artist/add

param:

name, type

PATCH: /artist/update/:name

param:

type

GET: /artist/retrieve/:artist

GET: /artist/all

POST: /track/add

param:
song_name, singer_name
```
## Authors

- [@kshwedha](https://www.github.com/kshwedha)


## Deployment

To deploy this project run

```bash
  go run main.go
```


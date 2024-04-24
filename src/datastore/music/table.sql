CREATE table album(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL CHECK (LENGTH(name) >= 5),
    release_date DATE NOT NULL,
    genre VARCHAR(255),
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 100 AND price <= 1000),
    description TEXT
);

CREATE TABLE musician (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL CHECK (LENGTH(name) >= 3),
    musician_type VARCHAR(255)
);

CREATE TABLE playlist (
    album_id int,
    musician_id int,
    CONSTRAINT fk_album_id FOREIGN KEY (album_id) REFERENCES album(id) ON DELETE CASCADE,
    CONSTRAINT fk_musician_id FOREIGN KEY (musician_id) REFERENCES musician(id) on DELETE CASCADE
);

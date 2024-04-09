CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    surname VARCHAR(100),
    patronymic VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS car (
    id SERIAL PRIMARY KEY,
    regNum VARCHAR(100),
    mark VARCHAR(100),
    model VARCHAR(100),
    year INTEGER,
    people_id INT NOT NULL REFERENCES people (id)
);
CREATE TABLE urls (
    url_id serial PRIMARY KEY, 
    url VARCHAR(255) NOT NULL,
    expireat VARCHAR(255) NOT NULL
);
CREATE TABLE urls (
    url_id serial PRIMARY KEY, 
    link VARCHAR(255) NOT NULL,
    expireat VARCHAR(255) NOT NULL
);
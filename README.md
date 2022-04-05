<h1 align="center">Shorten_url_API</h1>
![image](https://github.com/dinoslow/Shorten_url_API/blob/main/example.gif)

## Description
This project will shorten the given url and redirect the shorten url to original url

## Usage

## How API works
GET Method -> Redirect shorten url to the original url
```go=
func Get(w http.ResponseWriter, r *http.Request) {
    
    1. record time when GET Method is called
    2. get param from url
    3. using param to search original url in db
    4. redirect to the original url while exist or not expired
    
}
```
POST Method -> convert the given url to a shorter url
```go=
func Post(w http.ResponseWriter, r *http.Request) {
    
    1. get params from request body
    2. insert url, expireDate to db and return url index
    3. generate JSON file with shorten url
    4. write response
    
}
```
## Other functions
```
```

## Database
```sql=
CREATE TABLE urls (
    url_id serial PRIMARY KEY, 
    link VARCHAR(255) NOT NULL,
    expireat VARCHAR(255) NOT NULL
);
```
在這個作業裡面我使用了postgresSQL來作為資料庫。因為在大一下我修的一門課 - 軟體實驗設計(Software studio)就有學到postgresSQL，對我來說比MongoDB熟悉。

## Third Party Library
#### *"gorilla/mux"*

在這個作業裡面我用了mux做為router，因為mux在handleFunc傳入的參數和單純使用http做為route的handleFunc是一樣的，我認為比較好理解，而且寫的code也更直觀了。


#### *"lib/pq"*
我使用了這個library來連接postgresSQL資料庫。

#### *"stretchr/testify/assert"*
我使用了這個library來用作test.go裡面的斷言，使用上相對簡潔。

## Project setup (docker)
```
docker-compose up
```
<h1 align="center">Shorten_url_API</h1>

## Description
#### This project will shorten the given url and redirect the shorten url to the given original url.

## Usage
![image](https://github.com/dinoslow/Shorten_url_API/blob/main/example.gif)

Post a url with expire date.

```
curl -X POST -H "Content-Type:application/json" http://localhost:3001/api/v1/urls -d '{
   "url": "http://www.google.com",
    "expireAt": "2023-01-01T00:00:00Z"
}'
```
server will return a shorten url.
```
{
    "id": "3",
    "shortUrl": "http://localhost:3001/3"
}
```
Then, use the shorten url to go to the original website.
```
curl -L -X GET http://localhost:3001/3
```
Redirect to google.com

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
## Deploy on AWS
To implement reverse proxy. I deployed my project to AWS and used Nginx as reverse proxy.
*Post with req.Body*
```
http://ec2-52-197-102-90.ap-northeast-1.compute.amazonaws.com/api/v1/urls

    with 

{
    "url":"http://www.youtube.com",
    "expireAt":"2023-01-01T00:00:00Z"
}
```
http://ec2-52-197-102-90.ap-northeast-1.compute.amazonaws.com/1

---
http://ec2-52-197-102-90.ap-northeast-1.compute.amazonaws.com/2

---
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

## More feature
1. 實作cache - 用go-cache這個library或是Redis將url的id和redirect方向暫存，來增加效率。
2. 將function和struct分類變得更有架構 - 例如分成router, model, util等等。
## Project setup (docker)
```
docker-compose up
```
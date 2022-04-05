# shorten_url_API

## Description
This project will shorten the given url to a shorter url.

## Usage

## How API works
GET Method -> Redirect shorten url to the origin website
```go=
func Get(w http.ResponseWriter, r *http.Request) {
    
    1. record time when GET Method is called
    2. get param from url
    3. using param to search origin url in db
    4. redirect to the origin website
    
}
```
POST Method -> convert the given url to a shorter one 
```go=
func Post()
```
## Database

## Third Party Library

## Project setup
```
docker-compose up
```

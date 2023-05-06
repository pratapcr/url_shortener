# url_shortener

To start the service, 
1)docker-compose up -d  in your terminal


Service starts on port http://localhost:1234

APIS

POST - to insert url and get shorturl
GET (using short url) - to access long url with short link


2)Use postman to send urls 
   /POST http://localhost:1234/
body : formdata 
 url : "long url......"

response will be short url 

Redirection is not handled

Only response is handled.

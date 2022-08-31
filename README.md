# todont

API for things you absolutely don't want to do

## Introduction

Todont is the API that powers an imaginary application where you can add the things you don't want to do.

### Routes

| Route           	| Method 	| Require Auth 	| Description                                               	|
|-----------------	|--------	|--------------	|-----------------------------------------------------------	|
 /               	| GET    	| FALSE        	| Root route. Usually used to check if server is online.    	|
 /health         	| GET    	| FALSE        	| Display heatlh of the API service.                        	|
 /ping           	| GET    	| FALSE        	| Ping is used to check if the server is online.            	|
 /v1/auth/signup 	| POST   	| FALSE        	| Used to create a new account                              	|
 /v1/auth/login  	| POST   	| FALSE        	| used to authenticate a user                               	|
 /v1/me          	| GET    	| TRUE         	| used to fetch currently logged in user information        	|
 /v1/me          	| DELETE 	| TRUE         	| used to delete the currently logged in user               	|
 /v1/me          	| PATCH  	| TRUE         	| used to update user information                           	|
 /v1/item        	| GET    	| FALSE        	| used to retrieve all the items in the store               	|
 /v1/item        	| POST   	| TRUE         	| used to create an item in the store                       	|
 /v1/item/:id    	| GET    	| TRUE         	| used to retrieve a single item in the store               	|
 /v1/item/:id    	| PATCH  	| TRUE         	| used to update an item, only a creator can update an item 	|
 /v1/item/:id    	| DELETE 	| TRUE         	| used to delete an item, only a creator can delete an item 	|

#### Persistence

Information is to be stored in a PostgreSQL database.

##### Router

I opted for [go-chi](https://github.com/go-chi/chi/) instead of [gorilla/mux](https://github.com/gorilla/mux) because of the following reasons:

* `mux` & `go-chi` both have similar cons to their usage

* `mux` contains some advanced features (such as custom routing rules, host-based routing or route 'reversing') which I do not have a need for

* `go-chi` offers a number of middlewares out of the box

[Link](https://www.alexedwards.net/blog/which-go-router-should-i-use)

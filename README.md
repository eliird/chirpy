# chirpy
An API for a messaging app

## Motivation
To learn golang

## Installation
go get github.com/eliird/chirpy

## Endpoints
Provides the following endpoints

## GET /api/healthz", handlerReadiness   
>checks if the server is online

## GET /admin/metrics
>Returns the html page indicating the amount of hits on a page

## GET /api/reset", apiCfg.handlerReset
> reset the hit counter to the website


## POST /api/revoke", apiCfg.handlerRevoke
>Revokes the access token with login information
```
Header
{

}
Body
{

}
```
	
## "POST /api/refresh", apiCfg.handlerRefresh
>Refreshes the access token so user can stay logged in for longer period
```
Header
{

}
Body
{

}
```

## "POST /api/login", apiCfg.handlerLogin 
>Can be used to login to store and access messages accepts information in format
```
>Header
{

}

>Body
{
    "password": "something"
    "email":"some@gmail.com"
}
```


## "POST /api/users"
>Can be used to create a user in the database
```
>Header
{

}

>Body
{
    "password": "something"
    "email":"some@gmail.com"
}
```

## "PUT /api/users"
> Can be used to update the information of the user in the database
```
Header
{

}
Body
{

}
```

## POST /api/chirps"
> Add the message to the chirp
```
Header
{

}
Body
{

}
```

## GET /api/chirps"
>Get all the chirps. Has 2 query parameters `author_id` and `sort`.
>specifying `author_id` gets messages from the specific user
>`sort` can be specified as `asc` or `desc` to sort the default is `asc`
```
Header
{

}
Body
{

}
```

## GET /api/chirps/{chirpID}"
> Retreive the message based on the id from the database
```
Header
{

}
Body
{

}
```

## DELETE /api/chirps/{chirpID}"
> Delete the message based on the id from the database, after validating the user
```
Header
{

}
Body
{

}
```

## "POST /api/polka/webhooks"
> can be used as a way to manage a 3rd party payment servive to get a signal of if to update a user or no

```
Header
{

}
Body
{

}
```

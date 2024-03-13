# go-chirpy

ENV keys used are in the `.env.example`

# API links under /api

### GET /api/healtz

Health check endpoint

### GET /api/chirps

Returns an array of AllChirps

Accepts the following query params:

author_id

- id of desired author

sort

- asc
- desc

Example:

/api/chirps?author_id={id}

Returns an array of chirps for author with given id

Example

```
/api/chirps?sort=asc
```

Returns query sorted in ascending order

Those can be combined

```
/api/chirps?author_id=2&sort=asc
```

### POST /api/chirps

Requires Auth:
Accepts JSON:

```JSON
{
    "message": "Lorem ipsum dolor sit amet, qui minim"
}
```

Message must be max 140 char length

### GET /api/chirps/{id}

Get Chirp with given ID

### DELETE /api/chirps/{id}

Deletes Chirp with given ID if it belongs to current user

### GET /api/users

Returns an array of all users

### PUT /api/users

Updates currently authenticated user, accepts JSON of

```JSON
{
	"email": "test@gmail.com"
	"password": "abc123"
}
```

### POST /api/login

Returns Auth Bearer token and refresh token
Accepts body of:

```JSON
{
	"email": "test@gmail.com"
	"password": "abc123"
}
```

### POST /api/revoke

Revokes a given refresh token
Refresh token needs to be passed in the `Authorization` Header

### POST /api/refresh

Refreshes the API token given the refresh token
Refresh token needs to be passed in the `Authorization` Header like so

```
{
    "Authorization": "Bearer {token}"
}
```

### POST /api/polka/webhooks

Webhook for polkadok, requires Polkadot ApiKey to be passed like so:

```
{
    "Authorization": "ApiKey {token}"
}
```

# GET /admin/metrics

Requires auth used to be an admin, returns number of hits

# Night Owl

> Simple API

## How to run

You can simply run the project using following command

```
go run main.go
```

However, it's recommended to run this project using `docker-compose` with following command

```
docker-compose -f docker-compose.yml up --build
```

The app would available in `localhost:8000`, you can also set custom port by providing `APP_PORT` in `.env` file

## Api Docs

### POST /team

> Create a new team

```
curl --request POST \
  --url http://localhost:8000/team \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Chelsea"
}'

# Response sample
{
  "name": "Chelsea"
}
```

### GET /team

> Get team list

```
curl --request GET \
  --url http://localhost:8000/team

# Response sample
[
  {
    "id": 1,
    "name": "Chelsea"
  }
]
```

### GET /team/:id/player

> Get player list from a team id

```
curl --request GET \
  --url http://localhost:8000/team/1/player

# Response sample
{
  "id": 1,
  "name": "Chelsea",
  "Players": [
    {
      "name": "tesar",
      "teamId": 1
    },
    {
      "name": "hazard",
      "teamId": 1
    }
  ]
}
```

### POST /player

> Create a new player

```
curl --request POST \
  --url http://localhost:8000/player \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "hazard",
	"teamId": 1
}'

# Response sample
{
  "name": "hazard",
  "teamId": 1
}
```

### GET /player

> Get player list

```
curl --request GET \
  --url http://localhost:8000/player

# Response sample
[
  {
    "name": "tesar",
    "teamId": 1
  }
]
```

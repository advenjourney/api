# api

Based on this [tutorial](https://www.howtographql.com/graphql-go/1-getting-started/) to provide a basic api template.

## Run the api

### Start mysql as docker container
```bash
$ docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=advenjourney -d mysql:latest
```

### Clone Repository and start the server
```bash
$ git clone https://github.com/advenjourney/api
$ cd api
$ go run server.go
```

### Extend the model

Add api declaration to to `graph/schema.graphqls`
```bash
$ go run github.com/99designs/gqlgen generate
```
Implement the functions in `graph/schema.resolvers.go`

### Create a database migration

```bash
$ migrate create -ext sql -dir mysql -seq create_<model-name>_table
```

Implement the migration logic in the generated files in `internal/pkg/db/migrations/mysql`. There is on file for up and one for down migrations.

Then run the migration command to apply the migration and change the database schema

```bash
$   migrate -database mysql://root:dbpass@/advenjouirney -path internal/pkg/db/migrations/mysql up

```

### Test API

Open [Test server](http://http://localhost:8080) in your browser and insert the following queries or mutations

#### Get all offers

Send query
```
query {
  offers{
    title
    location,
    description,
    titleImageUrl,
    user{
      name
    }
  }
}
```

Results in
```
{
  "data": {
    "offer": [
      {
        "title": "something",
        "location": "somewhere",
        "description": "sfgsegsdgsd"
        "titleimageurl": "https://some.domain.com"
        "id": "1"
      },
            {
        "title": "something 2",
        "location": "somewhere 2",
        "description": "sfgsegsdgsd 2"
        "titleimageurl": "https://some.domain2.com"
        "id": "2"
      }
    ]
  }
}
```

#### Create a user

Send query
```
mutation {
  createUser(input: {username: "user1", password: "123"})
}
```

Results in
```
{
  "data": {
    "createUser": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODE0NjAwODUsImlhdCI6MTU4MTQ1OTc4NX0.rYLOM123kSulGjvK5VP8c7S0kgk03WweS2VJUUbAgNA"
  }
}
```

#### Create a offer

Send query
```
mutation {
  createLink(input: {title: "Coliving Request", location: "Gran Canaria", description: "Villa at the beach", titleimageurl: "https://image.adress.com"}){
    user{
      name
    }
  }
}
```

Results in
```
{
  "errors": [
    {
      "message": "access denied",
      "path": [
        "createLink"
      ]
    }
  ],
  "data": null
}
```

To create offr you must set the Authorization header
```
{
  "Authorization": "" // use your own generated token
}
```
Try again you should be able to create a new offer


### Known Issues

```bash
$ go run server.go
graph/schema.resolvers.go:9:2: package gqlgen/graph/generated is not in GOROOT (/usr/local/opt/go/libexec/src/gqlgen/graph/generated)
graph/schema.resolvers.go:10:2: package gqlgen/graph/model is not in GOROOT (/usr/local/opt/go/libexec/src/gqlgen/graph/model)
graph/schema.resolvers.go:11:2: package gqlgen/internal/auth is not in GOROOT (/usr/local/opt/go/libexec/src/gqlgen/internal/auth)
graph/schema.resolvers.go:12:2: package gqlgen/internal/offers is not in GOROOT (/usr/local/opt/go/libexec/src/gqlgen/internal/offers)
graph/schema.resolvers.go:13:2: package gqlgen/internal/pkg/jwt is not in GOROOT (/usr/local/opt/go/libexec/src/gqlgen/internal/pkg/jwt)
graph/schema.resolvers.go:14:2: package gqlgen/internal/users is not in GOROOT (/usr/local/opt/go/libexec/src/gqlgen/internal/users)
```

Help needed to solve this dependency issues that i do not understand yet
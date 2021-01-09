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

To create offer you must set the Authorization header
```
{
  "Authorization": "" // use your own generated token
}
```
Try again you should be able to create a new offer
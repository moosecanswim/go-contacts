# Golang contacs development

## Prerequisites

1. Docker

## Tutorial

Based on tutorial located here `https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b`

## Basic (static development) server

To run the basic server defined run defined in `main.go` use the following command:

```console
go run main.go
```

Go to `http://localhost:10000/` You will see the following:

```html
<p> W </>lcome to the HomePage </p>
```

That means its up and running on port 10000. This is defined in `main.go`'s`homePage()` function.

## Dynamic Development

```markdown
While in the same directory:

- To start this use `docker-compose up`. That will show the logs/ output from println. To run in backgrount append `-d` to the above command.

- To STOP run `docker-compose down`
```

To make this a dynamic development process i used go-watcher and docker-compose to create an isolated environment.

```yaml
services:
  contacts:
    image: canthefason/go-watcher:latest
    container_name: 'contacts'
    ports:
      - '8080:10000'
    environment:
      STAGE: 'docker-compose-land'
    volumes:
      - ~/go/src/github.com/moosecanswim/go-contacts:/go/src/github.com/moosecanswim/go-contacts
    command: watcher -run github.com/moosecanswim/go-contacts --config github.com/moosecanswim/go-contacts/settings.toml --watch github.com/moosecanswim/go-contacts
```

This docker-compose definition uses the `canthefason/go-watcher` image and volume mounts the directory that can change. This means that any changes located in `$PWD/go-contacts` directory will result in the docker container being rebuilt.

To use this environment run `docker-compose up`. This will spin up the containers defined in `docker-compose.yaml`. Check that it is working at `localhostt:8080`.

You might be wondering why this has a different port than in the Basic Development section. Well, in the `contacts` service we map the 8080 port on the host (your computers) port to the containers 10000 port.

## Define the Contacts strut and gloval variable

Self explanatory

```go
// Define struct of contacts
type Contact struct {
	First string `json:"first"`
	Last  string `json:"last'`
	Email string `"json: email"`
}

// Declare global Contacts array (simulation database)
var Contacts []Contact
```

To create an entry follow this (node the capitalized parts) :
`Contact{First: "Test FirstName:", Last: "Test LastName:", Email: "Test Email:"},`

Update the `main.go` main() do define two struct entries: kyle and dan

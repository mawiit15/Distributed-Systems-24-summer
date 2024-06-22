# Distributed systems summer 2024 lab

This is the Lab for the course `Parallele und Verteilte Systeme` in summer 2024.\
For this, a backend was implemented according to specs provided in `openapi-spec.json`. The code is written in `go` using a `MongoDB` databse service.\
The frontend component used here is not of our own creation and was provided by the lecturers of this course.

## Starting the application

To start the application, simply clone the repositoy and navigate to its root directory

```bash
cd path/to/the/repository
```

and then start up the application.

```bash
docker compose up --build
```

The application will be available at `http://localhost.com:8090`.

## Docker image

If you wish to just use the docker image of the implemented backend, you can pull the corresponding docker image from dockerhub using

```bash
docker pull gtorfo/todobackend:latest
```

or use it in your `docker-compose.yml` directly (excerpt shown):

```yaml
services:
  backend:
    image: gtorfo/todobackend:latest
```

## Documentation

Documentation on the API endpoints can be found under `http://localhost.com:8080/swagger/index.html`.
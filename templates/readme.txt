# {{.AppName}}

A simple description of what your project will do.

## Run the project

{{if .HasDocker}}### With docker
```sh
docker-compose up --build
docker-compose exec {{.AppName}}api /bin/sh

go build -o main .

./main
```

### Without docker

```sh
go build -o main .

./main
```
{{else}}```sh
go build -o main .

./main
```
{{end}}
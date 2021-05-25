# crud-go

A simple CRUD Golang API with MySQL & Docker

- Install
  * Go
  * Docker
  * docker-compose

- Usage: 

```bash
    cd $GOPATH/src

    git clone https://github.com/pedrocmart/crud-go

    cd crud-go

    docker-compose up
```

## Set up test database

```bash
    docker exec -it crud_go_db bash -l

    mysql -u root -p
```

## Api endpoint inside the container:

- localhost:5000/user

## Api endpoint out of container:

```bash
    cd $GOPATH/src/crud-go
    
    go run main.go -p 8080

```
- localhost:8080/user

-Usage + Demo:

[![DEMO](https://img.youtube.com/vi/ypUJ5-j-GD8/0.jpg)](https://youtu.be/ypUJ5-j-GD8)

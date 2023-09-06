# TODO CRUD API ![](https://img.shields.io/github/go-mod/go-version/mounis-bhat/todo-go)

An API developed in Go that performs CRUD operation on a Todo List

## Author

- [@mounis](https://www.github.com/mounis-bhat)

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

| ENV      | Example                                                       |
| -------- | ------------------------------------------------------------- |
| `PORT`   | 3000                                                          |
| `DSN`    | "host=x user=x password=x dbname=x port=3000 sslmode=disable" |
| `SECRET` | SECRET                                                        |

## Installation

```bash
  Install Go

  Install compile daemon on your machine to enable live reloading

  Create a postgres DB and use environment variables to sync
```

## Run

To run this project run

```bash
  CompileDaemon -command="./todo-go"

  or

  go run main.go
```

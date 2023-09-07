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

Install Go
Install compile daemon on your linux machine to enable live reloading (Only works on linux)
Create a postgres DB and use environment variables to sync

## Run

Run using compile daemon

```bash
  CompileDaemon -command="./todo-go"
```

or

```bash
  go run main.go
```

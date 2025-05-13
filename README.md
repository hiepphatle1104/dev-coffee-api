# dev-coffee-api

> simple coffee shop management

## Tech-stack
- Go
- Gorm
- Gin
- Mysql

## Setup

- Clone project
```zsh
git clone https://github.com/hiepphatle1104/dev-coffee-api.git
cd dev-coffee-api
```

- Setup environment
```dotenv
DB_CONN_URL=user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
PORT=3001
```

- Run
```zsh
go run main.go
```
An architectural foundation for web development projects built with Gin, offered as a reference template.

## Environment Requirements 

- [GO >= 1.8]

## Project Structure
```
├── README.md
├── config
│   └── cfg.go
├── config.yaml
├── controllers
│   ├── packages_controller.go
│   └── user_controller.go
├── dao
│   ├── packages_dao.go
│   └── user_dao.go
├── db
│   ├── db.go
│   └── redis.go
├── docs
├── go.mod
├── go.sum
├── logs
├── main.go
├── middleware
│   └── jwt_auth.go
├── models
│   ├── packages.go
│   └── user.go
├── pkg
├── router
│   └── routers.go
├── services
│   ├── packages_service.go
│   └── user_service.go
├── utils
│   ├── jwt.go
│   ├── logger
│   │   └── logger.go
│   └── resp
│       ├── resp.go
│       └── status_code.go
└── validator
    ├── custom_validator.go
    └── validator_errors.go
```

## Generate Code
```
go generate ./generate
```

## Deploy

First build the executable file
```
make build # for Linux
make cross # for mac/windows
```
Then put files of the ```build``` in your server and set the path of log and static file(html/css/js),and run the executable file.If 80 port is not allowed to use,consider the nginx proxy,or use the gin middleware [gin-reverseproxy](https://github.com/chenhg5/gin-reverseproxy) instead, which has some example in ```routers.go```. When the project start running, it will generate the ```pid```file in the root path of the project. Excute the following command to update your project. 
```
kill -INT $(cat pid) && ./morningo # graceful stop the process and restart
```

## Featrue

- [X] Test
- [X] Logger
- [X] Read & Write Connections
- [X] Redis 
- [X] Mysql 
- [X] Gorm 
- [X] Swagger
- [X] JWT
- [X] Code generator

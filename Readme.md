##  UP local pyroscope 

```sh
docker-compose up -d
```

## build projects

```go
go get -d ./
go build -o ./pushsimple ./cmd/app
```

run app 

```sh
./pushsimple
```

## Роуты для проверки работы профилирования 

http://127.0.0.1:8080/slowapp  # долгий запрос + mutex 500 ms  

http://127.0.0.1:8080/fastapp  # быстрый запрос  + mutex 50 ms  

http://127.0.0.1:8080/mem-leak # утекчка памяти +5 Мб при запросе 

http://127.0.0.1:8080/gorout   # не оптимальная работа goroutine   




Каталог авто


Запустить сервер 
```
go mod tidy
go run cmd/server/main.go
```

Получить список
```
http://localhost:37546/cars?page=1&perPage=10
```
```
http://localhost:37546/cars?page=1&perPage=10&id=001&regNum=a111aa11
```

Создать авто
```
curl -d '{"regNums":["a111aa11"]}' http://localhost:37546/cars
```

Обновить
```
curl -X "PATH" -d '{"id":"001","regNum":"e333ee33"}' http://localhost:37546/cars  
```

Удалить
```
curl -X "DELETE" http://localhost:37546/cars/001
```




Обычный круд на golang на тему "Каталог авто"

Запущенный сервер обрабатывает следующие эндпоинты:
 - Получть список - `GET http://localhost:37546/cars`
 - Удалить запись по id -`DELETE http://localhost:37546/cars/{id}` 
 - Частичное изменение записи - `PATH http://localhost:37546/cars` 
 - Создать запись -`POST http://localhost:37546/cars` 

### Запустить сервер 
```
go mod tidy
go run cmd/server/main.go
```

### Получить список:
```
http://localhost:37546/cars?page=1&perPage=10
```
Параметры по которым можео осуществлять фильтр: id, regNum, mark, model, owner
Обязательные параметры для пагинации: page, perPage
пример с фильтром:
```
http://localhost:37546/cars?page=1&perPage=10&id=001&regNum=a111aa11
```

### Создать авто:
```
curl -d '{"regNums":["a111aa11"]}' http://localhost:37546/cars
```

### Обновить:
Поля которые можно переать: id, regNum, mark, model, owner
Обязательные поля: id
```
curl -X "PATH" -d '{"id":"001","regNum":"e333ee33"}' http://localhost:37546/cars  
```

### Удалить:
```
curl -X "DELETE" http://localhost:37546/cars/001
```



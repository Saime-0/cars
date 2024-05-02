Реализовать каталог автомобилей. Необходимо реализовать следующее
1. Выставить rest методы
	1. Получение данных с фильтрацией по всем полям и пагинацией 
	2. Удаления по идентификатору
	3. Изменение одного или нескольких полей по идентификатору
	4. Добавления новых автомобилей в формате
```json
{
    "regNums": ["X123XX150"] // массив гос. номеров
}
```
2. При добавлении сделать запрос во внешнее АПИ, 
описанного сваггером (это описание некоторого внешнего АПИ,
 которого нет, но к которому надо обращаться. Реализованное, согласно
 описанию, АПИ будет использоваться при проверке)

```yaml
openapi: 3.0.3
info:
  title: Car info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: regNum
          in: query
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Car'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    Car:
      required:
        - regNum
        - mark
        - model
        - owner
      type: object
      properties:
        regNum:
          type: string
          example: X123XX150


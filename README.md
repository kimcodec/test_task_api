# Тестовое задание

## Эндпоинты

### Swagger

> /swagger/index.html

> http://158.160.57.179:8080/swagger/index.html

### Машины

#### Добавление

> /cars [POST]

##### Входные поля

- `[reg_num]`* - Массив гос номеров

#### Выходные поля

- `id` - ID машины
- `reg_num` - гос номер
- `mark` - марка автомобиля
- `model` - модель автомобиля
- `year` - год выпуска автомобиля
- `owner` - информация о владельце автомобиля

#### Получение

> /cars [GET]

##### Входные параметры(query)

- `offset` - стартовый элемент
- `limit` - лимит вывода
- `year` - год выпуска автомобиля
- `mark` - марка автомобиля
- `model` - модель автомобиля
- `reg_num` - гос номер

#### Выходные поля

- `id` - ID машины
- `reg_num` - гос номер
- `mark` - марка автомобиля
- `model` - модель автомобиля
- `year` - год выпуска автомобиля
- `owner` - информация о владельце автомобиля

#### Удаление

> /cars/{id} [DELETE]

#### Изменение

> /cars/{id} [PATCH]

##### Входные данные

- `reg_num` - гос номер
- `mark` - марка автомобиля
- `model` - модель автомобиля
- `year` - год выпуска автомобиля

#### Выходные поля

- `id` - ID машины
- `reg_num` - гос номер
- `mark` - марка автомобиля
- `model` - модель автомобиля
- `year` - год выпуска автомобиля
- `owner` - информация о владельце автомобиля

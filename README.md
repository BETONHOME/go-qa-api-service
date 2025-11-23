Тестовое задание: API-сервис для вопросов и ответов.

Инструкция по запуску.

git clone https://github.com/BETONHOME/go-qa-api-service   
cd go-qa-api-service  
docker compose up --build  
(Сервер работает на 8000 порту, БД на 5432)

Описание проекта.

REST API сервис, разработанный на языке Golang с использованием PostgreSQL.

Модели:

Questions:
- `id`         - идентификатор
- `text`       - содержимое вопроса
- `created_at` - дата создания

Answer:
- `id`          - идентификатор
- `question_id` - ссылка на вопрос
- `user_id`     - идентификатор пользователя
- `text`        - содержимое ответа
- `created_at`  - дата создания


Методы.

Questions:
- `GET /questions/` - получить список всех вопросов
- `POST /questions/` - создать новый вопрос
- `GET /questions/{id}` - получить вопрос со всеми ответами
- `DELETE /questions/{id}` - удалить вопрос

Answers:
- `POST /questions/{id}/answers/` - добавить ответ к вопросу
- `GET /answers/{id}` - получить конкретный ответ
- `DELETE /answers/{id}` - удалить ответ

Используемый стэк:
-`Go` 
-`GORM` - ORM
-`PostgreSQL` - база данных
-`Goose` - миграции
-`Docker, docker compose` - виртализация
-`testify` - тестирование
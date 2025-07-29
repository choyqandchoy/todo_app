# 📝 To-Do API на Go

Простое и чистое REST API для создания и управления задачами (To-Do) с поддержкой регистрации, логина и JWT-аутентификации.

---

## ⚙️ Стек технологий

- Go — основной язык
- PostgreSQL — база данных
- JWT (golang-jwt/v5) — для защиты API
- bcrypt — для хэширования паролей
- Gorilla Mux — роутинг
- Postman — тестирование API

---

## 🚀 Запуск проекта

### 1. Клонируй репозиторий

```bash
git clone https://github.com/choyqandchoy/todo_app.git
cd todo_app
```

### 2. Настрой базу данных PostgreSQL

```go
host=localhost
port=5432
user=postgres
password=Zamira30
dbname=postgres
```

```sql
-- Создание пользователя
CREATE USER user WITH PASSWORD 'Zamira30';

-- Создание базы данных
CREATE DATABASE user OWNER user;
CREATE DATABASE users OWNER user;
CREATE DATABASE user_roles OWNER user;
```

### 3. Создай таблицы

```sql
-- Таблица пользователей
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL
);

-- Таблица задач
CREATE TABLE tasks (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  title TEXT NOT NULL,
  description TEXT
);

-- Таблица ролей (опционально, если используешь)
CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE
);
```

### 4. Установите зависимости

```bash
go mod tidy
```

### 5. Запустите приложение

```bash
go run main.go
```
- Сервер запустится на: 📍 http://localhost:8080

---

## 🧪 Примеры запросов (Postman):

🔐 Регистрация  
• POST /register 
• Body (JSON):
```json
{
  "username": "john",
  "password": "123456"
}
```

🔑 Логин
• POST /login
• Body (JSON):
```json
{
  "username": "john",
  "password": "123456"
}
```
• Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR..."
}
```

✅ Создание задачи  
• POST /tasks  
• Headers:  
  • Authorization: 
  Bearer <ваш JWT токен>  
• Body (JSON):
```json
{
  "title": "Купить продукты",
  "description": "Хлеб, молоко, сыр"
}
```

🔍 Получить задачу по ID  
• GET /tasks/{id}  
• Headers:  
  • Authorization: Bearer <ваш JWT токен>
  


## 📌 Автор  Проект создан в обучающих целях. 
GitHub: @choyqandchoy

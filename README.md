# Todo App

Простое приложение для управления задачами, написанное на Go с использованием PostgreSQL.

## Требования

- Go 1.21+
- Docker и Docker Compose (рекомендуется)
- PostgreSQL 12+ (если не используете Docker)

## Установка и настройка

### Вариант 1: С Docker (рекомендуется)

1. Установите Docker и Docker Compose:
   ```bash
   sudo pacman -S docker docker-compose
   sudo systemctl start docker
   sudo systemctl enable docker
   sudo usermod -aG docker $USER
   # Перезайдите в систему для применения изменений группы
   ```

2. Запустите приложение с PostgreSQL в Docker:
   ```bash
   ./start_with_docker.sh
   ```

### Вариант 2: Локальная установка PostgreSQL

1. Установите зависимости:
   ```bash
   go mod tidy
   ```

2. Настройте базу данных:
   ```bash
   ./setup_db.sh
   ```

   Или настройте вручную:

   1. Установите PostgreSQL:
      ```bash
      sudo pacman -S postgresql
      ```

   2. Запустите PostgreSQL:
      ```bash
      sudo systemctl start postgresql
      sudo systemctl enable postgresql
      ```

   3. Создайте базу данных:
      ```bash
      sudo -u postgres psql
      CREATE DATABASE todo_app;
      CREATE USER postgres WITH PASSWORD 'password';
      GRANT ALL PRIVILEGES ON DATABASE todo_app TO postgres;
      \q
      ```

### 3. Настройка конфигурации

Отредактируйте файл `configs/config.yml` под ваши настройки базы данных:

```yaml
server:
  port: "8000"

database:
  host: "localhost"
  port: "5432"
  username: "postgres"
  password: "password"
  dbname: "todo_app"
  sslmode: "disable"
```

### 3. Запуск приложения

#### С Docker:
```bash
./start_with_docker.sh
```

#### Без Docker:
```bash
go run cmd/main.go
```

Приложение будет доступно по адресу: http://localhost:8000

### 4. Управление Docker контейнерами

```bash
# Запуск только PostgreSQL
docker-compose up -d postgres

# Остановка всех контейнеров
docker-compose down

# Просмотр логов PostgreSQL
docker-compose logs postgres

# Подключение к базе данных
docker-compose exec postgres psql -U postgres -d todo_app
```

## API Endpoints

### Аутентификация
- `POST /auth/sign-up` - Регистрация пользователя
- `POST /auth/sign-in` - Вход в систему

### Списки задач
- `POST /api/lists/` - Создать список
- `GET /api/lists/` - Получить все списки
- `GET /api/lists/:id` - Получить список по ID
- `PUT /api/lists/:id` - Обновить список
- `DELETE /api/lists/:id` - Удалить список

### Задачи
- `POST /api/lists/:id/items/` - Создать задачу
- `GET /api/lists/:id/items/` - Получить все задачи списка
- `GET /api/lists/:id/items/:item_id` - Получить задачу по ID
- `PUT /api/lists/:id/items/:item_id` - Обновить задачу
- `DELETE /api/lists/:id/items/:item_id` - Удалить задачу

## Структура проекта

```
├── cmd/
│   └── main.go              # Точка входа приложения
├── configs/
│   └── config.yml           # Конфигурация
├── pkg/
│   ├── handler/             # HTTP обработчики
│   ├── repository/          # Слой данных
│   └── service/             # Бизнес-логика
├── schema/                  # SQL миграции
└── server.go               # HTTP сервер
```

## Миграции

Миграции запускаются автоматически при старте приложения. Файлы миграций находятся в папке `schema/`:

- `000001_init.up.sql` - Создание таблиц
- `000001_init.down.sql` - Удаление таблиц

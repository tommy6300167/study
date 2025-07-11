# Todo API Documentation

Base URL: `http://localhost:8080/api/v1`

## Endpoints

### 1. Create Todo
**POST** `/todos`

Create a new todo item.

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go programming",
    "description": "Study Go fundamentals and build REST API"
  }'
```

**Response:**
```json
{
  "id": 1,
  "title": "Learn Go programming",
  "description": "Study Go fundamentals and build REST API",
  "completed": false,
  "created_at": "2025-07-11T11:30:20.123456Z",
  "updated_at": "2025-07-11T11:30:20.123456Z"
}
```

### 2. Get All Todos
**GET** `/todos`

Retrieve all todo items.

```bash
curl -X GET http://localhost:8080/api/v1/todos
```

**Response:**
```json
[
  {
    "id": 1,
    "title": "Learn Go programming",
    "description": "Study Go fundamentals and build REST API",
    "completed": false,
    "created_at": "2025-07-11T11:30:20.123456Z",
    "updated_at": "2025-07-11T11:30:20.123456Z"
  },
  {
    "id": 2,
    "title": "Write tests",
    "description": "Add unit tests for API endpoints",
    "completed": true,
    "created_at": "2025-07-11T11:35:15.654321Z",
    "updated_at": "2025-07-11T11:40:30.789012Z"
  }
]
```

### 3. Get Todo by ID
**GET** `/todos/{id}`

Retrieve a specific todo item by ID.

```bash
curl -X GET http://localhost:8080/api/v1/todos/1
```

**Response:**
```json
{
  "id": 1,
  "title": "Learn Go programming",
  "description": "Study Go fundamentals and build REST API",
  "completed": false,
  "created_at": "2025-07-11T11:30:20.123456Z",
  "updated_at": "2025-07-11T11:30:20.123456Z"
}
```

**Error Response (404):**
```json
{
  "error": "todo with id 999 not found"
}
```

### 4. Update Todo
**PUT** `/todos/{id}`

Update an existing todo item. All fields are optional.

```bash
# Update title and description
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Master Go programming",
    "description": "Complete advanced Go concepts and patterns"
  }'
```

```bash
# Mark as completed
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "completed": true
  }'
```

```bash
# Update only title
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated title"
  }'
```

**Response:**
```json
{
  "id": 1,
  "title": "Master Go programming",
  "description": "Complete advanced Go concepts and patterns",
  "completed": true,
  "created_at": "2025-07-11T11:30:20.123456Z",
  "updated_at": "2025-07-11T11:45:10.987654Z"
}
```

### 5. Delete Todo
**DELETE** `/todos/{id}`

Delete a todo item by ID.

```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

**Response:**
```json
{
  "message": "todo deleted successfully"
}
```

**Error Response (404):**
```json
{
  "error": "todo with id 999 not found"
}
```

## Error Responses

### 400 Bad Request
Invalid request body or parameters.

```json
{
  "error": "Key: 'CreateTodoRequest.Title' Error:Tag: 'required'"
}
```

### 404 Not Found
Todo item not found.

```json
{
  "error": "todo with id 123 not found"
}
```

## Testing Workflow

### Complete CRUD Test Sequence

```bash
# 1. Create a new todo
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Test Todo",
    "description": "This is a test todo item"
  }'

# 2. Get all todos
curl -X GET http://localhost:8080/api/v1/todos

# 3. Get specific todo (replace 1 with actual ID)
curl -X GET http://localhost:8080/api/v1/todos/1

# 4. Update the todo
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Test Todo",
    "completed": true
  }'

# 5. Delete the todo
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

## Data Models

### Todo Object
```json
{
  "id": "integer (auto-generated)",
  "title": "string (required)",
  "description": "string (optional)",
  "completed": "boolean (default: false)",
  "created_at": "timestamp (auto-generated)",
  "updated_at": "timestamp (auto-updated)"
}
```

### Create Todo Request
```json
{
  "title": "string (required)",
  "description": "string (optional)"
}
```

### Update Todo Request
```json
{
  "title": "string (optional)",
  "description": "string (optional)",
  "completed": "boolean (optional)"
}
```
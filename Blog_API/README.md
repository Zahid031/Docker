# Dockerized Blog API (Django+PostgreSQL+AWS+CI/CD)
## Hosted on AWS 


A RESTful API for a blogging platform built with Django Rest Framework 

# Live API 
```bash
http://ec2-67-202-2-33.compute-1.amazonaws.com:8001/
```

## Features

- User registration and authentication with JWT
- CRUD operations for blog posts
- Author-based permissions (only authors can edit/delete their posts)
- Image upload functionality for posts
- API pagination
- Dockerized application


## Technologies Used
- Django Rest Framework
- PostgreSQL
- JWT Authentication
- Docker & Docker Compose

## API Endpoints

### Authentication
- `POST /auth/register/` - User registration
- `POST /auth/login/` - User login (JWT token)
- `POST /api/token/refresh/` - Refresh JWT token

### Blog Posts
- `GET /blog/posts/` - List all posts (paginated)
- `POST /blog/posts/` - Create a new post (authenticated users only)
- `GET /blog/posts/<id>/` - Retrieve a specific post
- `PUT /blog/posts/<id>/` - Update a post (only by post author)
- `DELETE /blog/posts/<id>/` - Delete a post (only by post author)

## Setup Instructions

### Using Docker (recommended)

1. Clone the repository:
```bash
git clone origin git@github.com:Zahid031/Blog_API.git
cd Blog-API
```

2. Build and run the Docker containers:
```bash
docker-compose up --build
```

3. The API will be available at http://localhost:8001/

### Manual Setup

1. Clone the repository:
```bash
git clone origin git@github.com:Zahid031/Blog_API.git
cd Blog-API
```

2. Create and activate a virtual environment:
```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

3. Install dependencies:
```bash
pip install -r requirements.txt
```

4. Configure PostgreSQL:
   - Create a PostgreSQL database named `blog_db`
   - Update the database configuration in `settings.py` 

5. Run migrations:
```bash
python manage.py migrate
```

6. Start the development server:
```bash
python manage.py runserver
```

7. The API will be available at http://localhost:8000/

## API Usage Examples

## You can use any browser it will have a user interface to check the API.

### Register a new user
```bash
curl -X POST http://localhost:8000/auth/register/ \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser", "password":"TestPass123", "confirm_password":"TestPass123", "email":"test@example.com", "first_name":"Test", "last_name":"User"}'
```

### Get JWT token (login)
```bash
curl -X POST http://localhost:8000/auth/login/ \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser", "password":"TestPass123"}'
```

### Create a new post
```bash
curl -X POST http://localhost:8000/blog/posts/ \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"My First Post", "content":"This is the content of my first post"}'
```

### List all posts
```bash
curl -X GET http://localhost:8000/blog/posts/
```

### Get a specific post
```bash
curl -X GET http://localhost:8000/blog/posts/1/
```

### Update a post(Author)
```bash
curl -X PUT http://localhost:8000/blog/posts/1/ \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Post Title", "content":"Updated content"}'
```

### Delete a post (Author)
```bash
curl -X DELETE http://localhost:8000/blog/posts/1/ \
  -H "Authorization: Bearer <your_token>"
```

# Folio

## MVP

* Users can sign up and login to their account
* Users can create posts
* Users can view their feed
* Users can view profiles
* Users can like posts

## Data Models

```sql

User
- id
- email
- username
- password_hash
- created_at

Post
- id
- user_id
- title
- description
- url
- like_count
- created_at

```

## Endpoints

- `POST` `/auth/register`
- `POST` `/auth/login`

- `GET`    `/users/{username}`

- `POST   /posts`
- `GET    /feed`
- `POST   /posts/{id}/like`
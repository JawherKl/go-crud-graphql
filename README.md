# Go CRUD GraphQL API

![go-crud-graphql](https://github.com/JawherKl/go-crud-graphql/blob/main/go_graphql.png)

This project is a CRUD (Create, Read, Update, Delete) API built with Go, GraphQL, and PostgreSQL. It uses the `gqlgen` library for GraphQL and `go-pg` for PostgreSQL interaction. The API allows you to create and retrieve movies with a title, URL, and release date.

## Features
- Create a new movie with a title, URL, and automatically generated release date.
- Retrieve all movies.

## Technologies
- Go
- GraphQL
- PostgreSQL
- gqlgen
- go-pg
- godotenv

## Prerequisites

- Go (version 1.16+)
- PostgreSQL
- `go-pg`
- `gqlgen`
- `godotenv`

## Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/your-username/go-crud-graphql.git
    cd go-crud-graphql
    ```

2. **Install dependencies:**
    ```sh
    go mod download
    ```

3. **Create a PostgreSQL database:**
    ```sql
    CREATE DATABASE your_db_name;
    ```

4. **Set up environment variables:**
    Create a `.env` file in the root directory with the following content:

    ```plaintext
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    DB_HOST=localhost
    DB_PORT=5432
    ```

5. **Create the `movies` table:**
    ```sql
    CREATE TABLE movies (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        url TEXT NOT NULL,
        release_date TEXT NOT NULL
    );
    ```

6. **Run the server:**
    ```sh
    go run server.go
    ```

## Usage
Once the server is running, you can access the GraphQL playground at `http://localhost:8080/`.

### Example Mutation
To create a new movie, use the following GraphQL mutation:

```graphql
mutation createMovie {
  createMovie(
    input: {
      title: "Rise of GraphQL Warrior Pt1"
      url: "https://riseofgraphqlwarriorpt1.com/"
    }
  ){
    id
    releaseDate
  }
}
```

To update a movie, use the following GraphQL mutation:

```graphql
mutation {
  updateMovie(input: {
    id: 7
    title: "GraphQL in Action (Updated)"
  }) {
    id
    title
    url
    releaseDate
  }
}
```

To delete a movie, use the following GraphQL mutation:

```graphql
mutation {
  deleteMovie(id: 1)
}
```

To get all movie, use the following GraphQL mutation:

```graphql
query {
  movies {
    id
    title
    url
    releaseDate
  }
}
```

To get movie by id, use the following GraphQL mutation:

```graphql
query {
  movie(id: 1) {
    id
    title
    url
    releaseDate
  }
}
```

## Stargazers over time
[![Stargazers over time](https://starchart.cc/JawherKl/go-crud-graphql.svg?variant=adaptive)](https://starchart.cc/JawherKl/go-crud-graphql)

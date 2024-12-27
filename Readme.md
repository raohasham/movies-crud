This is a simple Movie API built with Go to demonstrate CRUD (Create, Read, Update, Delete) operations using the Gorilla Mux router. This project is for practice and to get familiar with basic CRUD operations in Go.

## Project Setup

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/movie-api.git
    cd movie-api
    ```

2. **Install dependencies:**
    Ensure you have Go installed on your machine. Install the Gorilla Mux package:
    ```sh
    go get -u github.com/gorilla/mux
    ```

3. **Run the application:**
    ```sh
    go run main.go
    ```

4. **API Endpoints:**
    - `GET /movies` - Get all movies
    - `GET /movie/{id}` - Get a single movie by ID
    - `POST /movies` - Create a new movie
    - `PUT /movie/{id}` - Update a movie by ID
    - `DELETE /movie/{id}` - Delete a movie by ID

## Example Requests

### Get All Movies

```sh
curl -X GET "http://localhost:8000/movies"

Get a Single Movie

curl -X GET "http://localhost:8000/movie/1"

Create a New Movie
sh

curl -X POST -H "Content-Type: application/json" -d '{
  "isbn": "978-3-16-148410-0",
  "name": "New Movie",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}' "http://localhost:8000/movies"

Update a Movie
sh

curl -X PUT -H "Content-Type: application/json" -d '{
  "id": "1",
  "isbn": "978-3-16-148410-0",
  "name": "Updated Movie",
  "director": {
    "firstname": "Jane",
    "lastname": "Doe"
  }
}' "http://localhost:8000/movie/1"

Delete a Movie

curl -X DELETE "http://localhost:8000/movie/1"

Project Structure

    main.go: The main file containing the API implementation.
    README.md: Documentation for the project.

Dependencies

    github.com/gorilla/mux: A powerful URL router and dispatcher for Golang.

License

This project is licensed under the MIT License. See the LICENSE file for details.
```
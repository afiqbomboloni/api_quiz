
# Quiz Api


## Installation

1. Clone the repository.

    git clone https://github.com/afiqbomboloni/api_quiz.git

3. Install dependencies using `go mod tidy`.

## Set Up Database

1. Open terminal and type

    mysql -u (your username) -p (your pass)

2. Create database

    create database quiz_db


## Running the Application

  

To start the application, use the following command:

  

go run cmd/main.go

The application will run on `http://localhost:8080` by default.


To run seeder, use the following command:

go run db/seeders/main.go

# Go Authentication with MySQL

This is a simple Go project for user authentication using the Fiber web framework and MySQL database.

## Technologies Used

- Go
- Fiber
- MySQL

## Installation

1. Clone the repository: `git clone https://github.com/arnabtechie/go-auth-mysql.git`
2. Install the dependencies: `go mod download`
3. Set up the database: 
   * Create a MySQL database
   * Create the `users` table using the following query:
     ```
     CREATE TABLE `users` (
       `id` int(11) NOT NULL AUTO_INCREMENT,
       `name` varchar(255) DEFAULT NULL,
       `email` varchar(255) DEFAULT NULL,
       `password` varchar(255) DEFAULT NULL,
       `created_at` datetime DEFAULT NULL,
       PRIMARY KEY (`id`),
       UNIQUE KEY `email` (`email`)
     ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
     ```
4. Rename the `.env.example` file to `.env` and fill in the details for your MySQL connection.
5. Run the application using `go run main.go`.

## Author

Arnab Banerjee - [@arnabtechie]

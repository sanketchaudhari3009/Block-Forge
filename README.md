# Block Forge - Simple Blockchain for Book Checkouts

Block Forge is a basic blockchain implementation in Go that enables the management of book checkouts. This project serves as a starting point for understanding blockchain concepts and can be expanded upon for more complex applications.

## Features

- **Blockchain Management:** Create and manage a simple blockchain for tracking book checkouts.
- **Book Checkout:** Add book checkout transactions to the blockchain.
- **Genesis Block:** Automatically generates a genesis block to start the blockchain.
- **Hash Validation:** Ensures the integrity of blocks through hash validation.

## Getting Started

### Prerequisites

- Go installed on your machine

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/sanketchaudhari3009/Block-Forge.git
2. Navigate to the project directory:

    ```bash
    cd block-forge
    ```

3. Run the Go application:

    ```bash
    go run main.go
    ```

    The application should now be running and listening on port 3000.

## Usage

Block Forge provides a basic RESTful API for managing the blockchain and book checkouts. You can use tools like Postman to interact with the API:

- **GET /:** Retrieve the entire blockchain.
- **POST /:** Add a new book checkout to the blockchain.
- **POST /new:** Create a new book and add it to the blockchain.

Example request for adding a book checkout:

```json
POST http://localhost:3000/
Body:
{
    "book_id": "12345",
    "user": "John Doe",
    "checkout_date": "2023-10-10"
}
```
Example request for creating a new book:

```json
POST http://localhost:3000/new
Body:
{
    "title": "Sample Book",
    "author": "Author Name",
    "publish_date": "2023-10-05",
    "isbn": "978-1234567890"
}
```


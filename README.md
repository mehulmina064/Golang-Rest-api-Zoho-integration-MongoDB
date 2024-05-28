# Golang REST API Boilerplate with Zoho Books Item Integration

This project is a Golang REST API boilerplate using the Gin framework, with MongoDB as the database. It is designed for managing teams and user roles, and additionally handles product management for an e-commerce website. The product structure is directly compatible with Zoho Books item structure, making integration with Zoho seamless.

## Features
- User and role management
- Product management
- Integration-ready with Zoho Books

## Technologies Used
- **Backend**: Golang with Gin framework
- **Database**: MongoDB

## Getting Started

### Prerequisites
- Go installed on your machine
- MongoDB set up and running

### Installation

1. **Install Go**

    ```sh
    sudo apt install golang-go
    ```

2. **Set up environment path**

    ```sh
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$PATH
    ```

3. **Set up environment variables**

    Create a `.env` file in the root directory and add the necessary environment variables.

4. **Install Air for auto-refresh and rerun setup**

    ```sh
    go install github.com/cosmtrek/air@latest
    ```

    **Set up the path**

    ```sh
    alias air='$(go env GOPATH)/bin/air'
    ```

    **Initialize Air**

    ```sh
    air init
    ```

### Running the Server

**Using Air for auto-refresh**

```sh
air
```

**Manual run**

```sh
go run main.go
```

### Troubleshooting

**Server not ending issue**

List open ports:

```sh
lsof -i
```

Kill the process by PID:

```sh
kill -9 <pid>
```

## API Documentation

You can find the API documentation and Postman collection [Link](https://elements.getpostman.com/redirect?entityId=23544751-7cb9dedf-4167-4c42-9d1c-449429507cfe&entityType=collection).

## Contributing

Contributions are welcome. Please fork the repository and create a pull request.

##  Deployed Server Link

The server is deployed on Render machine. You can access it at server-link or [BaseURL](https://restapi-golang-with-mongodb.onrender.com).

## License

This project is licensed under the MIT License.

---

For detailed API usage and examples, refer to the Postman collection linked above. Happy coding!

---

### Contact

For any inquiries or support, please contact [mehulmeena064@gmail.com](mailto:mehulmeena064@gmail.com).

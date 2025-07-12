### Project Description

This is a simple HTTP server written in Go that logs incoming requests. The server listens on `localhost` at port `8085`. For every incoming request, it prints detailed information to the console, including:

- **HTTP Method**: The method used (e.g., GET, POST).
- **URL**: The full URL of the request.
- **Path**: The path portion of the URL.
- **Query Parameters**: Any query parameters included in the URL.
- **Client IP Address**: The IP address of the client making the request.
- **Headers**: All headers sent with the request.
- **Request Body**: If present, the body of the request is also logged.

### Usage

1. Ensure you have Go installed on your system.
2. Save the following code into a file named `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the application using the command:
   ```bash
   go run main.go
   ```
5. The server will start, and you can test it by making HTTP requests to `http://localhost:8085`.

### Example

You can use tools like `curl` or Postman to send a request:

```bash
curl -X GET 'http://localhost:8085/?key=value' --header "Content-Type: text/plain" -d "Hello, World"
```

output:
```
----- New query -----
Mathod: GET
URL: /?key=value
Path: /
Query params: key=value
IP client: [::1]:58263
Headers:
  Content-Type: text/plain
  Content-Length: 12
  User-Agent: curl/8.7.1
  Accept: */*
Body:
Hello, World

```

This will trigger the logging of the HTTP method (GET), URL (`http://localhost:8085/?key=value`), path (`/`), query parameters (`key=value`), client IP address, headers (`Content-Type`), and body (`Hello, World!`).

### Dependencies

**Note:** This project has zero dependencies. It uses only the standard Go libraries.

### Notes

This server is designed for basic request logging and educational purposes. It does not implement any handling for responses beyond sending a 200 OK status.

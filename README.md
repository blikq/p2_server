# p2 Server

Player 2 is a lightweight, real-time entity tracking system written in Go. It utilizes TCP connections for efficient communication.

## Features

- Real-time entity tracking
- Lightweight and efficient
- Written in Go
- Utilizes TCP connections

## Installation

To install and run the p2 Server, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/p2-server.git`
2. Navigate to the project directory: `cd p2-server`
3. Build the server: `go build`
4. Run the server: `./p2-server`

## Usage

Once the p2 Server is up and running, you can start tracking entities by establishing a TCP connection to the server. Here's an example of how to do it in Go:

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Failed to connect to p2 Server:", err)
        return
    }
    defer conn.Close()

    // Send entity data to the server
    // ...

    // Receive updates from the server
    // ...
}
```

Make sure to replace `"localhost:8080"` with the actual address and port of your p2 Server instance.

## Contributing

Contributions are welcome! If you'd like to contribute to the p2 Server project, please follow these guidelines:

1. Fork the repository
2. Create a new branch
3. Make your changes
4. Test your changes
5. Submit a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

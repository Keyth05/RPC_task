# 🚀 RPC Hello World in Go

This project demonstrates how to create a basic implementation of **RPC (Remote Procedure Call)** in **Go**. The RPC server receives a request from the client, processes the function, and returns a response to the client.

## 🛠️ Requirements

Before getting started, make sure you have the following installed:

- **Go** (version 1.17 or higher)

---
## 🚀 Execution Process

### ⚙️ 1. Running the server

In the terminal, navigate to the project folder and run the RPC server using:
``` go run server.go```

You should see the following message: 
   ```Server is running on port 1234...```

### 👩🏻‍💻 2. Running the client

Then, in the terminal, navigate to the project folder and run the RPC client using:
```go run client.go ```

The client will connect to the server and receive a response like:
```Server response: Hello Keyth from RPC!```
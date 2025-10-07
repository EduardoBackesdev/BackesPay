# BackesPay - Payment System with Go

BackesPay is a payment system developed in Go that demonstrates the advanced use of goroutines and WebSocket for asynchronous transaction processing and real-time communication.

## 🚀 Main Features

### Goroutines for Transaction Processing
The system uses goroutines to process BED (TED) payments asynchronously, allowing:
- Execution of transactions during business hours (8 AM to 3 PM, business days)
- Non-blocking transaction processing
- Continuous time verification for execution
- Real-time transaction status notifications

```go
func Routine_bed(data repositories.BedRequest) {
    for {
        // Verifica horário comercial e processa a transação
        // Notifica o cliente via WebSocket
    }
}
```

### WebSocket for Real-Time Communication
WebSocket implementation for:
- Persistent client connection
- Instant notifications about transaction status
- Management of multiple connections using mutex for thread safety
- Connection pool for efficient resource control

```go
var Map_clients = make(map[int]*websocket.Conn)
var Mutex = &sync.Mutex{}
```

## 🛠 Technologies Used

- **Go** - Main language  
- **Gin** - Web framework  
- **MySQL** - Database  
- **Gorilla WebSocket** - WebSocket implementation  
- **JWT** - Authentication  
- **Decimal** - Precision in monetary operations  

## 📊 Project Structure

- `cmd/app`: Application entry point  
- `internal/api`: API handlers and services  
- `internal/repositories`: Database access layer  
- `internal/ws`: WebSocket connection management  
- `lib/goroutines`: Asynchronous routine implementations  
- `routes`: API route configuration  

## 🔐 Features

- Account creation and authentication  
- Transfers via **BIX** (instant)  
- Transfers via **BED** (scheduled/business hours)  
- Balance inquiry  
- Recipient verification  
- Real-time transaction notifications  

## 🔄 Transaction System

### BIX (Instant Transfer)
- Synchronous processing  
- Immediate validation  
- Instant balance update  

### BED (Scheduled Transfer)
- Asynchronous processing via goroutines  
- Execution only during business hours  
- Real-time notification via WebSocket  
- Thread-safe using mutex  

## 🚧 Upcoming Implementations

- **Transaction History**: System for tracking and querying all completed operations  
  - Integration of a notification queue with RabbitMQ  

## 💡 Project Purpose

This project was developed with the purpose of showcasing my skills in:
- Concurrent programming with goroutines  
- Real-time communication with WebSocket  

---

⌨️ Developed by **EduardoBackesdev**

# Bookings & Reservations Project

A web application for managing bookings and reservations, built with Go 1.23.

## Features

- RESTful routing with [Chi](https://github.com/go-chi/chi)
- Session management using [Alex Edwards SCS](https://github.com/alexedwards/scs)
- CSRF protection via [Nosurf](https://github.com/justinas/nosurf)

## Getting Started

1. **Clone the repository**
    ```bash
    git clone https://github.com/yourusername/bookings.git
    cd bookings
    ```

2. **Install dependencies**
    ```bash
    go mod tidy
    ```

3. **Run the application**
    ```bash
    go run main.go
    ```

## Technologies Used

- Go 1.23
- Chi Router
- SCS Session Management
- Nosurf CSRF Protection

## License

MIT

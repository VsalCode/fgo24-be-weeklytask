# E-Wallet Backend

This project is a Backend E-Wallet system developed by me to complete a weekly task in the Full-Stack Web Development Bootcamp at Koda Academy. It builds upon a previous E-Wallet project from an earlier weekly task, focusing on creating a robust backend for an E-Wallet application.

## API Endpoints Documentation
| Method | Endpoint                  | Description                |
|--------|---------------------------|----------------------------|
| POST   | ```/auth/register```            | Authentication Register    |
| POST   | ```/auth/login```               | Authentication Login       |
| PATCH  | ```/profile```            | Update profile             |
| GET    | ```/profile```            | Current User profile       |
| PUT    | ```/profile```            | Update/change Avatar/profile picture       |
| GET    | ```/users?search=```            | List users, Find user by name and phone   |
| GET    | ```/wallets```                  | Get wallet balance         |
| GET    | ```/wallets/records```          | Get finance records (balance, income, expense)  |
| POST   | ```/transactions/topup```       | Top up transaction         |
| POST   | ```/transactions/transfer```     | Transfer transaction       |
| GET    | ```/transactions```             | Get transaction history     |


## ERD (Entity Relationship Diagram)

  ```mermaid
erDiagram
  direction LR

  users {
      int id PK
      string fullname
      string email
      string password
      int pin
      string phone
      string avatar
      timestamp created_at
  }

  wallets {
      int id PK
      int user_id FK
      decimal balance
      timestamp updated_at
  }

  topup {
      int id PK
      int user_id FK
      decimal topup_amount
      timestamp topup_date
      int method_id FK
      bool success
  }

  payment_method {
      int id PK
      string method_name
  }

  transfers {
      int transfer_id PK
      int sender_user_id FK
      int receiver_user_id FK
      decimal transfer_amount
      string notes
      timestamp transfer_date
      bool success
  }

  users ||--o| wallets: has
  users ||--o| topup: "do"
  topup }o--|| payment_method: "uses"
  users ||--o| transfers: "sends"
  users ||--o{ transfers: "receives"
  ```

## Installation
1. Clone the repository:
```
git clone https://github.com/VsalCode/fgo24-be-weeklytask.git
cd fgo24-be-weeklytask
```

2. Create a .env file: Add the necessary environment variables in the .env file (e.g., database URL, JWT secret).

3.Install dependencies:
```
go mod tidy
```

4. Run the application:
```
go run main.go
```

## Depedencies
- Gin Gonic (github.com/gin-gonic/gin)
- JWT V5 (github.com/golang-jwt/jwt/v5)
- PGX (github.com/jackc/pgx/v5)
- Godotenv (github.com/joho/godotenv)

## How To Contribute
Pull requests are welcome! For major changes, please open an issue first to discuss your proposed changes. 

## License
This project is licensed under the [MIT](https://opensource.org/license/mit) License
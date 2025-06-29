  # Backend E-Wallet

  ### ERD

  ```mermaid
erDiagram
  direction LR

  users {
      int id PK
      string fullname
      string email
      string phone
      string password
      int pin
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
      timestamp transfer_date
      bool success
  }

  users ||--o| wallets: has
  users ||--o| topup: "do"
  topup }o--|| payment_method: "uses"
  users ||--o| transfers: "sends"
  users ||--o{ transfers: "receives"
  ```

  
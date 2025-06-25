  # Backend E-Wallet

  ### ERD

  ```mermaid
  erDiagram
  direction LR
  users ||--o{ sessions : "has"
  users ||--o{ wallets : "owns"
  users ||--o{ transactions : "initiates"
  payment_methods ||--o{ transactions : "used_in"
  wallets ||--o{ transactions : "affects"

  users {
    id int PK
    fullname string
    email string
    phone string
    password string
    pin int
    created_at timestamp
    updated_at timestamp
  }

  sessions {
    id int PK
    token string 
    created_at timestamp
    expired_at timestamp
    user_id int FK
  }

  payment_methods {
    id int PK
    name string
  }

  wallets {
    id int PK
    balance decimal
    user_id int FK 
  }

  transactions {
    id int PK
    amount decimal
    type string
    status string
    user_id int FK
    payment_method_id int FK
    wallet_id int FK
  }
  ```

  
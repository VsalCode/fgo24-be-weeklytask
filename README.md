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

    balance {
        int id PK
        int user_id FK
        decimal amount
        timestamp updated_at
    }

    transactions {
        int id PK
        int user_id FK
        decimal amount
        bool success
        string transaction_type  
        timestamp created_at
    }

    topup {
        int id PK
        int transaction_id FK
        decimal topup_amount
        timestamp topup_date
        int method_id FK
    }

    payment_method {
      int id PK
      string method_name
    }

    transfers {
        int transfer_id PK
        int transaction_id FK
        int sender_user_id FK
        int receiver_user_id FK
        decimal transfer_amount
        timestamp transfer_date
    }

    users ||--o| balance: has
    users ||--o| transactions: initiates
    transactions ||--o| topup: includes
    topup |o--|| payment_method : "has"  
    transactions ||--o| transfers: includes
    users ||--o| transfers: sender
    users ||--o| transfers: receiver


  ```

  
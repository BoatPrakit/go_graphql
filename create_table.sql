CREATE TABLE IF NOT EXISTS contact (
    contact_id INTEGER primary key autoincrement,
    name TEXT,
    first_name TEXT,
    last_name TEXT, 
    gender_id INTEGER,
    dob DATE,
    email TEXT,
    phone TEXT,
    address TEXT,
    photo_path TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT
);

CREATE TABLE IF NOT EXISTS bank_account (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    balance INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS deposit_transaction (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    bank_account_id TEXT NOT NULL,
    amount INTEGER NOT NULL,
    status INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (bank_account_id) REFERENCES bank_account(account_id)
);
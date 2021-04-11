CREATE TABLE users
(
    id             SERIAL NOT NULL CONSTRAINT users_pk PRIMARY KEY,
    email          VARCHAR(256) NOT NULL,
    first_name     VARCHAR(256),
    last_name      VARCHAR(256),
    password       VARCHAR(256),
    password_reset VARCHAR(256),
    active         BOOLEAN DEFAULT TRUE,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     BOOLEAN,
    deleted_at     BOOLEAN
);

CREATE UNIQUE INDEX users_id_uindex ON users (id);

CREATE TABLE currencies
(
    id     SERIAL NOT NULL CONSTRAINT currencies_pk PRIMARY KEY,
    name   VARCHAR(256) NOT NULL,
    symbol VARCHAR(256)
);

CREATE UNIQUE INDEX currencies_id_uindex ON currencies (id);

CREATE TABLE accounts
(
    id             SERIAL NOT NULL CONSTRAINT accounts_pk PRIMARY KEY,
    user_id        INT NOT NULL,
    type_id        INT,
    currency_id    INT,
    name           VARCHAR(256) NOT NULL,
    initial_amount FLOAT DEFAULT 0,
    icon           VARCHAR(256) NOT NULL,
    color          VARCHAR(256) NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     BOOLEAN,
    deleted_at     BOOLEAN
);

CREATE UNIQUE INDEX accounts_id_uindex ON accounts (id);
CREATE INDEX accounts_type_id_index ON accounts (type_id);
CREATE INDEX accounts_user_id_index ON accounts (user_id);
CREATE INDEX accounts_currency_id_index ON accounts (currency_id);

CREATE TABLE account_users
(
    id         SERIAL NOT NULL CONSTRAINT account_users_pk PRIMARY KEY,
    account_id INT NOT NULL,
    user_id    INT NOT NULL
);

CREATE UNIQUE INDEX account_users_id_uindex ON account_users (id);
CREATE INDEX account_users_account_id_index ON account_users (account_id);
CREATE INDEX account_users_user_id_index ON account_users (user_id);

CREATE TABLE account_types
(
    id   SERIAL NOT NULL CONSTRAINT account_types_pk PRIMARY KEY,
    name VARCHAR(256) NOT NULL
);

CREATE UNIQUE INDEX account_types_id_uindex ON account_types (id);

CREATE TABLE transactions
(
    id                  SERIAL NOT NULL CONSTRAINT transactions_pk PRIMARY KEY,
    account_id          INT NOT NULL,
    user_id             INT NOT NULL,
    currency_id         INT NOT NULL,
    category_id         INT NOT NULL,
    transfer_account_id INT,
    amount              FLOAT NOT NULL,
    description         VARCHAR(256),
    type                VARCHAR(256) NOT NULL,
    location            VARCHAR(256),
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP
);

CREATE UNIQUE INDEX transactions_id_uindex ON transactions (id);
CREATE INDEX transactions_account_id_index ON transactions (account_id);
CREATE INDEX transactions_user_id_index ON transactions (user_id);
CREATE INDEX transactions_currency_id_index ON transactions (currency_id);
CREATE INDEX transactions_category_id_index ON transactions (category_id);
CREATE INDEX transactions_transfer_account_id_index ON transactions (transfer_account_id);

CREATE TABLE categories
(
    id      SERIAL NOT NULL CONSTRAINT categories_pk PRIMARY KEY,
    user_id INT,
    type_id INT,
    name    VARCHAR(256) NOT NULL,
    icon    VARCHAR(256) NOT NULL,
    color   VARCHAR(256) NOT NULL
);

CREATE UNIQUE INDEX categories_id_uindex ON categories (id);
CREATE INDEX categories_type_id_index ON categories (type_id);
CREATE INDEX categories_user_id_index ON categories (user_id);

CREATE TABLE transaction_types
(
    id   SERIAL NOT NULL CONSTRAINT transaction_types_pk PRIMARY KEY,
    name VARCHAR(256) NOT NULL
);

CREATE UNIQUE INDEX transaction_types_id_uindex ON transaction_types (id);

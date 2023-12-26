CREATE TABLE IF NOT EXISTS account (
    account_id serial PRIMARY KEY,
    name varchar,
    email varchar
);

CREATE TABLE IF NOT EXISTS transaction (
    transaction_id serial PRIMARY KEY,
    date_transaction varchar,
    amount decimal(10,2) NOT NULL,
    input_timestamp timestamp default current_timestamp,
    filename varchar NOT NULL,
    account_id integer NOT NULL,
    FOREIGN KEY (account_id)
      REFERENCES account (account_id)
);

INSERT INTO account(name, email) VALUES ('<YOUR_NAME>','<YOUR_EMAIL>');
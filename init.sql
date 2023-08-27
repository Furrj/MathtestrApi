CREATE TABLE account_info (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(50)
);

INSERT INTO account_info (username, password)
VALUES ('Jackson', 'Pass');
CREATE DATABASE mathtestr;

\c mathtestr

CREATE TYPE role AS ENUM ('S', 'T', 'A');

CREATE TABLE user_info (
    user_id SERIAL PRIMARY KEY,
    first_name VARCHAR(32),
    last_name VARCHAR(16),
    username VARCHAR(16),
    password VARCHAR(16),
    role role,
    period SMALLINT,
    teacher INT
);

CREATE TABLE session_data(
    user_id INTEGER,
    session_key VARCHAR(36),
    expires BIGINT,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES user_info(user_id)
);

CREATE TABLE test_results(
    user_id INTEGER,
    score SMALLINT,
    min INTEGER,
    max INTEGER,
    question_count SMALLINT,
    operations VARCHAR(50),
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES user_info(user_id)
);

INSERT INTO user_info (first_name, last_name, username, password, role, period, teacher)
VALUES ('Jackson', 'Furr', 'Poemmys', 'password', 'S', 0, 2);

INSERT INTO session_data (user_id, session_key, expires)
VALUES (1, 'test_uuid', 1234);

INSERT INTO test_results(user_id, score, min, max, question_count, operations)
VALUES (1, 100, 1, 12, 5, 'Multiplication');

SELECT * FROM user_info;

SELECT * FROM session_data;

SELECT * FROM test_results;

SELECT * FROM user_info
INNER JOIN session_data
USING (user_id)
INNER JOIN test_results
USING (user_id);
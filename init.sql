CREATE DATABASE mathtestr;

\c mathtestr

CREATE TABLE user_info (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(16),
    password VARCHAR(16),
    first_name VARCHAR(16),
    last_name VARCHAR(16)
);

CREATE TABLE session_data(
    user_id INTEGER,
    uuid VARCHAR(36),
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

INSERT INTO user_info (username, password, first_name, last_name)
VALUES ('a', 'pad89!', 'Jackson', 'Furr');

SELECT * FROM user_info;

INSERT INTO session_data (user_id, uuid, expires)
VALUES (0, 'test_uuid', 1234);

SELECT * FROM session_data;

INSERT INTO test_results(user_id, score, min, max, question_count, operations)
VALUES (3, 100, 1, 12, 5, 'Multiplication');

SELECT * FROM user_info
INNER JOIN session_data
USING (user_id)
INNER JOIN test_results
USING (user_id);
DROP TABLE IF EXISTS student_info;
DROP TABLE IF EXISTS teacher_info;
DROP TABLE IF EXISTS test_results;
DROP TABLE IF EXISTS session_data;
DROP TABLE IF EXISTS user_info;
DROP TYPE IF EXISTS role;

CREATE TYPE role AS ENUM ('S', 'T', 'A');

CREATE TABLE user_info (
    user_id SERIAL PRIMARY KEY,
    first_name VARCHAR(32),
    last_name VARCHAR(16),
    username VARCHAR(16),
    password VARCHAR(16),
    role role
);

CREATE TABLE session_data(
    user_id INTEGER PRIMARY KEY,
    session_key VARCHAR(36),
    expires BIGINT,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES user_info(user_id)
);

CREATE TABLE teacher_info(
    user_id INTEGER PRIMARY KEY,
    periods SMALLINT,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES user_info(user_id)
);

CREATE TABLE student_info(
    user_id INTEGER PRIMARY KEY,
    teacher_id INTEGER,
    period SMALLINT,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES user_info(user_id),
    CONSTRAINT fk_teacher_id
        FOREIGN KEY (teacher_id)
            REFERENCES teacher_info(user_id)
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

INSERT INTO user_info (first_name, last_name, username, password, role)
VALUES ('Michelle', 'Furr', 'MFurr', 'password', 'T');

INSERT INTO teacher_info (user_id, periods)
VALUES (1, 8);

INSERT INTO user_info (first_name, last_name, username, password, role)
VALUES ('Jackson', 'Furr', 'Poemmys', 'password', 'A');

INSERT INTO user_info (first_name, last_name, username, password, role)
VALUES ('Thomas', 'Glenn', 'Tg3', 'password', 'S');

INSERT INTO student_info (user_id, teacher_id, period)
VALUES (3, 1, 1);

INSERT INTO session_data (user_id, session_key, expires)
VALUES (1, 'test_uuid', 1234);

INSERT INTO session_data (user_id, session_key, expires)
VALUES (2, 'test_uuid', 1234);

INSERT INTO session_data (user_id, session_key, expires)
VALUES (3, 'test_uuid', 1234);

INSERT INTO test_results(user_id, score, min, max, question_count, operations)
VALUES (3, 100, 1, 12, 5, 'Multiplication');
package dbHandler

// SEARCH

// QCheckIfUsernameExists takes in username and return username
const QCheckIfUsernameExists = `
	SELECT username FROM user_info WHERE username=$1
`

// QGetUserByUsername takes in username and returns all fields of AllUserData
const QGetUserByUsername = `
	Select user_id, username, password, first_name, last_name, role, period, teacher, session_key, expires
	FROM user_info
	NATURAL JOIN session_data
	WHERE username=$1
`

// QGetUserIDByUsername takes in username and returns user_id
const QGetUserIDByUsername = `
	SELECT user_id
	FROM user_info
	WHERE username=$1
`

// QGetSessionDataByUserID takes in user_id and returns session data
const QGetSessionDataByUserID = `
	SELECT user_id, session_key, expires
	FROM session_data
	WHERE user_id=$1
`

// QGetTestResultsByUsername takes in user_id and returns all test results
const QGetTestResultsByUserID = `
	SELECT user_id, score, min, max, question_count, operations
	FROM test_results
	WHERE user_id=$1
`

// INSERT

// EInsertUserInfo inserts all fields of AllUserData into user_info table
const EInsertUserInfo = `
	INSERT INTO user_info (username, password, first_name, last_name, role, period, teacher)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`

// EInsertSessionData inserts all fields of SessionData into session_data table
const EInsertSessionData = `
	INSERT INTO session_data (user_id, session_key, expires)
	VALUES ($1, $2, $3)
`

// EInsertTestResults inserts all fields of TestResults into test_results table
const EInsertTestResults = `
    INSERT INTO test_results (user_id, score, min, max, question_count, operations)
    VALUES ($1, $2, $3, $4, $5, $6)
`

// TESTING

// ECreateRole creates role enum for user_info table
const ECreateRole = `
	CREATE TYPE role AS ENUM ('S', 'T', 'A');
`

// EInitUserInfo contains SQL commands to create user_info table
const EInitUserInfo = `
	CREATE TABLE user_info (
		user_id SERIAL PRIMARY KEY,
		first_name VARCHAR(32),
		last_name VARCHAR(16),
		username VARCHAR(16),
		password VARCHAR(16),
		role role,
		period SMALLINT,
		teacher INT
	)
`

// EInitSessionData contains SQL commands to create session_data table
const EInitSessionData = `
	CREATE TABLE session_data(
		user_id INTEGER,
		session_key VARCHAR(36),
		expires BIGINT,
		CONSTRAINT fk_user_id
			FOREIGN KEY (user_id)
				REFERENCES user_info(user_id)
	)
`

// EInitTestResults contains SQL commands to create test_results table
const EInitTestResults = `
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
	)
`

// EDeleteAllSessionData drops session_data table
const EDeleteAllSessionData = `
	DROP TABLE session_data
`

// EDeleteAllTestResults drops test_results table
const EDeleteAllTestResults = `
	DROP TABLE test_results;
`

// EDeleteAllUserInfo drops user_info table, dependents: session_data and
// rest_results because of user_id
const EDeleteAllUserInfo = `
	DROP TABLE user_info;
`

// EDeleteRole deletes role type
const EDeleteRole = `
	DROP TYPE role;
`

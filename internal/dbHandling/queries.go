package dbHandling

// SEARCH
const QCheckIfUsernameExists = `
	SELECT username FROM user_info WHERE username=$1
`

const QGetUserByUsername = `
	Select user_id, username, password, first_name, last_name, uuid, expires
	FROM user_info
	NATURAL JOIN session_data
	WHERE username=$1
`

// INSERT
const EInsertUser = `
	INSERT INTO user_info (username, password, first_name, last_name)
	VALUES ($1, $2, $3, $4)
`

const EInsertSessionData = `
	INSERT INTO session_data (user_id, uuid, expires)
	VALUES ($1, $2, $3)
`

// TESTING
const EInitUserInfo = `
	CREATE TABLE user_info (
		user_id SERIAL PRIMARY KEY,
		username VARCHAR(16),
		password VARCHAR(16),
		first_name VARCHAR(16),
		last_name VARCHAR(16)
	)
`

const EInitSessionData = `
	CREATE TABLE session_data(
		user_id INTEGER,
		uuid VARCHAR(36),
		expires BIGINT,
		CONSTRAINT fk_user_id
			FOREIGN KEY (user_id)
				REFERENCES user_info(user_id)
	)
`

const EDeleteAllSessionData = `
	DROP TABLE session_data
`

const EDeleteAllUserInfo = `
	DROP TABLE user_info;
`

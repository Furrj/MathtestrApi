package dbHandlers

const QCheckIfUsernameExists = "SELECT username FROM user_info WHERE username=$1"
const QGetUserByUsername = `
	Select user_id, username, password, first_name, last_name, uuid, expires
	FROM user_info
	NATURAL JOIN session_data
	WHERE username=$1
`

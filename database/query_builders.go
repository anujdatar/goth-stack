package database

func GenerateQueryStr(tableName string, columns []string) string {
	sqlQuery := "INSERT INTO " + tableName + " ("
	for i, column := range columns {
		if i == len(columns)-1 {
			sqlQuery += column + ") VALUES ("
		} else {
			sqlQuery += column + ", "
		}
	}
	for i := 0; i < len(columns); i++ {
		if i == len(columns)-1 {
			sqlQuery += "?)"
		} else {
			sqlQuery += "?, "
		}
	}
	return sqlQuery
}

func NewUserQueryStr() string {

	sql_query := `
		INSERT INTO users (
			first_name,
			last_name,
			username,
			email,
			password,
			phone,
			role,
			account_state,
			incorrect_pwd_count,
			pwd_reset_flag,
			pwd_reset_code,
			invite_code,
			phone_verification,
			email_verification,
			created_at,
			updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	return sql_query
}

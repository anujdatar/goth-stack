package database

import (
	"database/sql"
	"log"
)

func CreateDbTables(db *sql.DB) {
	// Create user table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		username TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		phone TEXT UNIQUE NOT NULL,
		role TEXT NOT NULL DEFAULT 'user',
		account_state TEXT NOT NULL,
		incorrect_pwd_count INTEGER NOT NULL DEFAULT 0,
		pwd_reset_flag INTEGER NOT NULL DEFAULT 0,
		pwd_reset_code TEXT,
		invite_code TEXT,
		phone_verification TEXT,
		email_verification TEXT,
		created_at INTEGER DEFAULT (strftime('%s', 'now')),
		updated_at INTEGER DEFAULT (strftime('%s', 'now'))
	)`)
	if err != nil {
		log.Fatal("Error. Unable to create user table: ", err)
	}

	// Create subscriptions table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS subscriptions (
		subscription_id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		price INTEGER NOT NULL,
		deleted INTEGER NOT NULL DEFAULT 0,
		created_at INTEGER DEFAULT (strftime('%s', 'now')),
		updated_at INTEGER DEFAULT (strftime('%s', 'now'))
	)`)
	if err != nil {
		log.Fatal("Error. Unable to create subscriptions table: ", err)
	}

	// Create teams table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS teams (
		team_id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		owner_id INTEGER NOT NULL,
		subscription_id INTEGER NOT NULL,
		deleted INTEGER NOT NULL DEFAULT 0,
		created_at INTEGER DEFAULT (strftime('%s', 'now')),
		updated_at INTEGER DEFAULT (strftime('%s', 'now')),
		FOREIGN KEY (owner_id) REFERENCES users (user_id),
		FOREIGN KEY (subscription_id) REFERENCES subscriptions (subscription_id)
	)`)
	if err != nil {
		log.Fatal("Error. Unable to create teams table: ", err)
	}

	// Create projects table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS projects (
		project_id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		team_id INTEGER NOT NULL,
		owner_id INTEGER NOT NULL,
		deleted INTEGER NOT NULL DEFAULT 0,
		created_at INTEGER DEFAULT (strftime('%s', 'now')),
		updated_at INTEGER DEFAULT (strftime('%s', 'now')),
		FOREIGN KEY (team_id) REFERENCES teams (team_id),
		FOREIGN KEY (owner_id) REFERENCES users (user_id)
	)`)
	if err != nil {
		log.Fatal("Error. Unable to create projects table: ", err)
	}

	// Create team_users table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS team_users (
		team_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		PRIMARY KEY (team_id, user_id)
		FOREIGN KEY (team_id) REFERENCES teams (team_id)
		FOREIGN KEY (user_id) REFERENCES users (user_id)
	)`)
	if err != nil {
		log.Fatal("Error. Unable to create team_users table: ", err)
	}

	// Create team_admins table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS team_admins (
		team_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		PRIMARY KEY (team_id, user_id)
		FOREIGN KEY (team_id) REFERENCES teams (team_id)
		FOREIGN KEY (user_id) REFERENCES users (user_id)
	)`)
	if err != nil {
		log.Fatal("Error. Unable to create team_admins table: ", err)
	}

	// Create team_projects table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS team_projects (
		team_id INTEGER NOT NULL,
		project_id INTEGER NOT NULL,
		PRIMARY KEY (team_id, project_id)
		FOREIGN KEY (team_id) REFERENCES teams (team_id)
		FOREIGN KEY (project_id) REFERENCES projects (project_id)
	)`)
	if err != nil {
		log.Fatal("Error. Unable to create team_projects table: ", err)
	}

	// Create team_subscriptions table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS team_subscriptions (
	team_id INTEGER NOT NULL,
	subscription_id INTEGER NOT NULL,
	PRIMARY KEY (team_id, subscription_id)
	FOREIGN KEY (team_id) REFERENCES teams (team_id)
	FOREIGN KEY (subscription_id) REFERENCES subscriptions (subscription_id)
)`)
	if err != nil {
		log.Fatal("Error. Unable to create team_subscriptions table: ", err)
	}

	// Create project_users table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS project_users (
		project_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		PRIMARY KEY (project_id, user_id)
		FOREIGN KEY (project_id) REFERENCES projects (project_id)
		FOREIGN KEY (user_id) REFERENCES users (user_id)
	)`)
	if err != nil {
		log.Fatal("Error. Unable to create project_users table: ", err)
	}

	// Create project_admins table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS project_admins (
		project_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		PRIMARY KEY (project_id, user_id)
		FOREIGN KEY (project_id) REFERENCES projects (project_id)
		FOREIGN KEY (user_id) REFERENCES users (user_id)
	)`)
	if err != nil {
		log.Fatal("Error. Unable to create project_admins table: ", err)
	}
}

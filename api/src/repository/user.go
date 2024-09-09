package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) CreateUser(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastIdInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIdInserted), nil
}

func (repository users) GetAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := repository.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var user []models.User

	for lines.Next() {
		var u models.User

		if err = lines.Scan(&u.ID, &u.Name, &u.Nick, &u.Email, &u.CreateAt); err != nil {
			return nil, err
		}

		user = append(user, u)
	}

	return user, nil
}

func (repository users) GetById(userId uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, email, created_at from users where id = ?",
		userId,
	)

	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) UpdateUser(userId uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err

	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userId); err != nil {
		return err
	}

	return nil
}

func (repository users) DeleteUser(userId uint64) error {
	statement, err := repository.db.Prepare(
		"delete from users where id = ?",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId); err != nil {
		return err
	}

	return nil
}

func (repository users) GetUserByEmail(email string) (models.User, error) {

	line, err := repository.db.Query(
		"select id, password from users where email = ?",
		email,
	)

	if err != nil {
		return models.User{}, err
	}
	defer line.Close()
	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Follow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository users) UnFollow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository users) GetFollowersByUser(userId uint64) ([]models.User, error) {

	lines, err := repository.db.Query(`
		select u.id, u.name, u.nick,  u.email, u.created_at 
		from users u inner join followers f on u.id = f.follower_id where f.user_id = ?
	`, userId)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var follower models.User

		if err = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreateAt,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository users) GetFollowingByUser(userId uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
	select u.id, u.name, u.nick,  u.email, u.created_at 
	from users u inner join followers f on u.id = f.user_id where f.follower_id = ?
`, userId)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var follower models.User

		if err = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreateAt,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil

}

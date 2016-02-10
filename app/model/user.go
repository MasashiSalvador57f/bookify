package model

import "time"

type User struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	FacebookID  string    `db:"facebook_id"`
	Email       string    `db:"email"`
	AccessToken string    `db:"access_token"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (u *User) IsEmptyUser() bool {
	if u.ID != 0 {
		return false
	}

	return true
}

func (u *User) IsNameSet() bool {
	if u.Name != "" {
		return true
	}
	return false
}

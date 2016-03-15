package repository

import (
	"github.com/MasashiSalvador57f/bookify/app/model"

	"time"

	"gopkg.in/gorp.v1"
	"github.com/MasashiSalvador57f/bookify/app/shared/crypto"
)

type UserRepository interface {
	FindByID(int64) (*model.User, error)
	CreateByFacebookIDAndEmail(string, string) (*model.User, error)
	FindByFacebookID(string) (*model.User, error)
}

type UserRepo struct {
	UserRepository
	dbmap *gorp.DbMap
	now   time.Time
}

func NewUserRepo(dbm *gorp.DbMap, now time.Time) UserRepo {
	return UserRepo{
		dbmap: dbm,
		now:   now,
	}
}

func (ur *UserRepo) FindByID(id int64) (*model.User, error) {
	u := new(model.User)
	err := ur.dbmap.SelectOne(u, "SELECT * FROM user WHERE id = ?", id)

	return u, err
}

func (ur *UserRepo) CreateByFacebookIDAndEmail(facebookID, email string) (*model.User, error) {
	u := new(model.User)
	u.Email = email
	u.FacebookID = facebookID
	u.Name = ""
	u.AccessToken = accesstoken.ConvertToAccessToken(u.Email)
	u.CreatedAt = ur.now
	u.UpdatedAt = ur.now

	err := ur.dbmap.Insert(u)

	return u, err
}

func (ur *UserRepo) FindByFacebookID(facebookID string) (*model.User, error) {
	u := new(model.User)

	err := ur.dbmap.SelectOne(u, "SELECT * FROM user WHERE facebook_id = ?", facebookID)

	return u, err
}


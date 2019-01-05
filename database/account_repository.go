package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/flohero/Spongebot/database/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

const JWT_PASSWORD string = "JWT_PASSWORD"

func (p *Persistence) IsValid(acc *model.Account) (bool, error) {
	if p.FindAccountByUsername(acc.Username).Username != "" {
		return false, errors.New(fmt.Sprint("User not found!"))
	}
	return true, nil
}

func (p *Persistence) CreateAccount(acc *model.Account) (error, *model.Account) {
	if ok, err := p.IsValid(acc); !ok {
		return err, nil
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	acc.Password = string(hashedPassword)
	p.db.Create(acc)
	if acc.Id <= 0 {
		return errors.New("Failed to create account"), nil
	}
	acc.Password = "" //delete password
	return nil, acc
}

func (p *Persistence) Login(username, password string) (error, *model.Account) {

	account := &model.Account{}
	if account = p.FindAccountByUsername(username); account.Username == "" {
		return errors.New("Username not found"), nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return errors.New("Invalid login credentials. Please try again"), nil
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &model.Token{
		UserId: account.Id,
		Admin:  account.Admin,
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv(JWT_PASSWORD)))
	account.Token = tokenString //Store the token in the response

	return nil, account
}

func (p *Persistence) FindAccountByUsername(username string) (acc *model.Account) {
	acc = &model.Account{}
	p.db.Where(&model.Account{Username: username}).First(acc)
	acc.Token = ""
	return acc
}

func (p *Persistence) FindAllAccounts() ([]*model.Account, error) {
	rows, err := p.db.Table("accounts").Select("id, username, admin").Rows()
	if err != nil {
		return nil, errors.New("Error while getting all accounts")
	}
	return assignRowsToAccount(rows)
}

func (p *Persistence) FindAccountById(id int) *model.Account {
	acc := &model.Account{}
	p.db.Where(model.Account{Id: id}).First(acc)
	return acc
}

func assignRowsToAccount(rows *sql.Rows) ([]*model.Account, error) {
	accounts := make([]*model.Account, 0)
	for rows.Next() {
		acc := &model.Account{}
		if err := rows.Scan(&acc.Id, &acc.Username, &acc.Admin); err != nil {
			log.Fatal(err.Error())
			return nil, errors.New("Error with DB")
		}
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

func (p *Persistence) DeleteAccountById(id int) {
	p.db.Where(model.Account{Id: id}).Delete(model.Account{})
}

func (p *Persistence) UpdatePasswordById(id int, password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	p.db.Model(&model.Account{}).Updates(model.Account{Password: string(hashedPassword)})
}

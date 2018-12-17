package database

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/flohero/Spongebot/database/model"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
)

func (p *Persistence) IsValid(acc *model.Account) (bool, error) {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(acc.Email) {
		return false, errors.New("Not a valid email")
	}
	if p.FindByEmail(acc.Email).Email != "" {
		return false, errors.New("")
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
	tk := &model.Token{UserId: acc.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	acc.Token = tokenString

	acc.Password = "" //delete password
	return nil, acc
}

func (p *Persistence) Login(email, password string) (error, *model.Account) {

	account := &model.Account{}
	if p.FindByEmail(email).Email == "" {
		return errors.New("Email not found"), nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return errors.New("Invalid login credentials. Please try again"), nil
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &model.Token{UserId: account.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	return nil, account
}

func (p *Persistence) FindByEmail(email string) (acc *model.Account) {
	acc = &model.Account{}
	p.db.Where(&model.Account{Email: email}).First(acc)
	return acc
}

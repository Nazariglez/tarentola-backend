// Created by nazarigonzalez on 7/10/17.

package usermodel

import (
  "github.com/nazariglez/tarentola-backend/database/helpers"
  "golang.org/x/crypto/bcrypt"
  "errors"
  "fmt"
)

func DeleteByID(id uint) error {
  err := db.Select("id").Where("id = ?", id).First(&User{}).Error
  if helpers.IsNotFoundErr(err) {
      return errors.New("Invalid user id.")
  }

  if err != nil {
    return err
  }

  return db.Where("id = ?", id).Delete(&User{}).Error
}

func GetByID(id uint) (*User, error) {
  um := &User{}
  err := db.Preload("Role").Where("id = ?", id).Find(um).Error
  if err != nil {
    return nil, err
  }
  return um, nil
}

func FindOne(u *User) error {
  return db.Where(*u).First(&u).Error
}

func FindToLogin(email, password string) (*User, error) {
  u := User{}

  if err := db.Preload("Role").Select("id, email, hashed_password, role_refer").Where(map[string]interface{}{
    "email": email,
  }).First(&u).Error; err != nil {
    fmt.Println("here1",err)
    if helpers.IsNotFoundErr(err) {
      return nil, nil
    }

    return nil, err
  }

  fmt.Printf("match %+v\n", u)
  if !matchPassword(password, u.HashedPassword) {
    fmt.Println("not match", password, u.HashedPassword)
    return nil, errors.New("Invalid email or password.")
  }

  return &u, nil
}

func Create(u *User) error {
  fmt.Printf("create: %+v\n", u)
  return db.Create(u).Error
}

func ExistsEmail(email string) (bool, error) {
  u := User{Email: email}
  err := db.Select("ID").Where(u).First(&u).Error
  if err != nil {
    if helpers.IsNotFoundErr(err) {
      return false, nil
    }

    return false, err
  }

  return true, nil
}

//--
func hashPassword(pass string) (string, error) {
  p, err := bcrypt.GenerateFromPassword([]byte(pass), 11)
  if err != nil {
    return "", err
  }

  return string(p), nil
}

func matchPassword(password string, hashedPass string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
  if err != nil {
    return false
  }

  return true
}


package handlers

import (
	"crypto/sha256"
	"fmt"
	"isogate/pkg/models"
	"net/http"
	"time"

	"github.com/satori/go.uuid"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RegisterSessionEndpoints registers session endpoints
func RegisterSessionEndpoints(e *echo.Group) {
	e.POST("/signup", signup)
	e.POST("/login", login)
	e.POST("/logout", logout)
}

func signup(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.Account{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}

	existing := models.Account{}
	err = db.C("Accounts").FindId(a.Username).One(&existing)

	crypt := sha256.New()
	crypt.Write([]byte(a.Password))
	a.Password = string(crypt.Sum(nil))

	if err == nil {
		a.ID = existing.ID
		_, err = db.C("Accounts").UpsertId(existing.ID, &a)
		if err != nil {
			return err
		}
	} else {
		if a.ID == "" {
			a.ID = bson.NewObjectId()
		}

		err = db.C("Accounts").Insert(&a)
		if err != nil {
			return err
		}
	}

	return e.NoContent(http.StatusOK)
}

func login(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.Account{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}

	crypt := sha256.New()
	crypt.Write([]byte(a.Password))
	a.Password = string(crypt.Sum(nil))

	existing := models.Account{}
	err = db.C("Accounts").Find(bson.M{"username": a.Username, "password": a.Password}).One(&existing)

	if err != nil {
		return e.String(http.StatusUnauthorized, "Invalid user details")
	}

	s := models.Session{
		ID:              bson.NewObjectId(),
		Username:        existing.Username,
		CreatedDateTime: time.Now().UTC(),
		ExpiryDateTime:  time.Now().UTC().Add(models.SessionExpiry),
		Token:           uuid.NewV4().String(),
	}

	_, err = db.C("Sessions").Upsert(bson.M{"username": s.Username}, &s)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, &s)
}

func logout(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	s := models.Session{}
	err := e.Bind(&s)
	if err != nil {
		return err
	}

	existing := models.Session{}
	err = db.C("Session").Find(bson.M{"Token": s.Token}).One(&existing)

	if err != nil {
		return e.String(http.StatusNotFound, "Invalid user details")
	}

	err = db.C("Sessions").RemoveId(existing.ID)
	if err != nil {
		return err
	}

	return e.NoContent(http.StatusOK)
}

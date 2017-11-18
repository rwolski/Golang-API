package main

import (
	"fmt"
	"isogate/pkg/handlers"
	"isogate/pkg/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer db.Close() // clean up when we’re done

	e := echo.New()
	e.Use(attachMongoContext(db))

	session := e.Group("/session")
	api := e.Group("/api", checkSession())

	handlers.RegisterSessionEndpoints(session)

	handlers.RegisterLicenseEndpoints(api)
	handlers.RegisterBillingEndpoints(api)

	handlers.RegisterSiteEndpoints(api)
	handlers.RegisterGroupEndpoints(api)
	handlers.RegisterAdminEndpoints(api)
	handlers.RegisterSubjectEndpoints(api)

	handlers.RegisterBalanceStandardTestEndpoints(api)

	handlers.RegisterChoiceStandardTestEndpoints(api)
	handlers.RegisterChoiceMazeTestEndpoints(api)

	e.Logger.Fatal(e.Start(":8080"))
}

func attachMongoContext(db *mgo.Session) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			s := db.Copy()
			defer s.Close()

			db := s.DB("IsoGate")

			e.Set("database", db)
			if err := next(e); err != nil {
				e.Error(err)
			}
			return nil
		}
	}
}

func checkSession() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			db := e.Get("database").(*mgo.Database)
			if db == nil {
				return fmt.Errorf("Bad database session")
			}

			token := e.Request().Header.Get("X-CSRF-Token")
			if token == "" {
				return e.String(http.StatusUnauthorized, "No token provided")
			}

			existing := models.Session{}
			err := db.C("Sessions").Find(bson.M{"token": token}).One(&existing)
			if err != nil {
				return e.String(http.StatusUnauthorized, "Incorrect token provided")
			}

			currentTime := time.Now().UTC()
			if currentTime.After(existing.ExpiryDateTime) {
				return e.String(http.StatusUnauthorized, "Token has expired")
			}

			if err := next(e); err != nil {
				e.Error(err)
			}
			return nil
		}
	}
}

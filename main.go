package main

import (
	"isogate/pkg/handlers"
	"log"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer db.Close() // clean up when we’re done

	e := echo.New()
	e.Use(attachMongoContext(db))

	handlers.RegisterLicenseEndpoints(e)
	handlers.RegisterBillingEndpoints(e)

	handlers.RegisterSiteEndpoints(e)
	handlers.RegisterGroupEndpoints(e)
	handlers.RegisterAdminEndpoints(e)
	handlers.RegisterSubjectEndpoints(e)

	handlers.RegisterChoiceStandardTestEndpoints(e)
	handlers.RegisterChoiceMazeTestEndpoints(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func attachMongoContext(db *mgo.Session) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			s := db.Copy()
			defer s.Close()

			db := s.DB("IsoGate")

			c.Set("database", db)
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

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

	session := e.Group("/user")
	api := e.Group("") //, checkSession())

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

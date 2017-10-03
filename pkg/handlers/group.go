package handlers

import (
	"fmt"
	"isogate/pkg/models"
	"net/http"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func RegisterGroupEndpoints(e *echo.Echo) {
	g := e.Group("/groups")
	g.POST("", addGroup)
	g.PUT("/:id", updateGroup)
	g.GET("/:id", getGroup)
	g.GET("", getGroups)
}

func addGroup(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.Group{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}
	if a.GroupID == "" {
		a.GroupID = bson.NewObjectId()
	}

	_, err = db.C("Groups").UpsertId(a.GroupID, &a)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, a)
}

func updateGroup(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.Group{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}
	if a.GroupID == "" {
		return e.NoContent(http.StatusNotFound)
	}

	_, err = db.C("Groups").UpsertId(a.GroupID, &a)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, a)
}

func getGroup(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	id := bson.ObjectIdHex(e.Param("id"))
	if !id.Valid() {
		return fmt.Errorf("Bad parameters")
	}

	a := models.Group{}
	err := db.C("Groups").FindId(id).One(&a)
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, a)
}

func getGroups(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := []models.Group{}
	db.C("Groups").Find(nil).Sort("groupName").Limit(100).All(&a)
	return e.JSON(http.StatusOK, a)
}

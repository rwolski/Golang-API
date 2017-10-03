package handlers

import (
	"fmt"
	"isogate/pkg/models"
	"net/http"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func AddAdmin(e echo.Context) error {
	i := e.Get("database")
	db := i.(*mgo.Session)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.Admin{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}
	if a.AdminID == "" {
		a.AdminID = bson.NewObjectId()
	}

	_, err = db.DB("IsoGate").C("Admins").UpsertId(a.AdminID, &a)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, a)
}

func UpdateAdmin(e echo.Context) error {
	i := e.Get("database")
	db := i.(*mgo.Session)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.Admin{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}
	if a.AdminID == "" {
		return e.NoContent(http.StatusNotFound)
	}

	_, err = db.DB("IsoGate").C("Admins").UpsertId(a.AdminID, &a)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, a)
}

func GetAdmin(e echo.Context) error {
	i := e.Get("database")
	db := i.(*mgo.Session)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	id := bson.ObjectIdHex(e.Param("id"))
	if !id.Valid() {
		return fmt.Errorf("Bad parameters")
	}

	a := models.Admin{}
	err := db.DB("IsoGate").C("Admins").FindId(id).One(&a)
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, a)
}

func GetAdmins(e echo.Context) error {
	i := e.Get("database")
	db := i.(*mgo.Session)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := []models.Admin{}
	db.DB("IsoGate").C("Admins").Find(nil).Sort("groupName").Limit(100).All(&a)
	return e.JSON(http.StatusOK, a)
}

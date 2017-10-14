package handlers

import (
	"fmt"
	"isogate/pkg/models"
	"net/http"
	"time"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RegisterAdminEndpoints API registration
func RegisterAdminEndpoints(e *echo.Echo) {
	e.POST("/admin", saveAdmin)
	e.GET("/admin", getAdmin)
	e.GET("/admins", getAdmins)
}

func saveAdmin(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.Admin{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}
	a.ServerUpdateDateTime = time.Now().UTC()

	existing := models.Admin{}
	err = db.C("Admins").Find(bson.M{"adminUuid": a.AdminUUID}).One(&existing)

	if err == nil {
		if existing.ServerUpdateDateTime.After(a.LocalUpdateDateTime) {
			// Server version is more recent
			return e.JSON(http.StatusConflict, existing)
		}

		a.AdminID = existing.AdminID
		_, err = db.C("Admins").UpsertId(existing.AdminID, &a)
		if err != nil {
			return err
		}
	} else {
		if a.AdminID == "" {
			a.AdminID = bson.NewObjectId()
		}

		err = db.C("Admins").Insert(&a)
		if err != nil {
			return err
		}
	}

	return e.JSON(http.StatusOK, a)
}

func getAdmin(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	var id bson.ObjectId
	if idParam := e.QueryParam("id"); idParam != "" && bson.IsObjectIdHex(idParam) {
		id = bson.ObjectIdHex(idParam)
	}
	uuid, err := uuid.FromString(e.QueryParam("uuid"))
	if !id.Valid() && err != nil {
		return fmt.Errorf("Bad parameters")
	}

	a := models.Admin{}
	if id.Valid() {
		err = db.C("Admins").FindId(id).One(&a)
	} else {
		err = db.C("Admins").Find(bson.M{"adminUuid": uuid}).One(&a)
	}
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, a)
}

func getAdmins(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	var a []*models.Admin
	err := db.C("Admins").Find(nil).All(&a)
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, a)
}

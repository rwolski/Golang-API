package handlers

import (
	"fmt"
	"isogate/pkg/models"
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RegisterAdminEndpoints API registration
func RegisterAdminEndpoints(e *echo.Echo) {
	g := e.Group("/admin")
	g.POST("", saveAdmin)
	g.GET("", getAdmin)
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

	existing := models.Admin{}
	err = db.C("Admins").Find(bson.M{"adminUuid": g.AdminUUID}).One(&existing)

	if err == nil {
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

func getGroup(e echo.Context) error {
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
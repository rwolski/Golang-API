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

// RegisterGroupEndpoints API registration
func RegisterGroupEndpoints(e *echo.Echo) {
	g := e.Group("/groups")
	g.POST("", saveGroup)
	g.GET("", getGroup)
}

func saveGroup(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	g := models.Group{}
	err := e.Bind(&g)
	if err != nil {
		return err
	}

	existing := models.License{}
	err = db.C("Groups").Find(bson.M{"groupUuid": g.GroupUUID}).One(&existing)

	if err == nil {
		g.GroupID = existing.GroupID
		_, err = db.C("Groups").UpsertId(existing.GroupID, &g)
		if err != nil {
			return err
		}
	} else {
		if g.GroupID == "" {
			g.GroupID = bson.NewObjectId()
		}

		err = db.C("Groups").Insert(&g)
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

	g := models.Group{}
	if id.Valid() {
		err = db.C("Groups").FindId(id).One(&a)
	} else {
		err = db.C("Groups").Find(bson.M{"groupUuid": uuid}).One(&a)
	}
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, g)
}

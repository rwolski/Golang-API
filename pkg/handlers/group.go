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

// RegisterGroupEndpoints API registration
func RegisterGroupEndpoints(e *echo.Echo) {
	e.POST("/group", saveGroup)
	e.GET("/group", getGroup)
	e.GET("/groups", getGroups)
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
	g.ServerUpdateDateTime = time.Now().UTC()

	existing := models.Group{}
	err = db.C("Groups").Find(bson.M{"groupUuid": g.GroupUUID}).One(&existing)

	if err == nil {
		fmt.Printf("Found group: %+v", existing)

		if existing.ServerUpdateDateTime.After(g.LocalUpdateDateTime) {
			// Server version is more recent
			fmt.Printf("Group is out of date, returning")
			return e.JSON(http.StatusConflict, existing)
		}

		g.GroupID = existing.GroupID
		_, err = db.C("Groups").UpsertId(existing.GroupID, &g)
		if err != nil {
			return err
		}

		fmt.Printf("Updated group: %+v", g)
	} else {

		if g.GroupID == "" {
			g.GroupID = bson.NewObjectId()
		}

		fmt.Printf("New group: %+v", g)

		err = db.C("Groups").Insert(&g)
		if err != nil {
			return err
		}
	}

	return e.JSON(http.StatusOK, g)
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
		err = db.C("Groups").FindId(id).One(&g)
	} else {
		err = db.C("Groups").Find(bson.M{"groupUuid": uuid}).One(&g)
	}
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, g)
}

func getGroups(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	g := models.Groups{}
	err := db.C("Groups").Find(nil).All(&g.Groups)
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, g)
}

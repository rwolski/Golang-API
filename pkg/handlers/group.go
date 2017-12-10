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
func RegisterGroupEndpoints(e *echo.Group) {
	e.POST("/group", saveGroup, checkSession())
	e.GET("/group", getGroup, checkSession())
	e.GET("/groups", getGroups, checkSession())
}

// swagger:route POST /group SaveGroup
//
// Creates or updates a subject group.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Schemes: http, https
// Responses:
// 	200: GroupResponse
//  401: HttpResponse
//  409: GroupResponse A more recent version is available
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
		if existing.ServerUpdateDateTime.After(g.LocalUpdateDateTime) {
			// Server version is more recent
			return e.JSON(http.StatusConflict, existing)
		}

		g.GroupID = existing.GroupID
		_, err = db.C("Groups").UpsertId(existing.GroupID, &g)
		if err != nil {
			return err
		}
	} else {
		if g.GroupID == "" {
			g.GroupID = bson.NewObjectId().String()
		}

		err = db.C("Groups").Insert(&g)
		if err != nil {
			return err
		}
	}

	return e.JSON(http.StatusOK, g)
}

// swagger:route GET /group GetGroup
//
// Gets a subject group by it's UUID.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Schemes: http, https
// Responses:
// 	200: GroupResponse
//  401: HttpResponse
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

// swagger:route GET /groups GetGroups
//
// Gets all subject groups for a given Site UUID.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Schemes: http, https
// Responses:
// 	200: GroupsResponse
//  401: HttpResponse
func getGroups(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	// uuid, err := uuid.FromString(e.QueryParam("siteUuid"))
	// if err != nil {
	// 	return fmt.Errorf("Bad parameters")
	// }

	g := models.Groups{
		Groups: []models.Group{},
	}

	err := db.C("Groups").Find(nil).All(&g.Groups)
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, g)
}

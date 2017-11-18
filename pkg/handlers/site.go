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

// RegisterSiteEndpoints API registration
func RegisterSiteEndpoints(e *echo.Group) {
	g := e.Group("/site")
	g.POST("", saveSite)
	g.GET("", getSite)
}

func saveSite(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	s := models.Site{}
	err := e.Bind(&s)
	if err != nil {
		return err
	}
	s.ServerUpdateDateTime = time.Now().UTC()

	existing := models.Site{}
	err = db.C("Sites").Find(bson.M{"siteUuid": s.SiteUUID}).One(&existing)

	if err == nil {
		//fmt.Printf("Found site: %+v", existing)

		if existing.ServerUpdateDateTime.After(s.LocalUpdateDateTime) {
			// Server version is more recent
			//fmt.Printf("Subject is out of date, returning")
			return e.JSON(http.StatusConflict, existing)
		}

		s.SiteID = existing.SiteID
		_, err = db.C("Sites").UpsertId(existing.SiteID, &s)
		if err != nil {
			return err
		}
		//fmt.Printf("Updated site: %+v", s)
	} else {

		if s.SiteID == "" {
			s.SiteID = bson.NewObjectId()
		}

		//fmt.Printf("New site: %+v", s)

		err = db.C("Sites").Insert(&s)
		if err != nil {
			return err
		}
	}

	return e.JSON(http.StatusOK, s)
}

func getSite(e echo.Context) error {
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

	s := models.Site{}
	if id.Valid() {
		err = db.C("Sites").FindId(id).One(&s)
	} else {
		err = db.C("Sites").Find(bson.M{"siteUuid": uuid}).One(&s)
	}
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, s)
}

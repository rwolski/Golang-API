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

// RegisterSubjectEndpoints API registration
func RegisterSubjectEndpoints(e *echo.Echo) {
	e.POST("/subject", saveSubject)
	e.GET("/subject", getSubject)
	e.GET("/subjects", getSubjects)
}

func saveSubject(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	s := models.Subject{}
	err := e.Bind(&s)
	if err != nil {
		return err
	}
	s.ServerUpdateDateTime = time.Now().UTC()

	existing := models.Subject{}
	err = db.C("Subjects").Find(bson.M{"subjectUuid": s.SubjectUUID}).One(&existing)

	if err == nil {
		fmt.Printf("Found subject: %+v", existing)

		if existing.ServerUpdateDateTime.After(s.LocalUpdateDateTime) {
			// Server version is more recent
			fmt.Printf("Subject is out of date, returning")
			return e.JSON(http.StatusConflict, existing)
		}

		s.SubjectID = existing.SubjectID
		_, err = db.C("Subjects").UpsertId(existing.SubjectID, &s)
		if err != nil {
			return err
		}
		fmt.Printf("Updated subject: %+v", s)
	} else {
		if s.SubjectID == "" {
			s.SubjectID = bson.NewObjectId()
		}

		fmt.Printf("New subject: %+v", s)

		err = db.C("Subjects").Insert(&s)
		if err != nil {
			return err
		}
	}

	return e.JSON(http.StatusOK, s)
}

func getSubject(e echo.Context) error {
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

	s := models.Subject{}
	if id.Valid() {
		err = db.C("Subjects").FindId(id).One(&s)
	} else {
		err = db.C("Subjects").Find(bson.M{"subjectUuid": uuid}).One(&s)
	}
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, s)
}

func getSubjects(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	uuid, err := uuid.FromString(e.QueryParam("siteUuid"))
	if err != nil {
		return fmt.Errorf("Bad parameters")
	}

	s := models.Subjects{
		Subjects: []models.Subject{},
	}

	err = db.C("Subjects").Find(bson.M{"subjectSiteUuid": uuid}).All(&s.Subjects)
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, s)
}

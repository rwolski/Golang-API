package handlers

import (
	"fmt"
	"isogate/pkg/models"
	"net/http"
	"time"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RegisterChoiceMazeTestEndpoints API registration
func RegisterChoiceMazeTestEndpoints(e *echo.Echo) {
	e.POST("/choice/test/maze", saveMazeSession)
	//e.GET("/choice/test/maze", getMazeSession)
}

func saveMazeSession(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	s := models.ChoiceMazeSession{}
	err := e.Bind(&s)
	if err != nil {
		return err
	}
	s.ServerUpdateDateTime = time.Now().UTC()

	existing := models.ChoiceSession{}
	err = db.C("ChoiceSessions").Find(bson.M{"sessionUuid": s.SessionUUID}).One(&existing)

	if err == nil {
		return e.NoContent(http.StatusConflict)
	}

	if s.SessionID == "" {
		s.SessionID = bson.NewObjectId()
	}

	err = db.C("ChoiceSessions").Insert(&s.ChoiceSession)
	if err != nil {
		return err
	}

	err = saveMazeData(db, s.Tests)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, s)
}

func saveMazeData(db *mgo.Database, tests []models.ChoiceMazeTest) error {
	models := make([]interface{}, len(tests))
	for i := 0; i < len(tests); i++ {
		tests[i].TestID = bson.NewObjectId()
		tests[i].ServerUpdateDateTime = time.Now().UTC()
		models[i] = tests[i]
	}

	err := db.C("ChoiceMazeTests").Insert(models...)
	if err != nil {
		return err
	}
	return nil
}

// func getMazeSession(e echo.Context) error {
// 	db := e.Get("database").(*mgo.Database)
// 	if db == nil {
// 		return fmt.Errorf("Bad database session")
// 	}

// 	var id bson.ObjectId
// 	if idParam := e.QueryParam("id"); idParam != "" && bson.IsObjectIdHex(idParam) {
// 		id = bson.ObjectIdHex(idParam)
// 	}
// 	uuid, err := uuid.FromString(e.QueryParam("uuid"))
// 	if !id.Valid() && err != nil {
// 		return fmt.Errorf("Bad parameters")
// 	}

// 	s := models.ChoiceSession{}
// 	if id.Valid() {
// 		err = db.C("ChoiceSessions").FindId(id).One(&s)
// 	} else {
// 		err = db.C("ChoiceSessions").Find(bson.M{"adminUuid": uuid}).One(&s)
// 	}
// 	if err != nil {
// 		return e.NoContent(http.StatusNotFound)
// 	}
// 	return e.JSON(http.StatusOK, s)
// }

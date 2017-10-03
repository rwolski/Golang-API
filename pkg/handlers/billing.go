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

func RegisterBillingEndpoints(e *echo.Echo) {
	b := e.Group("/billing")
	b.POST("/account", saveAccount)
	b.GET("/account", getAccount)
	b.POST("/event", saveEvent)
	b.GET("/event", getEvent)
}

func saveAccount(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.BillingAccount{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}

	if a.BillingAccountID == "" {
		a.BillingAccountID = bson.NewObjectId()
	}

	_, err = db.C("BillingAccounts").Upsert(bson.M{"billingAccountUuid": a.BillingAccountUUID}, &a)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, a)
}

func getAccount(e echo.Context) error {
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

	a := models.BillingAccount{}
	if id.Valid() {
		err = db.C("BillingAccounts").FindId(id).One(&a)
	} else {
		err = db.C("BillingAccounts").Find(bson.M{"billingAccountUuid": uuid}).One(&a)
	}
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, a)
}

func saveEvent(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.BillingEvent{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}

	if a.BillingEventID == "" {
		a.BillingEventID = bson.NewObjectId()
	}

	_, err = db.C("BillingEvents").Upsert(bson.M{"billingEventUuid": a.BillingEventUUID}, &a)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, a)
}

func getEvent(e echo.Context) error {
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

	a := models.BillingEvent{}
	if id.Valid() {
		err = db.C("BillingEvents").FindId(id).One(&a)
	} else {
		err = db.C("BillingEvents").Find(bson.M{"billingEventUuid": uuid}).One(&a)
	}
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, a)
}

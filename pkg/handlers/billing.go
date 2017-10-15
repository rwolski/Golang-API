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

// RegisterBillingEndpoints registers endpoints
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
	a.ServerUpdateDateTime = time.Now().UTC()

	existing := models.BillingAccount{}
	err = db.C("BillingAccounts").Find(bson.M{"billingAccountUuid": a.BillingAccountUUID}).One(&existing)

	if err == nil {
		fmt.Printf("Found bill account: %+v", existing)

		if existing.ServerUpdateDateTime.After(a.LocalUpdateDateTime) {
			// Server version is more recent
			fmt.Printf("Billing account is out of date, returning")
			return e.JSON(http.StatusConflict, existing)
		}

		a.BillingAccountID = existing.BillingAccountID
		_, err = db.C("BillingAccounts").UpsertId(existing.BillingAccountID, &a)
		if err != nil {
			return err
		}

		fmt.Printf("Updated billing account: %+v", a)
	} else {
		if a.BillingAccountID == "" {
			a.BillingAccountID = bson.NewObjectId()
		}

		fmt.Printf("New billing account: %+v", a)

		err = db.C("BillingAccounts").Insert(&a)
		if err != nil {
			return err
		}
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
	billingID, billingErr := uuid.FromString(e.QueryParam("uuid"))
	adminID, adminErr := uuid.FromString(e.QueryParam("admin"))
	if !id.Valid() && (billingErr != nil || adminErr != nil) {
		return fmt.Errorf("Bad parameters")
	}

	a := models.BillingAccount{}
	var err error
	if id.Valid() {
		err = db.C("BillingAccounts").FindId(id).One(&a)
	} else {
		err = db.C("BillingAccounts").Find(bson.M{"billingAccountUuid": billingID}).One(&a)
	}
	if err != nil {
		fmt.Println("doesnt exist")
		// Create a billing account then
		u := models.Admin{}
		err = db.C("Admins").Find(bson.M{"adminUuid": adminID}).One(&u)
		if err != nil {
			return err
		}

		// Default to 20 credits if master admin (so 1 account per system only)
		if u.AdminStatus == 2 {
			a.BillingCredits = 20
		}
		a.BillingAccountID = bson.NewObjectId()
		a.BillingAccountUUID = billingID
		a.BillingAdminUUID = adminID
		err = db.C("BillingAccounts").Insert(&a)
		if err != nil {
			return err
		}
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
	if a.SessionStatus == 1 && a.ServerStartDateTime.IsZero() {
		a.ServerStartDateTime = time.Now().UTC()
	}
	if a.SessionStatus == 2 && a.ServerEndDateTime.IsZero() {
		a.ServerEndDateTime = time.Now().UTC()
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

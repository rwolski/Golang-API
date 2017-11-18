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
func RegisterBillingEndpoints(e *echo.Group) {
	b := e.Group("/billing")
	b.POST("/account", saveAccount)
	b.GET("/account", getAccount)
	b.GET("/accounts", getAccounts)
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
		//fmt.Printf("Found bill account: %+v", existing)

		if existing.ServerUpdateDateTime.After(a.LocalUpdateDateTime) {
			// Server version is more recent
			//fmt.Printf("Billing account is out of date, returning")
			return e.JSON(http.StatusConflict, existing)
		}

		a.BillingAccountID = existing.BillingAccountID
		_, err = db.C("BillingAccounts").UpsertId(existing.BillingAccountID, &a)
		if err != nil {
			return err
		}

		//fmt.Printf("Updated billing account: %+v", a)
	} else {
		if a.BillingAccountID == "" {
			a.BillingAccountID = bson.NewObjectId()
		}

		// Default to 20 credits if master admin (so 1 trial account per system only)
		if a.TrialAccount {
			a.BillingCredits = 20
		}

		//fmt.Printf("New billing account: %+v", a)

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

func getAccounts(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	uuid, err := uuid.FromString(e.QueryParam("siteUuid"))
	if err != nil {
		return fmt.Errorf("Bad parameters")
	}

	a := models.BillingAccounts{}
	err = db.C("BillingAccounts").Find(bson.M{"billingSiteUuid": uuid}).All(&a.Accounts)
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

	b := models.BillingEvent{}
	err := e.Bind(&b)
	if err != nil {
		return err
	}

	// Check billing credits first
	// and reject if not enough
	a := models.BillingAccount{}
	err = db.C("BillingAccounts").Find(bson.M{"billingAccountUuid": b.BillingAccountUUID}).One(&a)
	if err != nil {
		b.SessionStatus = models.SessionRejected
		return e.JSON(http.StatusForbidden, b)
	}

	eventCharge := 0 // change this to reflect cost per test type
	// if b.SessionType == 0 {
	// 	eventCharge = 10
	// }

	if a.BillingCredits-eventCharge < 0 {
		b.SessionStatus = models.SessionRejected
		return e.JSON(http.StatusUnauthorized, b)
	}

	// Start test event
	if b.SessionStatus == models.SessionStarted && b.ServerStartDateTime.IsZero() {
		b.ServerStartDateTime = time.Now().UTC()
	}

	// Finish test event
	if b.SessionStatus == models.SessionCompleted && b.ServerEndDateTime.IsZero() {
		b.ServerEndDateTime = time.Now().UTC()
	}

	if b.BillingEventID == "" {
		b.BillingEventID = bson.NewObjectId()
	}

	_, err = db.C("BillingEvents").Upsert(bson.M{"billingEventUuid": b.BillingEventUUID}, &b)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, b)
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

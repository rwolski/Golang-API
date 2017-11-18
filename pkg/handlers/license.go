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

// RegisterLicenseEndpoints registers endpoints
func RegisterLicenseEndpoints(e *echo.Group) {
	g := e.Group("/license")
	g.POST("", saveLicense)
	g.GET("", getLicense)
	g.DELETE("/:id", deleteLicense)
}

func saveLicense(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	a := models.License{}
	err := e.Bind(&a)
	if err != nil {
		return err
	}

	existing := models.License{}
	err = db.C("Licenses").Find(bson.M{"serialKey": a.SerialKey}).One(&existing)

	if err == nil {
		fmt.Printf("Found license: %+v", existing)

		if a.MachineKey != existing.MachineKey {
			// Found existing key with different machine key
			fmt.Printf("Trying to activate license with wrong PC")
			return e.String(http.StatusConflict, "Serial key already in use.")
		}

		a.LicenseID = existing.LicenseID
		_, err = db.C("Licenses").UpsertId(existing.LicenseID, &a)
		if err != nil {
			return err
		}

		fmt.Printf("Updated license: %+v", a)
	} else {
		if a.LicenseID == "" {
			a.LicenseID = bson.NewObjectId()
		}

		fmt.Printf("New license: %+v", a)

		err = db.C("Licenses").Insert(&a)
		if err != nil {
			return err
		}
	}

	return e.JSON(http.StatusOK, a)
}

func getLicense(e echo.Context) error {
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

	a := models.License{}
	if id.Valid() {
		err = db.C("Licenses").FindId(id).One(&a)
	} else {
		err = db.C("Licenses").Find(bson.M{"licenseUuid": uuid}).One(&a)
	}
	if err != nil {
		return e.NoContent(http.StatusNotFound)
	}
	return e.JSON(http.StatusOK, a)
}

func deleteLicense(e echo.Context) error {
	db := e.Get("database").(*mgo.Database)
	if db == nil {
		return fmt.Errorf("Bad database session")
	}

	id := bson.ObjectIdHex(e.Param("id"))
	if !id.Valid() {
		return fmt.Errorf("Bad parameters")
	}

	err := db.C("Licenses").RemoveId(id)
	if err != nil {
		return err
	}
	return e.NoContent(http.StatusOK)
}

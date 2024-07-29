package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersonInformation(c *gin.Context) {
	personID := c.Param("person_id")

	var info PersonInfo

	query := `select p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code from person p
        join phone ph on  p.id = ph.person_id
        join address_join aj on p.id = aj.person_id
        join address a on aj.address_id = a.id
        where p.id = ?`

	row := DB.QueryRow(query, personID)
	err := row.Scan(&info.Name, &info.PhoneNumber, &info.City, &info.State, &info.Street1, &info.Street2, &info.ZipCode)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "person not found "})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "issue while fetching details"})
		}
		return
	}

	c.JSON(http.StatusOK, info)
}

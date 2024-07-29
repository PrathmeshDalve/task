package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var request CreatePersonRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	tx, err := DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	defer tx.Rollback()

	res, err := tx.Exec("insert into person (name, age) VALUES (?, ?)", request.Name, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	personID, err := res.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	_, err = tx.Exec("insert into phone (person_id, number) VALUES (?, ?)", personID, request.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	res, err = tx.Exec("insert into address (city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)",
		request.City, request.State, request.Street1, request.Street2, request.ZipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	addressID, err := res.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	_, err = tx.Exec("insert into address_join (person_id, address_id) VALUES (?, ?)", personID, addressID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person created"})
}

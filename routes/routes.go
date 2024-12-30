package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Record struct
type Record struct {
	UserInstance UserInstance
	Text         string
	Id           int
	Status       string
	Saved        bool
}

// records collection
var records []Record

// user struct
type UserInstance struct {
	Name   string
	UserId int
}

// user interface
type User interface {
	welcomeUser()
	loadUserRecords(*Record)
	addRecord(string, *Record) Record
	updateRecord(int, string) Record
	deleteRecord(int)
	viewAllRecords(id int) []Record
}

// generate a random id

func randomIdGenerator() int {
	// seed the random number generator
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(100)
}

// add a new record
func (u UserInstance) addRecord(collection *[]Record, record string) Record {

	if record == "" {
		fmt.Println("Entry cannot be empty")
		fmt.Scan(&record)
	}
	defer fmt.Println("Record added successfully")
	fmt.Println("Adding a new record")
	fmt.Println("Enter the record text")
	newRecord := Record{UserInstance: UserInstance{Name: u.Name, UserId: u.UserId}, Text: record, Id: randomIdGenerator(), Status: "incomplete", Saved: false}
	*collection = append(records, newRecord)

	return newRecord
}

// setup router
func SetupRouter() *gin.Engine {
	app := gin.Default()

	// create a new record route
	app.POST("/records", func(c *gin.Context) {
		//bind body to record struct

		var body Record

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		body.Id = randomIdGenerator()

		// create a new record
		c.JSON(http.StatusOK, gin.H{
			"data": body,
		})
	})

	// view all records

	app.GET("/records", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"data": records,
		})

	})

	// update a record

	app.PUT("/records/:id", func(c *gin.Context) {
		// get the record id
		id := c.Param("id")
		newRecord := c.Param("record")

		var updatedRecord Record

		found := false

		// get the record
		for _, record := range records {
			if int, err := strconv.Atoi(id); err == nil {
				if record.Id == int {
					record.Text = newRecord
					record.Saved = false
					updatedRecord = record
					found = true
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}

		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		}

		c.JSON(http.StatusOK, gin.H{
			"data": updatedRecord,
		})
	})

	// get a record by id

	app.GET("/records/:id", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"data": c.Param("id"),
		})
	})

	// delete a record

	app.DELETE("/records/:id", func(c *gin.Context) {
		int, err := strconv.Atoi(c.Param("id"))
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id",
			})
		}

		for i, record := range records {
			if record.Id == int {
				records = append(records[:i], records[i+1:]...)
			}
		}

		c.JSON(200, gin.H{
			"data": c.Param("id"),
		})
	})

	return app
}

func RunApp() {
	fmt.Println("app running")
}

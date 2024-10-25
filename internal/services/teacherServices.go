package services

import (
	"calificationApi/internal/models"
	"calificationApi/internal/server/customErrors"
	"calificationApi/internal/utilities"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"sync"
)

// addTeacher handles the addition of a new teacher to the database.
// @Summary Add a new teacher
// @Description Inserts a new teacher record to the database
// @Accept json
// @Produce json
// @Param teacher body models.Teacher true "New Teacher"
// @Success 200 {object} models.Teacher
// @Failure 500 {object} map[string]string
// @Router /teacher [post]
// @Tags teacher
func addTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	err := utilities.ReadJson(w, r, &teacher)
	if err != nil {
		httpInternalError(w, err.Error())
		utilities.Log.Println(err)
		return
	}
	_, err = dbContext.Teachers.InsertOne(context.TODO(), teacher)
	if err != nil {
		httpInternalError(w, err.Error())
	}
}

// updateTeacher updates an existing teacher's information in the database.
// @Summary Update an existing teacher
// @Description Updates the information of an existing teacher in the database
// @Accept json
// @Produce json
// @Param teacher body models.Teacher true "Updated Teacher"
// @Success 200 {object} models.Teacher
// @Failure 500 {object} map[string]string
// @Router /teacher [put]
// @Tags teacher
func updateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	err := utilities.ReadJson(w, r, &teacher)
	if err != nil {
		httpInternalError(w, err.Error())
		utilities.Log.Println(err)
		return
	}
	filter := bson.M{"carnet": teacher.Carnet}
	update := bson.M{"$set": teacher}
	_, err = dbContext.Teachers.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		httpInternalError(w, err.Error())
		return
	}
}

// deleteTeacher removes a teacher from the database based on the provided "Carnet" parameter.
// @Summary Delete a teacher
// @Description Deletes an existing teacher record from the database using the "Carnet" query parameter.
// @Param Carnet query string true "Teacher Carnet"
// @Success 204 "No Content"
// @Failure 500 {object} map[string]string
// @Router /teacher [delete]
// @Tags teacher
func deleteTeacher(w http.ResponseWriter, r *http.Request) {
	carnet := r.URL.Query().Get("Carnet")
	filter := bson.M{"carnet": carnet}

	_, err := dbContext.Teachers.DeleteOne(context.TODO(), filter)
	if err != nil {
		httpInternalError(w, err.Error())
		return
	}

}

// getTeacher retrieves the details of a teacher based on the provided "Carnet" parameter.
// @Summary Get a teacher's details
// @Description Fetches the information of a teacher from the database using the "Carnet" query parameter.
// @Param Carnet query string true "Teacher Carnet"
// @Success 200 {object} models.Teacher
// @Failure 404 {object} map[string]string
// @Router /teacher [get]
// @Tags teacher
func getTeacher(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	ch := make(chan bool)
	carnet := r.URL.Query().Get("Carnet")
	go anyTeacher(carnet, &wg, ch)
	if <-ch {
		httpNotFoundError(w, customErrors.NewNotFoundMongoError("carnet").Error())
		return
	}
	wg.Wait()
	var teacher models.Teacher
	filter := bson.M{"carnet": carnet}
	err := dbContext.Teachers.FindOne(context.TODO(), filter).Decode(&teacher)
	if err != nil {
		httpNotFoundError(w, customErrors.NewNotFoundMongoError("carnet").Error())
		return
	}
	utilities.WriteJson(w, http.StatusOK, teacher)

}

// @Summary Retrieve all teachers
// @Description Fetch all teacher records from the database and return them as a JSON payload
// @Success 200 {array} models.Teacher "List of teachers"
// @Failure 500 {string} string "Internal server error"
// @Router /teachers [get]
// @Tags teachers
func getTeachers(w http.ResponseWriter) {
	var teachers []models.Teacher
	find, findErr := dbContext.Teachers.Find(context.TODO(), bson.M{})
	if findErr != nil {
		utilities.Log.Println(findErr)
		return
	}
	decodeErr := find.All(context.TODO(), &teachers)
	if decodeErr != nil {
		utilities.Log.Println(decodeErr)
		return
	}
	utilities.WriteJson(w, http.StatusOK, teachers)
}
func anyTeacher(carnet string, wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	defer close(ch)
	filter := bson.D{{"carnet", carnet}}
	err := dbContext.Teachers.FindOne(context.TODO(), filter).Decode(&models.Teacher{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		ch <- false
	} else if err != nil {
		utilities.Log.Println(err)
		ch <- false
	} else {
		ch <- true
	}
}

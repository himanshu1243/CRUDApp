package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	constant "main/Constants"
	models "main/Models"
	sqlconnect "main/SqlConnection"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var db *sql.DB

// CreateStudent godoc
// @Summary Create a new Student
// @Description Create a new Student and stores it in the database
// @Tags Studetn
// @Accept  json
// @Produce  json
// @Param Studentinfo body StudentInfo true "Create Student"
// @Success 200 {object} Studentinfo
// @Router /students [post]
func AddStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constant.HeaderType, constant.JsonString)
	db = sqlconnect.GetMySQLDB()
	defer db.Close()
	student := models.StudentInfo{}
	err := json.NewDecoder(r.Body).Decode(&student) //error
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	ID, err := strconv.Atoi(student.ID)
	if err != nil {
		log.Fatal("conversion not possible")
		return
	}
	_, err = db.Exec("insert into student(Id,Name,Course) values(?,?,?)", ID, student.Name, student.Course)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	json.NewEncoder(w).Encode(student)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constant.HeaderType, constant.JsonString)
	db = sqlconnect.GetMySQLDB()
	defer db.Close()
	params := mux.Vars(r)
	ID, err := strconv.Atoi(params["ID"])
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	_, err = db.Exec("delete from student where ID=?", ID)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

// getStudents godoc
// @Summary Get details of all students present in the database
// @Description Get details of all Students
// @Tags students
// @Accept  json
// @Produce  json
// @Success 200 {array} Students
// @Router /students [get]
func ReadStudents(w http.ResponseWriter, r *http.Request) {
	db = sqlconnect.GetMySQLDB()
	defer db.Close()
	ss := []models.StudentInfo{}
	student := models.StudentInfo{}
	rows, err := db.Query("select*from student")
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for rows.Next() {
		err := rows.Scan(&student.ID, &student.Name, &student.Course)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		ss = append(ss, student)
	}
	err = json.NewEncoder(w).Encode(ss)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Print("After error")
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constant.HeaderType, constant.JsonString)
	db = sqlconnect.GetMySQLDB()
	defer db.Close()
	student := models.StudentInfo{}
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	params := mux.Vars(r)
	fmt.Println(params["ID"])
	ID, err := strconv.Atoi(params["ID"])
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	_, err = db.Exec("update student set Name=?, Course=? where ID=?", student.Name, student.Course, ID)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	json.NewEncoder(w).Encode(student)
}

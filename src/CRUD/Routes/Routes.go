package route

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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
	w.Header().Set("Content-Type", "application/json")
	db = sqlconnect.GetMySQLDB()
	defer db.Close()
	s := models.StudentInfo{}
	err := json.NewDecoder(r.Body).Decode(&s) //error
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	ID, err := strconv.Atoi(s.ID)
	if err != nil {
		log.Fatal("conversion not possible")
		return
	}
	_, errin := db.Exec("insert into student(Id,Name,Course) values(?,?,?)", ID, s.Name, s.Course)
	if errin != nil {
		log.Fatal(err.Error())
		return
	}
	json.NewEncoder(w).Encode(s)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db = sqlconnect.GetMySQLDB()
	defer db.Close()
	params := mux.Vars(r)
	ID, err := strconv.Atoi(params["ID"])
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	_, errin := db.Exec("delete from student where ID=?", ID)
	if errin != nil {
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
func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db = sqlconnect.GetMySQLDB()
	defer db.Close()
	ss := []models.StudentInfo{}
	s := models.StudentInfo{}
	rows, err := db.Query("select*from student")
	if err != nil {
		log.Fatal(err.Error()) // log instead of fmt
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		for rows.Next() {
			err := rows.Scan(&s.ID, &s.Name, &s.Course)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			ss = append(ss, s)
		}
		err := json.NewEncoder(w).Encode(ss)
		if err != nil {
			log.Fatal(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Print("After error") // error handling
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db = sqlconnect.GetMySQLDB()
	defer db.Close()
	s := models.StudentInfo{}
	err := json.NewDecoder(r.Body).Decode(&s)
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

	_, errin := db.Exec("update student set Name=?, Course=? where ID=?", s.Name, s.Course, ID)
	if errin != nil {
		json.NewEncoder(w).Encode("ssssss")
		log.Fatal(err.Error())
		return
	}

	json.NewEncoder(w).Encode(s)
}

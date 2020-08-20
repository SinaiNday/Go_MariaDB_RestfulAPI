package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Restful API using MariaDB!")
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var Newstudent Student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "wrong data")
	}
	json.Unmarshal(reqBody, &Newstudent)
	insForm, err := db.Prepare("INSERT INTO students(id,firstname,lastname,age) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insForm.Exec(Newstudent.ID, Newstudent.Firstname, Newstudent.Lastname, Newstudent.Age)
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(Newstudent, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students := Student{}
	res := []Student{}
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM students") //  ORDER BY id DESC by the latest
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, age int
		var firstname, lastname string
		err = selDB.Scan(&id, &firstname, &lastname, &age)
		if err != nil {
			panic(err.Error())
		}
		students.ID = id
		students.Firstname = firstname
		students.Lastname = lastname
		students.Age = age

		res = append(res, students)
	}

	Conv, _ := json.MarshalIndent(res, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

	defer db.Close()

}
func GetOneStudent(w http.ResponseWriter, r *http.Request) {
	StudentID := mux.Vars(r)["id"]
	students := Student{}
	//res :=[]Student{}
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM students WHERE id=?", StudentID)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, age int
		var firstname, lastname string
		err = selDB.Scan(&id, &firstname, &lastname, &age)
		if err != nil {
			panic(err.Error())
		}
		students.ID = id
		students.Firstname = firstname
		students.Lastname = lastname
		students.Age = age

		// res = append(res, students)
	}

	Conv, _ := json.MarshalIndent(students, "", " ")

	fmt.Fprintf(w, "%s", string(Conv))

	defer db.Close()

}

func CountAllStudents(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	var count string

	err := db.QueryRow("SELECT COUNT(*) FROM students").Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "%s ", count)

}
func DeleteOneStudent(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	StudentID := mux.Vars(r)["id"]
	delForm, err := db.Prepare("DELETE FROM students WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(StudentID)
	fmt.Fprintf(w, "deleted successfully the student num %s ", StudentID)
	defer db.Close()

}

func DeleteAllStudents(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	delForm, err := db.Prepare("TRUNCATE students")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec()
	fmt.Fprintf(w, "deleted all successfully")
	defer db.Close()

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	StudentID := mux.Vars(r)["id"]
	var UpdateStudent Student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data properly")
	}
	json.Unmarshal(reqBody, &UpdateStudent)

	update, err := db.Prepare("UPDATE students SET firstname=?, lastname=?, age=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(UpdateStudent.Firstname, UpdateStudent.Lastname, UpdateStudent.Age, StudentID)
	fmt.Fprintf(w, "updated successfully")
	defer db.Close()

}

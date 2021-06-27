package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/julienschmidt/httprouter"
	"golang/models"
	"net/http"
)



// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.ToDo)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetDB() *pg.DB {
	opt, err := pg.ParseURL("postgres://b2b_portal:b2b_portal@localhost:5432/golang?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return pg.Connect(opt)
}

func main() {
	fmt.Println("Hello world")
	opt, err := pg.ParseURL("postgres://b2b_portal:b2b_portal@localhost:5432/golang?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	ctx := context.Background()
	if err = db.Ping(ctx); err != nil {
		panic(err)
	}

	var version string
	_, err = db.QueryOneContext(ctx, pg.Scan(&version), "SELECT version()")
	if err != nil {
		panic(err)
	}
	fmt.Println(version)

	//err = createSchema(db)
	//if err != nil {
	//	return
	//}

	//todo1 := ToDo{
	//	Id:         1,
	//	Title:      "sleep",
	//	IsComplete: false,
	//	UpdateAt:   time.Now(),
	//	CreatedAt:  time.Now(),
	//}

	//todo2 := models.ToDo{
	//	Id:         2,
	//	Title:      "eat",
	//	IsComplete: true,
	//	UpdateAt:   time.Now(),
	//	CreatedAt:  time.Now(),
	//}

	//_, err = db.Model(&todo2).Insert()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//todo := new(models.ToDo)
	//err = db.Model(todo).Where("id = ?", 1).Select()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//todo.IsComplete = true
	//_, err = db.Model(todo).Where("id = ?", 1).Update()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//todoList := new([]models.ToDo)
	//err = db.Model(todoList).Where("id = ?", 2).Select()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("%+v\n", todoList)

	router := httprouter.New()

	router.GET("/", hello)
	router.GET("/todo/all", ToDoGetAll)
	router.GET("/todo/get/:id", ToDoGetById)

	err = http.ListenAndServe(":8090", router)
	if err != nil {
		return
	}
}


func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, _ = w.Write([]byte("Hello"))
}

func ToDoGetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	todoList := new([]models.ToDo)
	err := GetDB().Model(todoList).Select()
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonByte, err := json.Marshal(todoList)
	if err != nil {
		return
	}

	_, _ = w.Write(jsonByte)
}


func ToDoGetById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	todoList := new([]models.ToDo)
	err := GetDB().Model(todoList).Where("id = ?", ps.ByName("id")).Select()
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonByte, err := json.Marshal(todoList)
	if err != nil {
		return
	}
	_, _ = w.Write(jsonByte)
}
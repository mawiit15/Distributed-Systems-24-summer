package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

//todo struct to hold tasks
type Todo struct {
    Task string `json:"task,omitempty" bson:"task,omitempty"`
}

var client *mongo.Client

//-----GET method for all todos-----//
func get_todo_all(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
	//'slice' -> basically an array
    var todos []Todo
	//get table
    collection := client.Database("todo_db").Collection("todos")
	//context -> will cancel db query after 10s if no response
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//grab all entries, this returns cursor with all todos
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
		//return err500 if query fails
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "error1": "` + err.Error() + `" }`))
        return
    }
	//keyword defer -> called when function returns, even if errors occur
    defer cursor.Close(ctx)
	//iterate through cursor, append entries to todo-slice
    for cursor.Next(ctx) {
        var todo Todo
        cursor.Decode(&todo)
        todos = append(todos, todo)
    }
    if err := cursor.Err(); err != nil {
		//return err500 if fetching fails
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "error2": "` + err.Error() + `" }`))
        return
    }
	//response 200 OK
	response.WriteHeader(http.StatusOK)
	//return list of todos in json format
    json.NewEncoder(response).Encode(todos)
}

//-----GET method for single task, by title-----//
func get_todo_single(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
	//get url parameters (url /x/y/z -> [x,y,z])
    params := mux.Vars(request)
	//grab the task from /todos/{task}
    task := params["task"]
    var todo Todo
	//get table
    collection := client.Database("todo_db").Collection("todos")
	//context -> will cancel db query after 10s if no response
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//grab the task specified by the url parameter (bson.M(...) is the filter)
	//task is saved into todo (.Decode(&todo))
	//err stays nil if task is found, FindOne will return error if no matching task is found
    err := collection.FindOne(ctx, bson.M{"task": task}).Decode(&todo)
    if err != nil {
		//return err500 if query fails
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "error3": "` + err.Error() + `" }`))
        return
    }
	//response 200 OK
	response.WriteHeader(http.StatusOK)
	//return single todo in json format
    json.NewEncoder(response).Encode(todo)
}

//-----DELETE method for single task, by title-----//
func delete_todo_single(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
	//get url parameters
    params := mux.Vars(request)
	//grab task from parameter list
    task := params["task"]
	//get table
    collection := client.Database("todo_db").Collection("todos")
	//context -> will cancel query after 10s if no response
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//delete task specified in url
	//err stays nil unless DeleteOne cant find specified task
    _, err := collection.DeleteOne(ctx, bson.M{"task": task})
    if err != nil {
		//return err500 if query fails
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "error4": "` + err.Error() + `" }`))
        return
    }
	//response 200 OK
	response.WriteHeader(http.StatusOK)
}

func main() {
    //-----mongoDB connection-----//

	//get dtails from environment variable
    mongoURI := os.Getenv("MONGO_URI")
	//context -> cancel after 10s if no connection established
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//cancel will stop db connection. call this upon returning to ensure proper cleanup on exit.
    defer cancel()
	//grab credentials
    client_opts := options.Client().ApplyURI(mongoURI)
	//connect to db
    client, _ = mongo.Connect(ctx, client_opts)

	//-----API mapping-----//

    //create router instance
    router := mux.NewRouter()
	//register endpoints
    router.HandleFunc("/todos/", get_todo_all).Methods("GET")
    router.HandleFunc("/todo/{task}", get_todo_single).Methods("GET")
    router.HandleFunc("/todo/{task}", delete_todo_single).Methods("DELETE")
	//default handler for incoming requests
    http.Handle("/", router)
    
    //start server on port 8080, use default handler
	//log.Fatal logs errors and exits app if error occurs
    log.Fatal(http.ListenAndServe(":8080", nil))
}
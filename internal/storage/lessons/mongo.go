package lessons

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/gorilla/mux"
// 	"github.com/imirjar/mongo-golang/internal/models"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

type Mongo struct {
	dbConn string
	client *mongo.Client
}

func NewDB(conn string) *Mongo {
	mongo := &Mongo{
		dbConn: conn,
	}
	log.Print(conn)

	if mongo.ping() {
		log.Print("mongo ping ok")
	}

	if err := mongo.connect(conn); err != nil {
		panic(err)
	}
	defer mongo.disconnect()
	log.Print("mongo conn ok")

	return mongo
}

func (m *Mongo) ping() bool {
	return true
}

func (m *Mongo) connect(conn string) error {
	client, err := mongo.Connect(options.Client().ApplyURI(conn))
	// client, err := mongo.Connect(conn)
	if err != nil {
		panic(err)
	}
	m.client = client

	return nil
}

func (m *Mongo) disconnect() error {
	return m.client.Disconnect(context.Background())
}

// // connect to MongoDB and get data by params
// func getData(db string, table string, obj Modeller, bsonV primitive.M) Modeller {

// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Printf("Error while parsing .env file: %v\n", err)
// 	}

// 	client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer mongo.Close(client, ctx, cancel)

// 	cursor, err := mongo.Query(client, ctx, db, table, bsonV, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := cursor.All(ctx, &obj); err != nil {
// 		fmt.Println(err)
// 	}

// 	return obj
// }

// func putData(db string, table string, obj Modeller, id primitive.ObjectID) Modeller {
// 	//подгружаем файл env
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Printf("Error while parsing .env file: %v\n", err)
// 	}
// 	//подключаемся к Mongodb по переменной подключения из env файла
// 	client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer mongo.Close(client, ctx, cancel)

// 	coll := client.Database(db).Collection(table)
// 	update := bson.D{{"$set", obj}}
// 	result, err := coll.UpdateOne(context.TODO(), bson.D{{"_id", id}}, update)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return result
// }

// func UploadFile(w http.ResponseWriter, r *http.Request) {

// 	// Parse our multipart form, 10 << 20 specifies a maximum
// 	// upload of 10 MB files.
// 	r.ParseMultipartForm(10 << 20)

// 	collection := r.FormValue("collection")
// 	// fmt.Println(collection)
// 	documentId := r.FormValue("documentId")
// 	// fmt.Println(documentId)
// 	uploadedFile, handler, err := r.FormFile("file")
// 	// fmt.Println(handler)
// 	if err != nil {
// 		fmt.Println("Error Retrieving the File")
// 		fmt.Println(err)
// 		return
// 	}

// 	defer uploadedFile.Close()

// 	tempFile, err := ioutil.TempFile("storage/files", "*"+handler.Filename)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer tempFile.Close()
// 	fileBytes, err := ioutil.ReadAll(uploadedFile)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	tempFile.Write(fileBytes)

// 	err = godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Printf("Error while parsing .env file: %v\n", err)
// 	}
// 	client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer mongo.Close(client, ctx, cancel)

// 	file := models.File{
// 		Id:   primitive.NewObjectID(),
// 		Name: handler.Filename,
// 		Link: tempFile.Name(),
// 	}

// 	coll := client.Database("sspkSite").Collection(collection)
// 	id, _ := primitive.ObjectIDFromHex(documentId)
// 	filter := bson.D{{"_id", id}}
// 	update := bson.D{{"$push", bson.D{{"documents", file}}}}
// 	result, err := coll.UpdateOne(context.TODO(), filter, update)
// 	// fmt.Println(result)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
// 	// fmt.Printf("File Size: %+v\n", handler.Size)
// 	// fmt.Printf("MIME Header: %+v\n", handler.Header)

// 	json.NewEncoder(w).Encode(result)

// 	// fmt.Fprintf(w, "Successfully Uploaded File\n")
// }

// func DeleteFile(w http.ResponseWriter, r *http.Request) {

// 	var file models.File
// 	err := json.NewDecoder(r.Body).Decode(&file)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	vars := mux.Vars(r)
// 	collection := vars["collection"]
// 	elementId, err := primitive.ObjectIDFromHex(vars["elementId"])
// 	if err != nil {
// 		fmt.Printf("Can't make primitive %v\n", err)
// 	}

// 	client, ctx, cancel, err := Connect(os.Getenv("MONGODB_URL"))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer mongo.Close(client, ctx, cancel)

// 	coll := client.Database("sspkSite").Collection(collection)

// 	filter := bson.M{
// 		"_id": elementId,
// 		"documents": bson.M{
// 			"$elemMatch": bson.M{
// 				"_id": file.Id,
// 			},
// 		},
// 	}

// 	update := bson.M{
// 		"$pull": bson.M{
// 			"documents": bson.M{
// 				"_id": file.Id,
// 			},
// 		},
// 	}

// 	result, err := coll.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = os.Remove(file.Link)
// 	if err != nil {
// 		fmt.Println("Ну удалось удалить файл", err)
// 	}

// 	json.NewEncoder(w).Encode(result)

// }

// // This is a user defined method to close resources.
// // This method closes mongoDB connection and cancel context.
// func Close(client *mongo.Client, ctx context.Context,
// 	cancel context.CancelFunc) {

// 	// CancelFunc to cancel to context
// 	defer cancel()

// 	// client provides a method to close
// 	// a mongoDB connection.
// 	defer func() {

// 		// client.Disconnect method also has deadline.
// 		// returns error if any,
// 		if err := client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()
// }

// // This is a user defined method that returns mongo.Client,
// // context.Context, context.CancelFunc and error.
// // mongo.Client will be used for further database operation.
// // context.Context will be used set deadlines for process.
// // context.CancelFunc will be used to cancel context and
// // resource associated with it.

// func Connect(uri string) (*mongo.Client, context.Context,
// 	context.CancelFunc, error) {

// 	// ctx will be used to set deadline for process, here
// 	// deadline will of 30 seconds.
// 	ctx, cancel := context.WithTimeout(context.Background(),
// 		30*time.Second)

// 	// mongo.Connect return mongo.Client method
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	return client, ctx, cancel, err
// }

// // This is a user defined method that accepts
// // mongo.Client and context.Context
// // This method used to ping the mongoDB, return error if any.
// func Ping(client *mongo.Client, ctx context.Context) error {

// 	// mongo.Client has Ping to ping mongoDB, deadline of
// 	// the Ping method will be determined by cxt
// 	// Ping method return error if any occurred, then
// 	// the error can be handled.
// 	if err := client.Ping(ctx, readpref.Primary()); err != nil {
// 		return err
// 	}
// 	fmt.Println("connected successfully")
// 	return nil
// }

// // insertOne is a user defined method, used to insert
// // documents into collection returns result of InsertOne
// // and error if any.
// func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

// 	// select database and collection ith Client.Database method
// 	// and Database.Collection method
// 	collection := client.Database(dataBase).Collection(col)

// 	// InsertOne accept two argument of type Context
// 	// and of empty interface
// 	result, err := collection.InsertOne(ctx, doc)
// 	return result, err
// }

// // insertMany is a user defined method, used to insert
// // documents into collection returns result of
// // InsertMany and error if any.
// func insertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {

// 	// select database and collection ith Client.Database
// 	// method and Database.Collection method
// 	collection := client.Database(dataBase).Collection(col)

// 	// InsertMany accept two argument of type Context
// 	// and of empty interface
// 	result, err := collection.InsertMany(ctx, docs)
// 	return result, err
// }

// func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

// 	// select database and collection.
// 	collection := client.Database(dataBase).Collection(col)

// 	// collection has an method Find,
// 	// that returns a mongo.cursor
// 	// based on query and field.
// 	result, err = collection.Find(ctx, query,
// 		options.Find().SetProjection(field))
// 	return
// }

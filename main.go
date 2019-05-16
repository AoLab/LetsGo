package testing101
import (
	"context"
	"time"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type test struct {
 	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
func JsonHandler(c echo.Context) (err error) {
  t :=test{}
  defer c.Request().Body.Close()
  b,err := ioutil.ReadAll(c.Request().Body)
  if err!=nil {
	  log.Printf("failed to rendering request body:%s",err)
	  return c.String(http.StatusInternalServerError,"")
  }
  err=json.Unmarshal(b,&t)
  if err!=nil {
          log.Printf("failed to Unmarshaling:%s",err)
          return c.String(http.StatusInternalServerError,"")
  }
  client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
  if err!= nil{
	  log.Fatal(err)
  }
  ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err!= nil{
          log.Fatal(err)
  }
  collection := client.Database("testing").Collection("geo")
  res, err := collection.InsertOne(ctx,t)
  if err!= nil{
          log.Fatal(err)
  }
  id := res.InsertedID
  defer cancel()
  return c.String(http.StatusOK, fmt.Sprintf("%s\n%s\n%s\n",id, string(t.Lat),string(t.Lng)))
}
func Pinging(c echo.Context) error {
                return c.String(http.StatusOK, "man amdam")
        }
func NewServer() {
	e := echo.New()
	e.GET("/ping",Pinging)
	e.POST("/posting",JsonHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

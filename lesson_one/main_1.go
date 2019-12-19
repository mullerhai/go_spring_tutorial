package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-spring/go-spring-boot-starter/starter-gin"
	_ "github.com/go-spring/go-spring-boot-starter/starter-web"
	SpringWeb "github.com/go-spring/go-spring-web/spring-web"
	SpringBoot "github.com/go-spring/go-spring/spring-boot"
	"net/http"
	"os"
)

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(client)
	SpringBoot.RegisterBean(new(Controller))
}

type Controller struct{}

func (c *Controller) InitWebBean(wc SpringWeb.WebContainer) {

	wc.SetPort(8894)
	//wc.GET("/", c.Home)
	wc.GET("/",c.Detail)

}

func (c *Controller) Home(ctx SpringWeb.WebContext) {
	ctx.String(http.StatusOK, "OK! daiji liudong ")
}

func (c *Controller) Detail(ctx SpringWeb.WebContext){
	fmt.Println("come ing")
	//ctx.HTML(http.StatusOK,"<h1> puck hockey </h1> <h5>book </h5>")
	jk :="{'name':45,}"
	kks ,_:=json.Marshal(jk)
	fmt.Println(kks)
	//ctx.JSON(http.StatusOK,kks)
	path :="templates/html/home.html"  //"./template/html/home.html"
	image_data, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer image_data.Close()
	//ctx.File(path)
	ctx.Stream(http.StatusOK,"html",image_data)

	name :=ctx.QueryParam("name")
	fmt.Println(name)
	ctx.ResponseWriter()
}
func main() {
	SpringBoot.RunApplication("config/")

}




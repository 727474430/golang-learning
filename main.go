package main

import (
	"fifa_server/rest"
	_ "fifa_server/rest"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
)

type Config struct {
	Hostname string `ini:"db_hostname"`
	Port     string `ini:"db_port"`
	Username string `ini:"db_username"`
	Password string `ini:"db_password"`
}

func main() {
	r := gin.Default()
	rest.ApiServerInit(r)

	/*res, err := http.Get("https://www.fifa.com/worldcup/matches/")
	if err != nil {
		log.Fatal(err)
	}
	defer  res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	content := document.Find(".fi-matchlist .fi-mu-list .fi-mu__link div")
	fmt.Print(content)*/

	/*config, err := loadConfigConf("E:\\wl-projects\\go-projects\\fifa_server\\configs\\db.conf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(config)*/
}

func loadConfigConf(path string) (Config, error) {
	var config Config
	conf, err := ini.Load(path)
	if err != nil {
		log.Println("load config error")
		return config, err
	}

	conf.BlockMode = false
	err = conf.MapTo(&config)
	if err != nil {
		log.Panicln("mapto config error")
		return config, err
	}

	return config, nil
}

package main

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/cloudfoundry-community/gautocloud/connectors/generic"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"
)

func init() {
	gautocloud.RegisterConnector(generic.NewConfigGenericConnector(CloudConfig{}))
}

func main() {
	var config CloudConfig
	err := gautocloud.Inject(&config)
	exitOnError(err)

	var dbConfig dbtype.MysqlDatabase
	err = gautocloud.Inject(&dbConfig)
	exitOnError(err)

	port := 8080
	if os.Getenv("PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	adminUiConf := AdminUiConfig{
		CloudConfig: config,
		Port:        port,
		BindAddress: "0.0.0.0",
		DbURI:       fmt.Sprintf("mysql2://%s:%s@%s:%d/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database),
	}
	b, err := yaml.Marshal(adminUiConf)
	exitOnError(err)
	exitOnError(ioutil.WriteFile("config.yml", b, 0666))

	wd, err := os.Getwd()
	exitOnError(err)
	args := []string{path.Join(wd, "bin", "admin"), "-c", path.Join(wd, "config.yml")}
	var adminUiExec *exec.Cmd
	if len(os.Args) > 1 {
		argsFromUser := os.Args[1:]
		args = append(args, argsFromUser...)
	}
	adminUiExec = exec.Command("ruby", args...)
	adminUiExec.Stdout = os.Stdout
	adminUiExec.Stderr = os.Stderr
	go func() {
		for {
			time.Sleep(5 * time.Second)
			sinkLogsFromFileToStdout(config.LogFile)
		}

	}()
	exitOnError(adminUiExec.Run())
}

func sinkLogsFromFileToStdout(logfile string) {
	if _, err := os.Stat(logfile); os.IsNotExist(err) {
		return
	}
	b, err := ioutil.ReadFile(logfile)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Print(string(b))
	err = ioutil.WriteFile(logfile, []byte{}, 0666)
	if err != nil {
		log.Println(err.Error())
	}
}
func exitOnError(err error) {
	if err == nil {
		return
	}
	log.Fatal(err.Error())
}

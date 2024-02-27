package models

import (
	"ButterHost69/PKr-client/encrypt"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Connections struct {
	ConnectionSlug 	string 	`json:"connection_slug"`
	Password       	string 	`json:"password"`
	CurrentIP		string	`json:"current_ip"`

	SendFiles		
}

type Files struct 

type UsersConfig struct {
	User           string        `json:"user"`
	AllConnections []Connections `json:"all_connections"`
}

const (
	ROOT_DIR     = "tmp"
	MY_KEYS_PATH = "tmp/mykeys"
	CONFIG_FILE  = "tmp/userConfig.json"
)

var (
	MY_USERNAME string
)

func CreateUserIfNotExists() {
	if _, err := os.Stat(ROOT_DIR + "/userConfig.json"); os.IsNotExist(err) {
		fmt.Println("!! 'tmp' No such DIR exists ")

		var username string
		fmt.Println(" [*] Register [*]")
		fmt.Print(" > Username: ")
		fmt.Scanln(&username)
		MY_USERNAME = username

		usconf := UsersConfig{
			User: username,
		}

		jsonbytes, err := json.Marshal(usconf)
		if err != nil {
			fmt.Println("~ Unable to Parse Username to Json")
		}

		if err = os.Mkdir(ROOT_DIR, 0777); err != nil {
			fmt.Println("~ Folder tmp exists")
		}
		err = os.WriteFile(ROOT_DIR+"/userConfig.json", jsonbytes, 0777)
		if err != nil {
			log.Fatal(err.Error())
		}

		if err = os.Mkdir(MY_KEYS_PATH, 0777); err != nil {
			fmt.Println("~ Folder tmp exists")
		}

		private_key, public_key := encrypt.GenerateRSAKeys()
		if private_key == nil && public_key == nil {
			panic("Could Not Generate Keys")
		}

		if err = encrypt.StorePrivateKeyInFile(MY_KEYS_PATH+"/privatekey.pem", private_key); err != nil {
			panic(err.Error())
		}

		if err = encrypt.StorePublicKeyInFile(MY_KEYS_PATH+"/publickey.pem", public_key); err != nil {
			panic(err.Error())
		}

		fmt.Printf(" ~ Created User : %s\n", username)
	}
}

func AddConnection(connection_slug string, password string) {

}

func AddConnectionInUserConfig(connection_slug string, password string, connectionIP string) error {
	file, err := os.Open(CONFIG_FILE)
	if err != nil {
		fmt.Println("error in opening config file.... pls check if tmp/userConfig.json available ")
		return err
	}
	defer file.Close()

	var userConfig UsersConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&userConfig)
	if err != nil {
		fmt.Println("error in decoding json data")
		return err
	}

	connection := Connections{
		ConnectionSlug: connection_slug,
		Password:       password,
		CurrentIP: connectionIP,
	}

	userConfig.AllConnections = append(userConfig.AllConnections, connection)
	newUserConfig := UsersConfig{
		User:           userConfig.User,
		AllConnections: userConfig.AllConnections,
	}	
	jsonData, err := json.MarshalIndent(newUserConfig, "", "	")
	if err != nil {
		fmt.Println("error occured in Marshalling the data to JSON")
		fmt.Println(err)
		return err
	}

	// fmt.Println(string(jsonData))
	err = os.WriteFile(CONFIG_FILE, jsonData, 0777)
	if err != nil {
		fmt.Println("error occured in storing data in userconfig file")
		fmt.Println(err)
		return err
	}

	return nil
}

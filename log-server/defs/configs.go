package defs
import (
	// "bufio"
	// "io"
	"io/ioutil"
	// "os"
	"fmt"
	"encoding/json"
	"log"
	"reflect"
)

type DbConfig struct {
	Nickname string		`json:"nickname"`
	Hostname string		`json:"hostname"`
	Port int			`json:"port"`
    User string			`json:"user"`
    Password string		`json:"password"`
    ConnPoolSize int	`json:"connPoolSize"`
}

func UnmarshalDbConfig(bytes []byte) map[string]DbConfig {
	var obj map[string]DbConfig
	if err := json.Unmarshal(bytes, &obj); err != nil {
		log.Fatal(err)
	}
	keys := reflect.ValueOf(obj).MapKeys()
	for _, k := range keys {
		key := k.String()
		var conf DbConfig = obj[key]
		conf.Nickname = key
		obj[key] = conf
	}
	return obj
}

var (
	MongoDbConfigs map[string]DbConfig
)

func LoadConfigs() {
	LoadEnvs();

	fmt.Println(CONFIG_MONGO)

	blob, err := ioutil.ReadFile(CONFIG_MONGO)
	if err != nil {
		panic(err)
		log.Printf(`Failed to read %s\n`, CONFIG_MONGO)
	}
	// fmt.Print(string(blob))
	fmt.Println("Mongo DB Configs")
	MongoDbConfigs = UnmarshalDbConfig([]byte(blob))
	for k := range MongoDbConfigs {
		fmt.Println(MongoDbConfigs[k])
	}


	// counts := make(map[Size]int)
	// for _, size := range inventory {
	// 	counts[size] += 1
	// }

	// fmt.Printf("Inventory Counts:\n* Small:        %d\n* Large:        %d\n* Unrecognized: %d\n",
	// 	counts[Small], counts[Large], counts[Unrecognized])

}
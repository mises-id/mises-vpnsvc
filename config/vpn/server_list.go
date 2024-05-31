package vpn

import (
	"encoding/json"
	"fmt"
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"os"
	"sync"
)

var (
	serverListInit sync.Once
	ServerList []*pb.GetServerListItem
	ServerAddressList map[string]struct{}
)

func InitConfig() {
	serverListInit.Do(func() {
		fmt.Println("vpnsvc server list initializing...")
		projectRootPath, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		configFile := projectRootPath + "/vpn-server-list.json"
		jsonData, err := os.ReadFile(configFile)
		if err!= nil {
			fmt.Println(err)
			panic(err)
		}
		err = json.Unmarshal(jsonData, &ServerList)
		if err!= nil {
			fmt.Println(err)
			panic(err)
		}
		if len(ServerList) == 0 {
			panic("empty server list")
		}
		ServerAddressList = make(map[string]struct{}, len(ServerList))
		for _, v := range ServerList {
			ServerAddressList[v.Ip] = struct{}{}
		}
		fmt.Println("vpnsvc server list:", ServerList)
	})
}
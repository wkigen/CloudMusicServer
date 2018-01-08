package utils

import (
	"encoding/xml"
    "io/ioutil"
    "os"
    "errors"
    "github.com/golang/glog"
)

type Config struct{
    XMLName         xml.Name `xml:"config"`
    BasePath        string   `xml:"basePath"`
    ServerToken     string   `xml:"server_token"`
    Zookeeper       ZookeeperConfig `xml:"zookeeper"`
    GateServer      GateServerConfig `xml:"gate_server"`
    DataBase        DataBaseConfig `xml:"database"`
    Servers         ServersConfig `xml:"servers"`
}

type ZookeeperConfig struct{
    XMLName         xml.Name `xml:"zookeeper"`
    ZookeeperNode   []ZookeeperNodeConfig `xml:"zookeeper_node"`
}

type ZookeeperNodeConfig struct{
    XMLName     xml.Name `xml:"zookeeper_node"`
    IP          string `xml:"ip"`
}

type GateServerConfig struct{
    XMLName     xml.Name `xml:"gate_server"`
    Ip        string `xml:"ip"`
}

type DataBaseConfig struct{
    DataBase    xml.Name `xml:"database"`
    Type        string `xml:"type"`
    Name        string `xml:"name"`
    Accout      string `xml:"accout"`
    Password    string `xml:"passowrd"`
    Ip          string `xml:"ip"`
}

type ServersConfig struct{
    XMLName     xml.Name `xml:"servers"`
    Server      []ServerConfig `xml:"server"`
}

type ServerConfig struct{
    XMLName     xml.Name `xml:"server"`
    Name        string `xml:"name"`
    Ip        string `xml:"ip"`
}

func (self *Config)GetZookeeperIp() (ip []string){
    zkAddr := []string{}
    for _,zookeeper := range self.Zookeeper.ZookeeperNode {
        zkAddr = append(zkAddr,zookeeper.IP)
    }
    return zkAddr
}

func (self *Config)GetServerIp(name string) (ip string){
	var serverIp string
	for _,s := range self.Servers.Server {
		if( s.Name == name){
			serverIp = s.Ip
		}
	}
	return serverIp
}

func ReadConfig() (conf Config,err  error){

    config := Config{}

    file, err := os.Open("conf/config.xml")    
    if err != nil {
        glog.Fatal("can not read the config ,%s",err)
        return config, errors.New("can not read the config")
    }
    defer file.Close()
    
    data, err := ioutil.ReadAll(file)
    if err != nil {
        glog.Fatal("can not read the config ,%s",err)
        return config, errors.New("can not read the config")
    }

    err = xml.Unmarshal(data, &config)
    if err != nil {
        glog.Fatal("can not read the config ,%s",err)
        return config, errors.New("can not read the config")
    }

    return config,nil
}
package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
)

type ClashXConfig struct {
	Port               int      `yaml:"port"`
	SocksPort          int      `yaml:"socks-port"`
	RedirPort          int      `yaml:"redir-port"`
	AllowLan           bool     `yaml:"allow-lan"`
	Mode               string   `yaml:"mode"`
	LogLevel           string   `yaml:"log-level"`
	ExternalController string   `yaml:"external-controller"`
	Secret             string   `yaml:"secret"`
	Proxy              []Proxy  `yaml:"Proxy"`
	Rule               []string `yaml:"Rule"`
}
type Proxy struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Server  string `yaml:"server"`
	Port    string `yaml:"port"`
	Uuid    string `yaml:"uuid"`
	AlterId string `yaml:"alterId"`
	Cipher  string `yaml:"cipher"`
}
type ProxyGroup struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Proxies []string `yaml:"proxies"`
}

func main() {
	url := ""
	for idx, args := range os.Args {
		//fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
		if idx == 1 {
			url = args
		}
	}
	resp, err := http.Get(url)
	conf := &ClashXConfig{}
	if err != nil {
		// handle error
		fmt.Print(err.Error())
	} else {
		defer resp.Body.Close()

		yaml.NewDecoder(resp.Body).Decode(conf)
		for _, p := range conf.Proxy {
			s := p.Name + " = " + p.Type + ", " + p.Server + ", " + p.Port + ", username=" + p.Uuid + ", tls=false, ws-path=/v2, ws-headers=alterId:2|X-Header-2:value"
			//println(s)
			fmt.Fprintln(os.Stdout, s)
		}
	}
}

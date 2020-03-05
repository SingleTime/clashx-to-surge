package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"strings"
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
	var url string
	var blackKeys string
	var isHelp bool
	flag.StringVar(&url, "url", "", "订阅地址")

	flag.StringVar(&blackKeys, "blackKeys", "", "黑名单关键词")
	flag.BoolVar(&isHelp, "help", false, "帮助文档")
	flag.Parse()
	if isHelp {
		flag.Usage()
		return
	}
	if url == "" {
		flag.Usage()
		return
	}
	//flag.Usage()
	keys := strings.Split(blackKeys, ",")
	resp, err := http.Get(url)
	conf := &ClashXConfig{}
	if err != nil {
		// handle error
		fmt.Print(err.Error())
	} else {
		defer resp.Body.Close()

		err := yaml.NewDecoder(resp.Body).Decode(conf)
		if err != nil {
			// handle error
			fmt.Print(err.Error())
		}
		for _, p := range conf.Proxy {
			var inBlack = false
			for _, k := range keys {
				if len(k) > 0 && strings.Contains(p.Name, k) {
					inBlack = true
				}

			}
			if !inBlack {
				s := p.Name + " = " + p.Type + ", " + p.Server + ", " + p.Port + ", username=" + p.Uuid + ", tls=false, ws-path=/v2, ws-headers=alterId:2|X-Header-2:value"
				//println(s)
				_, err := fmt.Fprintln(os.Stdout, s)
				if err != nil {
					// handle error
					fmt.Print(err.Error())
				}
			}
		}
	}
}

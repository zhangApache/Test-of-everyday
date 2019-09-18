package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 镜像结构
type Image struct {
	Created  uint64
	Id string
	ParentId string
	RepoTags []string
	Size uint64
	VirtualSize uint64
}

// 容器结构
type Container struct {
	Id string `json:"Id"`
	Names []string `json:"Names"`
	Image string `json:"Image"`
	ImageID string `json:"ImageID"`
	Command string `json:"Command"`
	Created uint64 `json:"Created"`
	State string `json:"State"`
	Status string `json:"Status"`
	Ports []Port `json:"Ports"`
	Labels map[string]string `json:"Labels"`
	HostConfig map[string]string `json:"HostConfig"`
	NetworkSettings map[string]interface{} `json:"NetworkSettings"`
	Mounts []Mount `json:"Mounts"`
}

// docker 端口映射
type Port struct {
	IP string `json:"IP"`
	PrivatePort int `json:"PrivatePort"`
	PublicPort int `json:"PublicPort"`
	Type string `json:"Type"`
}

// docker 挂载
type Mount struct {
	Type string `json:"Type"`
	Source string `json:"Source"`
	Destination string `json:"Destination"`
	Mode string `json:"Mode"`
	RW bool `json:"RW"`
	Propatation string `json:"Propagation"`
}

// 连接列表
var SockAddr = "/var/run/docker.sock"
var imagesSock = "GET /images/json HTTP/1.0\r\n\r\n"
var containerSock = "GET /containers/json?all=true HTTP/1.0\r\n\r\n"
var startContainerSock = "POST /containers/%s/start HTTP/1.0\r\n\r\n"

// 白名单
var whiteList []string


func main() {
	// 读取命令行参数
	// 白名单列表
	list := flag.String("list", "", "docker white list to restart, eg: token,explorer")
	// 轮询的时间间隔，单位秒
	times := flag.Int64("time", 10, "time interval to set read docker containers [second], default is 10 second")

	flag.Parse()

	// 解析list => whiteList
	whiteList = strings.Split(*list, ",")

	log.SetOutput(os.Stdout)
	log.Println("start docker watching...")
	log.Printf("Your whiteList: %v\n", *list)
	log.Printf("Your shedule times: %ds\n", *times)

	for {
		// 轮询docker
		err := listenDocker()
		if err != nil {
			log.Println(err.Error())
		}

		time.Sleep(time.Duration(*times)*time.Second)
	}

}

func listenDocker() error {
	// 获取容器列表
	containers, err := readContainer()
	if err != nil {
		return errors.New("read container error: " + err.Error())
	}

	// 先遍历白名单快，次数少
	for _, name := range whiteList {
	Name:
		for _, container := range containers {
			for _, cname := range container.Names {
				// 如果匹配到白名单
				if cname[1:] == name {
					// 关心一下容器状态
					log.Printf("id=%s, name=%s, state=%s", container.Id[:12], container.Names, container.Status)
					if strings.Contains(container.Status, "Exited") {
						// 如果出现异常退出的容器，启动它
						log.Printf("find container: [%s] has exited, ready to start it. ", name)
						e := startContainer(container.Id)
						if e != nil {
							log.Println("start container error: ", e.Error())
						}
						break Name
					}
				}
			}
		}
	}
	return nil
}

// 获取 unix sock 连接
func connectDocker() (*net.UnixConn, error) {
	addr := net.UnixAddr{SockAddr, "unix"}
	return net.DialUnix("unix", nil, &addr)
}

// 启动容器
func startContainer(id string) error {
	conn, err := connectDocker()
	if err != nil {
		return errors.New("connect error: " + err.Error())
	}

	start := fmt.Sprintf(startContainerSock, id)
	fmt.Println(start)
	cmd := []byte(start)
	code, err := conn.Write(cmd)
	if err != nil {
		return err
	}
	log.Println("start container response code: ", code)
	// 启动容器等待20秒，防止数据重发
	time.Sleep(20*time.Second)
	return nil
}

// 获取容器列表
func readContainer() ([]Container, error) {
	conn, err := connectDocker()
	if err != nil {
		return nil, errors.New("connect error: " + err.Error())
	}

	_, err = conn.Write([]byte(containerSock))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	body := getBody(result)

	var containers []Container
	err = json.Unmarshal(body, &containers)
	if err != nil {
		return nil, err
	}

	log.Println("len of containers: ", len(containers))
	if len(containers) == 0 {
		return nil, errors.New("no containers")
	}
	return containers, nil
}

// 获取镜像列表
func readImage(conn *net.UnixConn) ([]Image, error) {
	_, err := conn.Write([]byte(imagesSock))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	body := getBody(result[:])

	var images []Image
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// 从返回的 http 响应中提取 body
func getBody(result []byte) (body []byte) {
	for i:=0; i<=len(result)-4; i++ {
		if result[i] == 13 && result[i+1] == 10 && result[i+2] == 13 && result[i+3] == 10 {
			body = result[i+4:]
			break
		}
	}
	return
}


/*
error log :
	1、write unix @->/var/run/docker.sock: write: broken pipe
		建立的tcp连接不能复用，每次操作都建立连接
 */
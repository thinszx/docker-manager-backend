package agent
//
//import (
//	"container/list"
//	"net"
//)
//
//// DockerAgentLabel 定义了agent的属性，例如master/slave
//type DockerAgentLabel int
//
//// DockerAgentStatus 定义了Agent的状态，online/offline
//type DockerAgentStatus int
//
//
//type DockerAgent struct {
//	Label DockerAgentLabel
//	Address net.IPAddr
//	Name string
//	Status DockerAgentStatus
//	DockerList list.List
//}
//
//func (agent *DockerAgent) CheckStatus(){
//
//}
//
////func
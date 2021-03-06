package main

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/phpcyy/TarsProxy/Local/Tarsproxy/Local"
)

type RegistryObjImp struct {
}

func (imp *RegistryObjImp) Add(a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	*c = a + b
	return 0, nil
}
func (imp *RegistryObjImp) Sub(a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	*c = a * b
	return 0, nil
}

func (imp *RegistryObjImp) FindObjectById(Id string) (ret []Local.EndpointF, err error) {
	fmt.Println("FindObjectById")
	return findObjectById(Id), nil
}

func (imp *RegistryObjImp) FindObjectById4Any(Id string, ActiveEp *[]Local.EndpointF, InactiveEp *[]Local.EndpointF) (int32, error) {
	*ActiveEp = findObjectById(Id)
	return 0, nil
}

func (imp *RegistryObjImp) FindObjectById4All(Id string, ActiveEp *[]Local.EndpointF, InactiveEp *[]Local.EndpointF) (int32, error) {
	*ActiveEp = findObjectById(Id)
	return 0, nil
}

func (imp *RegistryObjImp) FindObjectByIdInSameGroup(Id string, ActiveEp *[]Local.EndpointF, InactiveEp *[]Local.EndpointF) (int32, error) {
	fmt.Println("FindObjectByIdInSameGroup")

	*ActiveEp = findObjectById(Id)
	return 0, nil
}

func (imp *RegistryObjImp) FindObjectByIdInSameStation(Id string, SStation string, ActiveEp *[]Local.EndpointF, InactiveEp *[]Local.EndpointF) (int32, error) {
	*ActiveEp = findObjectById(Id)
	return 0, nil
}

func (imp *RegistryObjImp) FindObjectByIdInSameSet(Id string, SetId string, ActiveEp *[]Local.EndpointF, InactiveEp *[]Local.EndpointF) (int32, error) {
	*ActiveEp = findObjectById(Id)
	return 0, nil
}

func findObjectById(Id string) (ret []Local.EndpointF) {
	cfg := tars.GetServerConfig()

	var tarsGateWay Local.EndpointF
	tarsGateWay.Host = cfg.Adapters["Local.Tarsproxy.localTcpProxy"].Endpoint.Host
	tarsGateWay.Port = cfg.Adapters["Local.Tarsproxy.localTcpProxy"].Endpoint.Port
	tarsGateWay.Istcp = 1

	ret = append(ret, tarsGateWay)
	println("findObjectById ret ", ret[0].Host, ret[0].Port)
	return ret
}

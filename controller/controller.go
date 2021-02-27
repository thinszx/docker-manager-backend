package controller

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/dockermanage/model"
	"time"
)

func CheckDockerAgentOnline(agent *model.HostAgentModel) {
	//if
}

func ListHostDocker() {
	timeLayout := "2000-00-00 00:00:00"
	ilo := types.ImageListOptions{
		All: false,
		Filters: filters.Args{

		},
	}
	cli, _ := client.NewClientWithOpts(client.FromEnv)
	is, _ := cli.ImageList(context.Background(), ilo)
	for _, i := range is {
		//fmt.Printf("%T",i)
		fmt.Println("容器:", i.Containers)
		fmt.Println("创建时间:", time.Unix(i.Created, 0).Format(timeLayout))
		fmt.Println("ID:", i.ID)
		for _, l := range i.Labels {
			fmt.Println("Labels:", l)
		}
		fmt.Println("ParentId:", i.ParentID)
		//fmt.Println("标签:",i.Labels)
		for _, r := range i.RepoDigests {
			fmt.Println("RepoDigests", r)
		}
		for _, r := range i.RepoTags {
			fmt.Println("repository:tag:", r)
		}

		fmt.Println("大小:", i.Size/(1000*1000))
	}
}

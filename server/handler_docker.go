package server

import (
	"errors"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/znkisoft/watchdog/schema"
)

func ContainerList(_ http.ResponseWriter, r *http.Request) (any, error) {
	req := &schema.ListContainersRequest{}

	// 	types.ContainerListOptions{
	// 	Quiet:   false,
	// 	Size:    false,
	// 	All:     false,
	// 	Latest:  false,
	// 	Since:   "",
	// 	Before:  "",
	// 	Limit:   0,
	// 	Filters: filters.Args{},
	// }

	if err := UnmarshalReq(r, req); err != nil {
		return nil, err
	}

	arr := make([]*types.Container, 0)
	arr = append(arr, &types.Container{
		Names: []string{"test"},
	})

	return arr, nil
}

func ContainerCreate(_ http.ResponseWriter, r *http.Request) (any, error) {
	return nil, errors.New("not implemented")
}
func ContainerRemove(_ http.ResponseWriter, r *http.Request) (any, error) {
	return nil, errors.New("not implemented")
}
func ContainerRename(_ http.ResponseWriter, r *http.Request) (any, error) {
	return nil, errors.New("not implemented")
}
func ContainerStart(_ http.ResponseWriter, r *http.Request) (any, error) {
	return nil, errors.New("not implemented")
}
func ContainerStop(_ http.ResponseWriter, r *http.Request) (any, error) {
	return nil, errors.New("not implemented")
}
func ContainerRestart(_ http.ResponseWriter, r *http.Request) (any, error) {
	return nil, errors.New("not implemented")
}
func ContainerKill(_ http.ResponseWriter, r *http.Request) (any, error) {
	return nil, errors.New("not implemented")
}
func ContainersPrune(_ http.ResponseWriter, r *http.Request) (any, error) {
	return nil, errors.New("not implemented")
}

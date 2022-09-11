package server

import (
	"errors"
	"net/http"

	"github.com/docker/docker/api/types"
)

func ContainerList(_ http.ResponseWriter, r *http.Request) (any, error) {

	arr := make([]*types.Container, 0)
	arr = append(arr, &types.Container{
		Names: []string{"test"},
	})

	return nil, nil
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

package docker

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

type API interface {
	Ping(host string) (bool, error)
	ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	ContainerCreate(ctx context.Context, config *ContainerConfig, hostConfig *HostConfig, networkingConfig *NetworkingConfig, containerName string) (ContainerCreateCreatedBody, error)
	ContainerRemove(ctx context.Context, container string, options types.ContainerRemoveOptions) error
	ContainerRename(ctx context.Context, container, newContainerName string) error
	ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error
	ContainerStop(ctx context.Context, container string, timeout *time.Duration) error
	ContainerRestart(ctx context.Context, container string, timeout *time.Duration) error
	ContainerKill(ctx context.Context, container, signal string) error
	ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
	// ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
	// ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
	// ContainerDiff(ctx context.Context, container string) ([]containertypes.ContainerChangeResponseItem, error)
	// ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error)
	// ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
	// ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
	// ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
	// ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
	// ContainerExport(ctx context.Context, container string) (io.ReadCloser, error)
	// ContainerInspect(ctx context.Context, container string) (types.ContainerJSON, error)
	// ContainerInspectWithRaw(ctx context.Context, container string, getSize bool) (types.ContainerJSON, []byte, error)
	// ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
	// ContainerPause(ctx context.Context, container string) error
	// ContainerResize(ctx context.Context, container string, options types.ResizeOptions) error
	// ContainerStatPath(ctx context.Context, container, path string) (types.ContainerPathStat, error)
	// ContainerStats(ctx context.Context, container string, stream bool) (types.ContainerStats, error)
	// ContainerStatsOneShot(ctx context.Context, container string) (types.ContainerStats, error)
	// ContainerTop(ctx context.Context, container string, arguments []string) (containertypes.ContainerTopOKBody, error)
	// ContainerUnpause(ctx context.Context, container string) error
	// ContainerUpdate(ctx context.Context, container string, updateConfig containertypes.UpdateConfig) (containertypes.ContainerUpdateOKBody, error)
	// ContainerWait(ctx context.Context, container string, condition containertypes.WaitCondition) (<-chan containertypes.ContainerWaitOKBody, <-chan error)
	// CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error)
	// CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error
}

type Client struct {
}

func (c Client) Ping(host string) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (c Client) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	// TODO implement me
	panic("implement me")
}

func (c Client) ContainerCreate(ctx context.Context, config *ContainerConfig, hostConfig *HostConfig, networkingConfig *NetworkingConfig, containerName string) (ContainerCreateCreatedBody, error) {
	// TODO implement me
	panic("implement me")
}

func (c Client) ContainerRemove(ctx context.Context, container string, options types.ContainerRemoveOptions) error {
	// TODO implement me
	panic("implement me")
}

func (c Client) ContainerRename(ctx context.Context, container, newContainerName string) error {
	// TODO implement me
	panic("implement me")
}

func (c Client) ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error {
	// TODO implement me
	panic("implement me")
}

func (c Client) ContainerStop(ctx context.Context, container string, timeout *time.Duration) error {
	// TODO implement me
	panic("implement me")
}

func (c Client) ContainerRestart(ctx context.Context, container string, timeout *time.Duration) error {
	// TODO implement me
	panic("implement me")
}

func (c Client) ContainerKill(ctx context.Context, container, signal string) error {
	// TODO implement me
	panic("implement me")
}

func (c Client) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error) {
	// TODO implement me
	panic("implement me")
}

# watchdog

## package info
- `exec`: execute commands
- `zfs`: file system
- `zssh`: ssh interface
- `ssm`: server status monitor

### SSM

- [ ] self-monitoring
- [ ] check other servers using ssh



## REST API

### SSM

- `GET /api/v1/ssm/list` get plugin list for server status monitor
- `GET /api/v1/ssm/[:plugins]` get plugin status for server status monitor

<details>
    <summary>plugin list</summary>

- [ ] Ping
- [ ] Uptime
- [ ] CPU
- [ ] Memory
- [ ] Load
- [ ] Process list
- [ ] Network interface
- [ ] Disk I/O
- [ ] IRQ / Raid
- [ ] Sensors
- [ ] Filesystem(and folders)
- [ ] Monitor
- [ ] Alert
</details>

## Credits:
- [glances](https://github.com/nicolargo/glances)
- [gossm](https://github.com/ssimunic/gossm)
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

- uptime
- 
</details>

## Credits:
- [glances](https://github.com/nicolargo/glances)
- [gossm](https://github.com/ssimunic/gossm)
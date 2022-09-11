# watchdog

## package info

- `exec`: execute commands
- `zfs`: file system
- `zssh`: ssh interface
- `ssm`: server status monitor

## REST APIs

### Userver

- `GET /api/v1/userver`: get all userver
- `GET /api/v1/userver/:id`: get userver by id
- `POST /api/v1/userver`: create userver
- `PUT /api/v1/userver/:id`: update userver
- `DELETE /api/v1/userver/:id`: delete userver

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

### Docker

## Credits:

- [glances](https://github.com/nicolargo/glances)
- [pocketbase](https://github.com/pocketbase/pocketbase)
- [gossm](https://github.com/ssimunic/gossm)
# Ports

Ports exposed by HIoT Platform.

To modify a port, update `docker/.env` and to be on the safe side, search the code base for any hard-coded values.

| Port | Description                                                                              |
|------|------------------------------------------------------------------------------------------|
| 3000 | Auth - HTTP port.                                                                        |
| 8086 | InfluxDB - Web Port.                                                                     |
| 8125 | Telegraf - StatsD Port.                                                                  |
| 8443 | Kong - Takes incoming HTTP traffic from Consumers, and forwards it to upstream Services. |
| 8444 | Kong - Admin API. Listens for calls from the command line over HTTP.                     |
| 9000 | Grafana - Web Port.                                                                      |

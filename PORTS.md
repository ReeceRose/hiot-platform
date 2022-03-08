# Ports

Ports exposed by HIoT Platform.

To modify a port, update `docker/.env` and to be on the safe side, search the code base for any hard-coded values.

| Port | Description                                                                              |
|------|------------------------------------------------------------------------------------------|
| 8000 | Kong - Takes incoming HTTP traffic from Consumers, and forwards it to upstream Services. |
| 8001 | Kong - Admin API. Listens for calls from the command line over HTTP.                     |
| 8086 | InfluxDB - Web Port                                                                      |
| 8125 | Telegraf - StatsD Port                                                                   |
| 9000 | Grafana - Web Port                                                                       |

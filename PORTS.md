# Ports

Ports exposed by HIoT Platform.

To modify a port, update `docker/.env` and to be on the safe side, search the code base for any hard-coded values.

| Port | Description                    |
|------|--------------------------------|
| 80   | Traefik - Incoming HTTP port.  |
| 443  | Traefik - Incoming HTTPS port. |
| 8080 | Traefik - Dashboard port.      |
| 8086 | InfluxDB - Web Port            |
| 8125 | Telegraf - StatsD Port         |
| 9000 | Grafana - Web Port             |

# Environment Variables

Environment variables used by HIoT Platform

| Environment Variable    | Description                    | Default Value |
|-------------------------|--------------------------------|------------------------|
| TELEGRAF_VERSION        | Telegraf container version     | 1.12.4                 |
| TELEGRAF_CONTAINER_NAME | Telegraf docker container name | hiot-platform-telegraf |
| TELEGRAF_STATSD_PORT    | Telegraf StatsD Port           | 8125                   |
| INFLUXDB_VERSION        | InfluxDB container version     | 2.1.2                  |
| INFLUXDB_CONTAINER_NAME | InfluxDB docker container name | hiot-platform-influxdb |
| INFLUXDB_INIT_MODE      | InfluxDB init mode             | setup                  |
| INFLUXDB_USERNAME       | InfluxDB username              | admin                  |
| INFLUXDB_PASSWORD       | InfluxDB password              | password               |
| INFLUXDB_ORG            | InfluxDB organization          | hiot-platform          |
| INFLUXDB_BUCKET         | InfluxDB default bucket        | data                   |
| INFLUXDB_ADMIN_TOKEN    | InfluxDB admin token           | hiot-platform-token    |
| INFLUXDB_WEB_PORT       | InfluxDB web port              | 8086                   |
| GRAFANA_VERSION         | Grafana container version      | 8.4.2                  |
| GRAFANA_CONTAINER_NAME  | Grafana docker container name  | hiot-platform-grafana  |
| GRAFANA_USERNAME        | Grafana username               | admin                  |
| GRAFANA_PASSWORD        | Grafana password               | password               |
| GRAFANA_WEB_PORT        | Grafana web port               | 9000                   |
| TRAEFIK_VERSION         | Traefik container version      | 2.6                    |
| TRAEFIK_CONTAINER_NAME  | Traefik container name         | hiot-platform-traefik  |
| TRAEFIK_HTTP_PORT       | Traefik incoming HTTP port     | 80                     |
| TRAEFIK_HTTPS_PORT      | Traefik incoming HTTPS port    | 443                    |
| TRAEFIK_DASHBOARD_PORT  | Traefik dashboard port         | 8080                   |
# Environment Variables

Environment variables used by HIoT Platform

| Environment Variable    | Description                         | Default Value                    |
|-------------------------|-------------------------------------|----------------------------------|
| CERTS_DIR               | Path to certs directory             | ../certs                         |
| TELEGRAF_VERSION        | Telegraf container version          | 1.12.4                           |
| TELEGRAF_CONTAINER_NAME | Telegraf docker container name      | hiot-platform-telegraf           |
| TELEGRAF_STATSD_PORT    | Telegraf StatsD Port                | 8125                             |
| INFLUXDB_VERSION        | InfluxDB container version          | 2.1.2                            |
| INFLUXDB_CONTAINER_NAME | InfluxDB docker container name      | hiot-platform-influxdb           |
| INFLUXDB_INIT_MODE      | InfluxDB init mode                  | setup                            |
| INFLUXDB_USERNAME       | InfluxDB username                   | admin                            |
| INFLUXDB_PASSWORD       | InfluxDB password                   | password                         |
| INFLUXDB_ORG            | InfluxDB organization               | hiot-platform                    |
| INFLUXDB_BUCKET         | InfluxDB default bucket             | data                             |
| INFLUXDB_ADMIN_TOKEN    | InfluxDB admin token                | hiot-platform-token              |
| INFLUXDB_WEB_PORT       | InfluxDB web port                   | 8086                             |
| GRAFANA_VERSION         | Grafana container version           | 8.4.2                            |
| GRAFANA_CONTAINER_NAME  | Grafana docker container name       | hiot-platform-grafana            |
| GRAFANA_USERNAME        | Grafana username                    | admin                            |
| GRAFANA_PASSWORD        | Grafana password                    | password                         |
| GRAFANA_WEB_PORT        | Grafana web port                    | 9000                             |
| KONG_VERSION            | Kong container version              | 2.7.1                            |
| KONG_CONTAINER_NAME     | Kong docker container name          | hiot-platform-kong               |
| KONG_LOG                | Kong log location                   | /dev/stdout                      |
| KONG_CONSUMER_PORT      | Kong consumer port                  | 8443                             |
| KONG_ADMIN_API_PORT     | Kong API port                       | 8444                             |
| AUTH_VERSION            | Auth microservice container version | latest                           |
| AUTH_CONTAINER_NAME     | Auth microservice container name    | hiot-platform-auth               |
| AUTH_JWT_SECRET         | JWT signing key                     | u8x/A?D(G+KbPeShVmYq3t6v9y$B&E)H |
| AUTH_HTTP_PORT          | Auth microservice HTTP port         | 3000                             |

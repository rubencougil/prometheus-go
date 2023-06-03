# 📶 Prometheus/Grafana Docker + Sample Go App

Monitoring System build using **Prometheus** (with some companion modules like **Alert Manager** and **MySQL Exporter**). It uses **Grafana** as the visualization dashboard.

[![Actions Status](https://github.com/rubencougil/prometheus-go/workflows/Build/badge.svg)](https://github.com/rubencougil/prometheus-go/actions)
[![Build Release](https://github.com/rubencougil/prometheus-go/workflows/Build%20Release/badge.svg)](https://github.com/rubencougil/prometheus-go/actions)

![Clase 3 - Practica Prometheus](https://github.com/rubencougil/prometheus-go/assets/1073799/b4d7ba04-cbf2-4e06-b072-d0cbf9bf1974)

## How to run it?

To spin up needed services

`docker-compose up -d`

Once done, the Sample Go App will automatically send some metrics to Prometheus. Can be visualized in **Grafana** at:

```
http://localhost:3000
User: admin
Pass: foobar
```

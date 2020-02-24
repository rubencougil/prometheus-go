# 📶 Prometheus/Grafana Docker + Sample Go App

Monitoring System build using **Prometheus** (with some companion modules like **CAAdvisor**, **Node Exporter**, **Alert Manager** and **MySQL Exporter**). It uses **Grafana** as the visualization dashboard.

![Clase 3 - Practica Prometheus](https://user-images.githubusercontent.com/1073799/75155005-52976180-570f-11ea-9163-b7af42a03349.jpg)

## How to run it?

To spin up needed services

`docker-compose up -d`

Once done, the Sample Go App will automatically send some metrics to Prometheus. Can be visualized in **Grafana** at:

```
http://localhost:3000
User: admin
Pass: foobar
```

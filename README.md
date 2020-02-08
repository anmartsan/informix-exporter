# Informix Exporter Prometheus

Prometheus exporter para varias metricas de Informix escrito en GO. 

### Pre-requisitos 📋

- Docker y docker-compose.
- Informix CSDK
- unixodbc

### Instalación 🔧

Configuramos el fichero sqlhosts :

prueba		onsoctcp	192.168.1.50	1527
prueba2		onsoctcp	192.168.1.50	1530

Configuramos el fichero odbc.ini

[ODBC]
UNICODE=UCS-2
[prueba]
Driver=/opt/IDS12/lib/cli/libifcli.so
Server=prueba
Database=sysmaster
TRANSLATIONDLL=/opt/IDS12/lib/esql/igo4a304.so
LogonID=informix
pwd=password
[prueba2]
Driver=/opt/IDS12/lib/cli/libifcli.so
Server=prueba2
Database=sysmaster
TRANSLATIONDLL=/opt/IDS12/lib/esql/igo4a304.so
LogonID=informix
pwd=password

Configuramos el fichero server.yaml
---
servers:
- name: pruebaids
  informixserver: prueba
  user: informix
  password: password
- name: pruebaids2
  informixserver: prueba2
  user: informix
  password: password
  
  Configuramos el fichero config.yml con querys que queramos ejecutar.
  
  metrics:
- parametro: logs_without_backup
  type: gauge
  description: Number of logs without backup   
  query: select count(*) as logssinbackup from syslogs where is_backed_up=0  
  label: custom
- parametro: pf_bufreads
  type: gauge
  description: Number of pf_bufreads   
  query: select value  from sysshmhdr where name ="pf_bufreads"
  label: custom
  
  ## Arranque del sistema ⚙️

```
docker-compose up -d
./informix-exporter
```

## Autores ✒️



* **Antonio Martinez Sanchez-Escalonilla ** - [anmartsan](https://github.com/anmartsan)
    www.scmsi.es







---
⌨️ con ❤️ por [anmartsan](a.martinez@scmsi.es) 😊



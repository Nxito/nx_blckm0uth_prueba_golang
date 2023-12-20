> [!IMPORTANT]  
> Ante la falta de comunicacion con esta empresa,
> publico la prueba técnica. La cual me hicieron realizar en menos de 48 horas y otorgando nulo feedback incluso tras haber pedido una minima confirmacion
> 




# Survivor Battle Royale API

# Index <!-- omit in toc -->

- [Survivor Battle Royale API](#survivor-battle-royale-api)
  - [Prerequisites](#prerequisites)
  - [How to configure](#how-to-configure)
    - [Docker-compose](#docker-compose)
    - [Config.json](#configjson)
    - [Data at start](#data-at-start)
  - [How to start](#how-to-start)

## Prerequisites

- Docker Compose

## How to configure

### Docker-compose

- app: You can change the endpoint port.
- database: You can change the **port** , **user** and **password**. They must be configured in config.json as same as here.

### Config.json

On the file [config.json](app/config.json) you can configure the following parametters:

| Key           | Description        |
| ------        | ---------------    |
| database      |  Parameters to connect to PostgreeSQL database|
| max_queues    | Value to limit the Queues on the database|

### Data at start

Inside [1_dump_data.sql](app/data/postgres/sql/1_dump_data.sql) you can use SQL functions that be initialized at start
You can specify more files that be executed in aphabetic order.

## How to start

To start the proyect just use

`docker compose -f "docker-compose.yml" up -d --build`

*note: if you are using the old pluging, you must to uso docker-compose command*

Docker will install an app container and a postgree container.

When initialized, go to [localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) to check the api urls

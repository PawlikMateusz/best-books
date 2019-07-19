# best-books

The project used to learning Go and practice a microservice approach for websites. Please open *./overview.jpg* to become familiar with the software architecture

## Project technology stack

Backend

* **Go** as main backend language
* **Nginx** as api gateway
* **RabbitMQ** is used for internal between services communication
* **MongoDB** as main document non-relational database

Frontend

* **Vue.js** JavaScript Framework
* **REST** software architectural style for *frontend -> backend* communication

Other

* **Docker** as container platform
* **docker-compose** used for defining and running multi-container Docker applications

## Quick Start Guide

Required software

* Docker
* docker-compose

Running project locally

1. Build docker containers

    ```sh
    docker-compose build
    ```

2. Run all services

    ```sh
    docker-compose up
    ```

3. To visit site open <http://localhost:80>

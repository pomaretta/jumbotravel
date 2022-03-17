# JumboTravel Infrastructure

## MainDB

Create a volume to store the main database.
```
docker volume create --name jumbotravel-infrastructure-maindb
```

Info about the volume:
```
docker volume inspect jumbotravel-infrastructure-maindb
```

Create compose file
```yaml
version: "3.7"

services:
    database:
        image: arm64v8/mysql:8.0-oracle
        container_name: jumbotravel-infrastructure-maindb
        restart: unless-stopped
        volumes:
            - jumbotravel-infrastructure-maindb:/var/lib/mysql
        environment:
            MYSQL_ROOT_PASSWORD: root
            MYSQL_DATABASE: jumbotravel
            MYSQL_USER: jumbotravel
            MYSQL_PASSWORD: jumbotravel
        network_mode: host
        expose:
            - 3306

volumes:
    jumbotravel-infrastructure-maindb:
        external: true
```
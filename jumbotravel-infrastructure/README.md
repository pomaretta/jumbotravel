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
        image: postgres:14
        container_name: jumbotravel-infrastructure-maindb
        restart: unless-stopped
        volumes:
            - jumbotravel-infrastructure-maindb:/var/lib/postgresql/data
        environment:
            - POSTGRES_DB=jumbotravel
            - POSTGRES_USER=jumbotravel
            - POSTGRES_PASSWORD=jumbotravel
            - PGDATA=/var/lib/postgresql/data/pgdata
        network_mode: host
        expose:
            - 5432

volumes:
    jumbtravel-infrastructure-maindb:
        external: true
```
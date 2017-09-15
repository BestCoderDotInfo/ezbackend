---
title: Setup single Docker Postgresql for Rails application on AWS
date: '2017-08-07 22:00:00'
comments: true
external-url: null
categories: internet
keywords: docker, rails, postgresql, docker postgresql
excerpt: When I have been deploy my Rails application on AWS, my Dockerfile is containt Rails app, nginx, redis and postgresql. It work fine, but when I have deploy new sersion for my app, sometime I forgot and rebuild new containers and database  is reset :( . I think I need one Database server look like Amazon RDS.
---

When I have been deploy my Rails application on AWS, my Dockerfile is containt Rails app, nginx, redis and postgresql. It work fine, but when I have deploy new sersion for my app, sometime I forgot and rebuild new containers and database  is reset :( . I think I need one Database server look like Amazon RDS. 

### Geting start

#### 1. Simple

Youn setup `PostgresSQL server with Docker` step by step with resource: [Dockerize PostgreSQL](https://docs.docker.com/engine/examples/postgresql_service/).

After setup successfully.

**Conection**

First, run `docker ps`

```
$ docker ps

CONTAINER    ID    IMAGE   COMMAND   CREATED    STATUS        PORTS       NAMES

5e24362f27f6   eg_postgresql:latest   /usr/lib/postgresql/   About an hour ago   Up About an hour    0.0.0.0:49153->5432/tcp               pg_test

```

We see PORT : **49157**

Conection setup:

- IP adress: your server IP
- Port: 49157
- user: docker
- pass: docker

Note: Wee need port **49157** to `security group` if you using **AWS EC2**

**Tab Inbound** : 

Type | Protocol | Port range | Source
-----|----------|------------|-------
Custom TCP Rule | TCP | 49157 | ::/0
Custom TCP Rule | TCP | 49157 | 0.0.0.0/0

**Stop docker image:**

```
docker stop image-name
```

#### Advance

At here we have `Dockerfile`:

```bash
FROM ubuntu

ARG DB_USER=slackbot
ARG DB_PASS=slackbot

# Add the PostgreSQL PGP key to verify their Debian packages.
# It should be the same key as https://www.postgresql.org/media/keys/ACCC4CF8.asc
RUN apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8

# Add PostgreSQL's repository. It contains the most recent stable release
#     of PostgreSQL, ``9.6``.
RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ precise-pgdg main" > /etc/apt/sources.list.d/pgdg.list

# Install ``python-software-properties``, ``software-properties-common`` and PostgreSQL 9.6
#  There are some warnings (in red) that show up during the build. You can hide
#  them by prefixing each apt-get statement with DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y python-software-properties software-properties-common postgresql-9.6 postgresql-client-9.6 postgresql-contrib-9.6

# Note: The official Debian and Ubuntu images automatically ``apt-get clean``
# after each ``apt-get``

# Run the rest of the commands as the ``postgres`` user created by the ``postgres-9.6`` package when it was ``apt-get installed``
USER postgres

# Create a PostgreSQL role named ``docker`` with ``docker`` as the password and
# then create a database `docker` owned by the ``docker`` role.
# Note: here we use ``&&\`` to run commands one after the other - the ``\``
#       allows the RUN command to span multiple lines.
RUN    /etc/init.d/postgresql start &&\
    psql --command "CREATE USER ${DB_USER} WITH SUPERUSER PASSWORD '${DB_PASS}';" &&\
    createdb -O ${DB_USER} ${DB_PASS}

# Adjust PostgreSQL configuration so that remote connections to the
# database are possible.
RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/9.6/main/pg_hba.conf

# And add ``listen_addresses`` to ``/etc/postgresql/9.6/main/postgresql.conf``
RUN echo "listen_addresses='*'" >> /etc/postgresql/9.6/main/postgresql.conf

# Expose the PostgreSQL port
EXPOSE 5432

# Add VOLUMEs to allow backup of config, logs and databases
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

# Set the default command to run when starting the container
CMD ["/usr/lib/postgresql/9.6/bin/postgres", "-D", "/var/lib/postgresql/9.6/main", "-c", "config_file=/etc/postgresql/9.6/main/postgresql.conf"]
```

You just care with somethings:

```bash
ARG DB_USER=slackbot
ARG DB_PASS=slackbot
```

Default `DB_USER` and `DB_PASS`.

After commit all your changes and push to your repo (Github, Gitlab,...). We will using auto deploy with Capistrano.

You must install `Ruby`. Open terminal and run:

```
cp config/application.yml.sample config/application.yml
```

You need config for your application at `config/application.yml`:

```yml
APP_NAME:
REPO_URL: git@github.com:minhquan4080/slackapp-postgresql-server.git
SERVER_IP:
SERVER_DOMAIN:
DEPLOY_USER:
DB_USER:
DB_PASS:
IMAGE_TAG:
PORT: '5432:5432'
```

- `DEPLOY_USER`: user of your Compute Cloud (AWS, Digital Ocean,...).

- Your Compute Cloud, you must install Docker for it. I recommend you should use Ubuntu OS for your Compute Cloud.

### Getting Start

Install Ruby gems:

```
gem install bundler && bundler install
```

After config complete.

```
cap production deploy # For production env
```

After deployed successfully. We can use database url look like:

`postgres://DB_USER:DB_USER_PASS@SERVER_IP:5432/database-name`

Change `DB_USER`, `DB_PASS` and `SERVER_IP` with your config. If you have been config `PORT: '5432:5432'`, you just use **5432**.

`database-name` do not exist, you need create it with `Rails migration` (If you using with Rails application) or you can create with SQL query.

[Source](https://github.com/minhquan4080/slackapp-postgresql-server)


Look Good! Now, you can use it to hosting dabase for your application :smile:.
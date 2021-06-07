# UOJ judging system refactor

## Current architecture

```text
judger_client -poll-> [web]judge/submit
```

## Planned architecture

```text
judger_client <-push(streaming)-> judger_server -(pool / push)-> [web]http endpoint / database
```

## TODO

- [x] Deprecate V8JS
- [x] SQL docker
- [x] Web docker
  - [x] Fix PHP version
- [ ] Nginx reverse proxy docker
- [ ] Web refactor
  - [ ] Remove exec
  - [ ] Remove run_program dependency
  - [ ] Remove main_judger
  - [ ] Push judger request
- [ ] Implements judger_server
- [ ] Implements judger_client
- [ ] Implements websocket push server

## Command Save

build

```bash
docker build -t uoj_web -f Dockerfile.web .
docker run -it -p 4567:80 --rm --link uoj_mysql uoj_web

docker build -t uoj_render_server -f Dockerfile.render_server .
```

mysql

```bash
docker run -it --rm --name uoj_mysql -e MYSQL_ROOT_PASSWORD=test -p 3306:3306 \
    -v $(pwd)/install/db/app_uoj233.sql:/docker-entrypoint-initdb.d/app_uoj233.sql \
    -v $(pwd)/config/uoj_mysqld.cnf:/etc/mysql/conf.d/uoj_mysqld.cnf \
    mysql
```

render server

```bash
docker run -it --rm --name uoj_render_server -p 3456:3456 \
 -v $(pwd)/src:/render_server/src \
 -v $(pwd)/package.json:/render_server/package.json \
 -v $(pwd)/package-lock.json:/render_server/package-lock.json \
  node /bin/bash

```

web:

```bash
docker run -it --rm --link uoj_mysql --link uoj_render_server \
  -p 4567:80 \
  -v $(pwd)/config/000-uoj.conf:/etc/apache2/sites-available/000-uoj.conf \
  -v $(pwd)/judger:/opt/uoj/judger \
  -v $(pwd)/web:/var/www/uoj \
  -v $(pwd)/config/conf.php:/var/www/uoj/app/.config.php \
  php:apache /bin/bash

apt-get update \
 && apt-get install -y \
 build-essential \
 libglib2.0-dev \
 zip \
 unzip \
 libzip-dev \
 libapache2-mod-xsendfile \
 libyaml-dev \
 && rm -rf /var/lib/apt/lists/* \
 && pecl install yaml \
 && a2enmod rewrite headers \
 && docker-php-ext-configure zip \
 && docker-php-ext-install mysqli \
 && docker-php-ext-enable yaml 

a2ensite 000-uoj.conf \
 && a2dissite 000-default.conf \
 && sed -i -e '172s/AllowOverride None/AllowOverride All/' /etc/apache2/apache2.conf \
 && mkdir -p --mode 733 /var/lib/php/uoj_sessions \
 && chmod +t /var/lib/php/uoj_sessions \
 && mkdir -p /var/www/uoj/app/storage \
 && chown -R www-data:www-data /var/www/uoj/app/storage \
 && mkdir -p /var/uoj_data/upload \
 && chown -R www-data:www-data /var/uoj_data

apache2-foreground
```

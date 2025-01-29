FROM php:8.2-apache
RUN apt-get update && docker-php-ext-install pdo_mysql mysqli
WORKDIR /var/www/html
COPY index.php index.php


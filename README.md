# Creating the MySql database for development.

To create the MySql database for development purposes use the following command.

 ```
 docker run -d \
 -p 3306:3306 \
 --name mysql \
 -e MYSQL_ROOT_PASSWORD=root \
 -e MYSQL_DATABASE=fitcal \
 -e MYSQL_USER=root \
 -e MYSQL_PASSWORD=root \
 mysql/mysql-server:latest
 ```
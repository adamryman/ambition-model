#! /bin/bash


function mysql_docker {
	mysql='exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -D"ambition" -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD" -e '
	docker run --link ambition-mysql:mysql --rm mysql sh -c "$mysql \"$1\""
}


# Stop mysql
docker stop ambition-mysql
docker rm ambition-mysql 

# Start mysql
docker run --name ambition-mysql -e MYSQL_ROOT_PASSWORD=ambition -e MYSQL_DATABASE=ambition -e MYSQL_USER=ambition -e MYSQL_PASSWORD=ambition -p 3306:3306 -d mysql

sleep 10

while read line; do
	mysql_docker "$line"
done < createTables.sql


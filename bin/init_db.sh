#!/bin/bash
cd docker || exit
docker compose cp ./mysql/sql mysql:/var/lib/mysql/
docker compose exec mysql bash -c "mysql -uroot -ppassword < /var/lib/mysql/sql/reset_database.sql"

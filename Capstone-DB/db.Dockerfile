FROM mysql:8.0.28

# Copy files into capstone-db container
COPY ./database/mysql.cnf /etc/mysql/conf.d/

COPY ./database/sample.csv /input/

COPY ./database/*.sql /docker-entrypoint-initdb.d/

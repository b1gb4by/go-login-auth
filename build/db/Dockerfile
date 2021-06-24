FROM mysql:8.0.25

ENV MYSQL_ROOT_PASSWORD=admin
ENV MYSQL_USER=admin
ENV MYSQL_PASSWORD=admin
ENV MYSQL_DATABASE=go_auth

ADD ./my.cnf /etc/mysql/conf.d/my.cnf
CMD ["mysqld"]

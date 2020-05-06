\c workermonitordb;

create table client_data(
    client_id int,
    name varchar(20),
    startup_datetime timestamp,
    shutdown_datetime timestamp,
    primary key(client_id, startup_datetime)
);

insert into client_data values (
1,
'Taro Sato',
'2020/04/30 10:00:00',
'2020/04/30 20:00:00'
);
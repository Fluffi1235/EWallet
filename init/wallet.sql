CREATE table wallets(
    id varchar(100) primary key,
    balance bigint default 10000
);

create table transactions_history(
    fromId varchar(100) not null,
    toId varchar(100) not null,
    amount bigint,
    time timestamp
)

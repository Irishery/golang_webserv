CREATE TABLE currency (
    id bigserial not null primary key,
    symbol varchar not null,
    price float8 not null,
    volume float8 not null,
    last_trade float8 not null
)

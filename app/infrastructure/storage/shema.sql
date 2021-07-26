CREATE DATABASE IF NOT EXISTS egame;

CREATE TABLE IF NOT EXISTS egame.events
(
    client_time String,
    device_id String,
    device_os String,
    session String,
    sequence UInt8,
    event String,
    param_int UInt8,
    param_str String,
    ip String,
    server_time DateTime
)
    ENGINE = MergeTree()
    PRIMARY KEY(device_id, session, sequence, server_time);
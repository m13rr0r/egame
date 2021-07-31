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

CREATE TABLE egame.events_buffer AS egame.events
    ENGINE = Buffer(egame, events, 16, 10, 100, 10000, 1000000, 10000000, 100000000)
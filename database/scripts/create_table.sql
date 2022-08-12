create table if not exists pixel
(
    source_id     varchar not null,
    datatype_id   int     not null,
    sourcetype_id int     not null,
    data_inst     varchar not null
);
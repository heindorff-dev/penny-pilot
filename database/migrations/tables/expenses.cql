create table expense (
    id uuid primary key,
    name text,
    amount int
)

CREATE CUSTOM INDEX fn_prefix ON expense (name)
USING 'org.apache.cassandra.index.sasi.SASIIndex';
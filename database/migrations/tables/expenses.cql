create table expense (
    id uuid primary key,
    name text,
    amount int
)

CREATE CUSTOM INDEX fn_prefix ON expense (name)
USING 'org.apache.cassandra.index.sasi.SASIIndex';

CREATE CUSTOM INDEX idx_name ON penny_pilot.expense (name) 
USING 'org.apache.cassandra.index.sasi.SASIIndex' 
WITH OPTIONS = {
'mode': 'CONTAINS', 
'analyzer_class': 'org.apache.cassandra.index.sasi.analyzer.NonTokenizingAnalyzer', 
'case_sensitive': 'false'};
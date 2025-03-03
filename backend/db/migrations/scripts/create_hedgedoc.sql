-- create-hedgedoc.sql

-- Creates a separate database for HedgeDoc
CREATE DATABASE IF NOT EXISTS hedgedoc;

-- Create a dedicated MySQL user for HedgeDoc
CREATE USER IF NOT EXISTS 'hedgedoc'@'%' IDENTIFIED BY 'hedgedoc_pass';

-- Give that user privileges on just the hedgedoc DB
GRANT ALL PRIVILEGES ON hedgedoc.* TO 'hedgedoc'@'%';

FLUSH PRIVILEGES;

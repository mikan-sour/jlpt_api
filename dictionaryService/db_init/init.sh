#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$POSTGRES_DB"<<-EOSQL
   CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

   CREATE TABLE IF NOT EXISTS words 
      (
         id SERIAL PRIMARY KEY,
         foreign1 VARCHAR(255) NOT NULL, 
         foreign2 VARCHAR(255) NOT NULL,
         definitions TEXT NOT NULL,
         level  VARCHAR(255) NOT NULL
      );

   CREATE TABLE IF NOT EXISTS sets 
      (
         id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
         owner_id uuid NOT NULL, 
         set_name VARCHAR(255) NOT NULL, 
         is_public BOOL NOT NULL DEFAULT true, 
         created_date timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'), 
         last_login timestamp without time zone DEFAULT (current_timestamp AT TIME ZONE 'UTC'), 
         last_modified_date timestamp without time zone 
      );

   CREATE TABLE IF NOT EXISTS cards 
      (
         id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
         owner_id uuid NOT NULL,
         set_id uuid NOT NULL,
         created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
         FOREIGN KEY (set_id) REFERENCES sets (id) ON DELETE CASCADE
      );

EOSQL
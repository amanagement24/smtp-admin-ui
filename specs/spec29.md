I have two files @docs/compare/new.sql and docs/compare/old.sql

- I want you to look to establish the differences between new and old for the tables that are only mentioned in new, not all tables, as old has more tables
- create ddl script that will patch the old database for the tables in new.sql 
- the ddl script will be placed under docs/compare/diff.sql and will consist in alter table etc

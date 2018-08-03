---
layout: post
author: derek
image: assets/images/mysql.png
featured: false
hidden: false
title: Simple MySql Command Line Cheatsheet
date: '2018-05-14 16:00:00'
comments: true
external-url: null
categories: database
keywords: mysql, database, command line, cheatsheet
excerpt: You are web developer, mobile developer or devops developer. But you should know simply MySQL command line cheat sheet below this article.
---

## Commands

Access monitor:

```bash
mysql -u [username] -p; #(will prompt for password)
```

Show all databases:

```bash
show databases;
```

Access database:

```bash
mysql -u [username] -p [database] #(will prompt for password)
```

Create new database:

```bash
create database [database];
```

Select database:

```bash
use [database];
```

Determine what database is in use:

```bash
select database();
```

Show all tables:

```bash
show tables;
```

Show table structure:

```bash
describe [table];
```

List all indexes on a table:

```bash
show index from [table];
```

Create new table with columns:

```bash
CREATE TABLE [table] ([column] VARCHAR(120), [another-column] DATETIME);
```

Adding a column:

```bash
ALTER TABLE [table] ADD COLUMN [column] VARCHAR(120);
```

Adding a column with an unique, auto-incrementing ID:

```bash
ALTER TABLE [table] ADD COLUMN [column] int NOT NULL AUTO_INCREMENT PRIMARY KEY;
```

Inserting a record:

```bash
INSERT INTO [table] ([column], [column]) VALUES ('[value]', [value]');
```

MySQL function for datetime input:

```bash
NOW()
```

Selecting records:

```bash
SELECT * FROM [table];
```

Explain records:

```bash
EXPLAIN SELECT * FROM [table];
```

Selecting parts of records:

```bash
SELECT [column], [another-column] FROM [table];
```

Counting records:

```bash
SELECT COUNT([column]) FROM [table];
```

Counting and selecting grouped records:

```bash
SELECT *, (SELECT COUNT([column]) FROM [table]) AS count FROM [table] GROUP BY [column];
```

Selecting specific records:

```bash
SELECT * FROM [table] WHERE [column] = [value];` (Selectors: `<`, `>`, `!=`; combine multiple selectors with `AND`, `OR`)
```

Select records containing

```bash
[value]`: `SELECT * FROM [table] WHERE [column] LIKE '%[value]%';
```

Select records starting with

```bash
[value]`: `SELECT * FROM [table] WHERE [column] LIKE '[value]%';
```

Select records starting with `val` and ending with `ue`:

```bash
SELECT * FROM [table] WHERE [column] LIKE '[val_ue]';
```

Select a range:

```bash
SELECT * FROM [table] WHERE [column] BETWEEN [value1] and [value2];
```

Select with custom order and only limit:

```bash
SELECT * FROM [table] WHERE [column] ORDER BY [column] ASC LIMIT [value];` (Order: `DESC`, `ASC`)
```

Updating records:

```bash
UPDATE [table] SET [column] = '[updated-value]' WHERE [column] = [value];
```

Deleting records:

```bash
DELETE FROM [table] WHERE [column] = [value];
```

Delete *all records* from a table (without dropping the table itself):

```bash
DELETE FROM [table];
```

(This also resets the incrementing counter for auto generated columns like an id column.)

Delete all records in a table:

```bash
truncate table [table];
```

Removing table columns:

```bash
ALTER TABLE [table] DROP COLUMN [column];
```

Deleting tables:

```bash
DROP TABLE [table];
```

Deleting databases:

```bash
DROP DATABASE [database];
```

Custom column output names:

```bash
SELECT [column] AS [custom-column] FROM [table];
```

Export a database dump (more info [here](http://stackoverflow.com/a/21091197/1815847)):

```bash
mysqldump -u [username] -p [database] > db_backup.sql
```

Use
```bash
--lock-tables=false
```
 option for locked tables (more info [here](http://stackoverflow.com/a/104628/1815847)).

Import a database dump (more info [here](http://stackoverflow.com/a/21091197/1815847)):

```bash
mysql -u [username] -p -h localhost [database] < db_backup.sql
```

Logout:

```bash
exit;
```
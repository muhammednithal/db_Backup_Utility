# db_Backup_Utility

`db_Backup_Utility` is a powerful and flexible CLI tool built in Go that enables seamless **backup** and **restoration** of MySQL and PostgreSQL databases. It supports multiple interaction modes including saved configurations, direct flag inputs, and interactive prompts.

## Features

- Supports MySQL and PostgreSQL
- Backup and restore using CLI or interactive mode
- Save reusable configurations for future use
- Structured logging of operations (success/failure)
- Cron-based scheduling support
- Discord Notification for backup & restore status
- Advanced backup options (e.g., incremental backups, compression)

---

## Installation

```bash
git clone https://github.com/muhammednithal/db_Backup_Utility.git
cd db_Backup_Utility
go mod tidy
go build -o db_Backup_Utility

```

## Usage

### Get help

```bash
    ./db_Backup_Utility help
```

You can perform database backups and restores in 3 ways:

1. Using command flags

2. Using saved configurations

3. Using interactive prompts

### Backup

#### Using flg

```bash
./db_Backup_Utility backup --type=mysql --host=localhost --port=3306 --user=root --password=secret --database=mydb --output=backup.sql
```

#### Using saved config:

```bash
./db_Backup_Utility backup --savedconfig "prod_db"
```

#### Interactive mode:

```bash
./db_Backup_Utility backup
```

simillar for restore also

### Configuration

#### Save a config interactively:

```bash
./db_Backup_Utility config
```

#### Delete a saved config:

```bash
./db_Backup_Utility config delete --name=prod_db
```

## Logging

Every backup or restore operation is logged with:

Action (backup or restore)

DB type, host, port, and database name

File path used

Status (success or failure)

Timestamp and optional error message

Logs are stored locally (logs/operations.log)

## Feature Roadmap

• Cloud storage integration (AWS S3, GCS, Azure)

• Add support for additional database systems (e.g., SQLite, MongoDB)

• Implement encryption for backup files

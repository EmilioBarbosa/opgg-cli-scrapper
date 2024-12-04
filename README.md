# OPGG Scrapper

CLI tool to scrape data about League of Legends champions from [op.gg](https://www.op.gg/).

## Instalation

Follow the steps bellow to install and run the cli.

### 1. Clone the repository

### 2. Install the dependencies

```
go mod tidy
```

### 3. Run

```
go run .
```

### 4. Parameters

Default value is nidalee, to search for specific champion you can use the `champion` flag.

```
go run . --champion=leblanc
```

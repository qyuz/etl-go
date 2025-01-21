# ETL Go

General purpose Extract Transform Load tool for

- loading series watchlist from TMDB into Notion database


## Loading series watchlist from TMDB into Notion database
### Setup
- Get [TMDB API key](https://www.themoviedb.org/settings/api)
- Get [Notion API key](https://www.notion.so/profile/integrations)
- Create Notion database
- Database fields should be named as follows:
  - Name
  - Movie ID
- Give integration access to the Notion database
  - open Database page
  - click `...`
  - open Connections
  - find your integration in `Search for connections`
- Copy `.env.prod.example` to `.env.prod` and fill in the required values

### Usage
```bash
go run main.go
```


### Testing
#### Unit tests
```bash
go test ./...
```

#### Loading series watchlist from TMDB into test Notion database
Copy `.env.test.example` to `.env.test` and fill in the required values

```bash
go run test.go
```

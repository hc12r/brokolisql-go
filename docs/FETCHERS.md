# Data Fetchers

This document describes the data fetching functionality in brokolisql-go, which allows retrieving data from various remote sources.

## Overview

The fetcher system is designed with a loosely coupled architecture to support multiple data sources. Currently, it supports:

- REST API endpoints (JSON data)

Future implementations will include:

- Database connections
- Other remote data sources

## Fetcher Interface

All fetchers implement the common `Fetcher` interface:

```go
type Fetcher interface {
    Fetch(source string, options map[string]interface{}) (*loaders.DataSet, error)
}
```

This interface allows for a consistent way to retrieve data regardless of the source.

## Available Fetchers

### REST Fetcher

The REST fetcher allows retrieving data from REST API endpoints that return JSON data.

#### Usage

```go
// Get a REST fetcher
fetcher, err := fetchers.GetFetcher("rest")
if err != nil {
    // Handle error
}

// Define options
options := map[string]interface{}{
    "method": "GET",
    "headers": map[string]string{
        "Accept": "application/json",
        "Authorization": "Bearer YOUR_TOKEN",
    },
    "timeout": 30 * time.Second,
}

// Fetch data
dataset, err := fetcher.Fetch("https://api.example.com/data", options)
if err != nil {
    // Handle error
}

// Use the dataset
for _, row := range dataset.Rows {
    // Process each row
    fmt.Println(row["id"], row["name"])
}
```

#### Supported Options

The REST fetcher supports the following options:

- `method`: HTTP method (default: "GET")
- `headers`: map[string]string of HTTP headers
- `body`: string or []byte request body
- `timeout`: time.Duration for request timeout (default: 30s)

## Integration with Existing Loaders

The fetchers return data in the same `DataSet` format used by the loaders, making it easy to integrate with the existing functionality. You can:

1. Fetch data from a remote source
2. Process it directly
3. Or save it to a file and use the existing loaders

Example of saving fetched data to a file:

```go
// Fetch data
dataset, err := fetcher.Fetch("https://api.example.com/data", options)
if err != nil {
    // Handle error
}

// Save to a file
// (implementation depends on your needs)
saveToFile(dataset, "data.json")

// Later, you can load it using the existing loaders
loader, err := loaders.GetLoader("data.json")
if err != nil {
    // Handle error
}
dataset, err = loader.Load("data.json")
```

## Extending with New Fetchers

To add a new fetcher:

1. Create a new file in the `pkg/fetchers` directory
2. Implement the `Fetcher` interface
3. Update the `GetFetcher` function in `fetcher.go` to return your new fetcher

Example:

```go
// In a new file like database_fetcher.go
type DatabaseFetcher struct {
    // Your implementation details
}

func (f *DatabaseFetcher) Fetch(source string, options map[string]interface{}) (*loaders.DataSet, error) {
    // Your implementation
}

// Then in fetcher.go, update GetFetcher:
func GetFetcher(sourceType string) (Fetcher, error) {
    switch sourceType {
    case "rest":
        return &RESTFetcher{}, nil
    case "database":
        return &DatabaseFetcher{}, nil
    default:
        return nil, ErrUnsupportedSourceType
    }
}
```

## Error Handling

The fetcher system defines several error types to help with error handling:

- `ErrUnsupportedSourceType`: When an unsupported fetcher type is requested
- `ErrInvalidURL`: When an invalid URL is provided
- `ErrHTTPRequestFailed`: When an HTTP request fails
- `ErrEmptyResponse`: When an empty response is received

You can check for these errors using Go's error wrapping:

```go
if errors.Is(err, fetchers.ErrHTTPRequestFailed) {
    // Handle HTTP request failure
}
```
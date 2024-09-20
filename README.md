# Mdparser

Markdown-like parser written in go. Does not fully support all markdown standard.

## How to use

```go
md := `
    # Markdown here
`

parsedMd := mdparser.Parse(md)
```

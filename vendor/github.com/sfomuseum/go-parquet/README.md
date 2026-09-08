# go-parquet

Opinionated Go package for working with Parquet files.

## Motivation

This is not a general purpose Parquet package. Currently it has exactly two functions:

```
Iterate[T any](ctx context.Context, uris ...string) iter.Seq2[*T, error]
```

Which is designed to iterate over a series of URIs which may be local files on disk or remote files served over HTTP(S).

And

```
NewWriter[T any](ctx context.Context, uri string) (*ParquetWriter[T], error)
```

Which mimics the `parquet-go/parquet-go.GenericWriter` but buffers rows and writes them in batches.

I have copy-pasted these functions in to different packages enough times that it seemed worth putting them in a common `import`-able package of their own.
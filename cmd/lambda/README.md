# bfconv-lambda

bfconv-lambda is an AWS Lambda implementation of `go-bfconv` that converts Hatena Bookmark RSS feeds to JSON Feed and serves them over HTTPS.

## Overview

- Requests ending with `.json` or `/json` are mapped to the corresponding `.rss` or `/rss` path on `b.hatena.ne.jp`, fetched, and returned as JSON Feed (e.g., `/entrylist.json` → `/entrylist.rss`).
- This service currently supports only selected Hatena Bookmark RSS endpoints. See the below section for the exact list.
- The root path (/) returns an overview of the service.

### Supported RSS feeds

- `https://b.hatena.ne.jp/hotentry.rss`
- `https://b.hatena.ne.jp/hotentry/{category}.rss`
- `https://b.hatena.ne.jp/entrylist.rss`
- `https://b.hatena.ne.jp/entrylist/{category}.rss`
- `https://b.hatena.ne.jp/{user}/rss`
- `https://b.hatena.ne.jp/{user}/bookmark.rss`

### Upcoming / Unsupported feeds

- `https://b.hatena.ne.jp/q/{tag}?mode=rss`
- `https://b.hatena.ne.jp/keyword/{keyword}?mode=rss`
- `https://b.hatena.ne.jp/entrylist?sort={kind}&url={url}&mode=rss`
- `https://b.hatena.ne.jp/entry/rss/http://example.com/`

## Usage

```bash
$ git clone https://github.com/sigsignv/go-bfconv.git
$ cd go-bfconv/cmd/lambda
$ make package
$ # Upload the generated lambda-handler.zip artifact to your Lambda function
```

## Author

- Sigsign <<sig@signote.cc>>

## License

Apache-2.0

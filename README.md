# URL requester with thread pool

A tool for making parallel requests to more than one url.
The default parallel number is 10, but the user can increase/decrease this number.

## Compile

```bash
    go build myhttp.go
```

## Execute

1. Without specify parallel request.

```bash
    ./myhttp google.com stackoverflow www.facebook.com www.twitter.com
```

2With specify parallel request.

```bash
    ./myhttp -parallel 4 google.com http://stackoverflow.com www.facebook.com www.twitter.com
```

## Test

```bash
    go test -v
```

Tool help.

```bash
    myhttp -h 
```

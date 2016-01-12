# Stream

# Install

```
$ go get -u github.com/bluele/stream/cmd/stream
```

# How to use

## `cat` command

```sh
$ echo "ok" | stream cat s3://mybucket/sample.txt

$ stream cat s3://mybucket/sample.txt
ok
```

## `ls` command

```sh
$ stream ls s3://mybucket/
sample.txt
```

## `cp` command

```sh
$ stream cp ./text.txt s3://mybucket/
```

github-cli
==========

github command line tool golang

## Install ##


## Usage ##

### Init Config

```
$ github-cli init --host https://github.com/ --api-path api/v3/ --token aaaaaaaaaaa
```

- `--host`: github host url
- `--api-path`: github api version path
- `--token`: your access token

### Issue List

```
$ github-cli list
```

- `--state`: state filter (`opened` or `closed`)

### Create Issue

```
$ github-cli add -t title -d hoge -l aaa,bbbb,hoge,tag
```

- `-t`: issue title
- `-d`: issue detail
- `-l`: issue labels (array of a comma delimited string)

### Issue Detail

```
$ github-cli issue --id 28
```

- `--issue-id, --id`: issue localID

## LICENSE

[MIT](https://github.com/kyokomi/github-cli/blob/master/LICENSE)

## Author

[kyokomi](https://github.com/kyokomi)


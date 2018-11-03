# Gitlab Repostiory Downloader

### Download & Run

see [Release](https://github.com/AmirSoleimani/GitlabRepoDownloader/releases)

### Build & Run
##### Step 1. Git Clone
```
git clone https://github.com/AmirSoleimani/GitlabRepoDownlaoder
```

##### Step 2. Get Dependencies
```
go get ./...
```
##### Step 3. Build
```
go build .
```
##### Step 4. Run
```
./glrd -token BlaBLaBlaBla-SB1bLA2 -host https://gitlab.example.com -datadir ./data
```


### CLi


| -        | description  |
|------------------|:--------------:|
| `host`         | gitlab host url (https://... , http://....)            |
| `token`          | gitlab user private token            |
| `datadir`        | data directory path |



# Thanks.
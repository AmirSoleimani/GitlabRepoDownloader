# Gitlab Repostiory Downloader
----
# Build & Run
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
| `host`         | gitlab host url            |
| `token`          | gitlab user private token            |
| `datadir`        | data directory path |



# Thanks.
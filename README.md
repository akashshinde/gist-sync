# gist-sync

Sync gist file to local disk

### Usage

#### Config
Create config file at `$HOME/.gist-sync/config`.
Update config file with Gist Id and local file path.
```
GistId: 4ac4e3beae76e671e78f199997d7c937
SyncFilePath: /Users/akash/.kube/config_backup
```

#### Run gist-sync
Download latest release https://github.com/akashshinde/gist-sync/releases

run `gist-sync`

It would download the Gist content(`GistId`) and store it in the `SyncFilePath`

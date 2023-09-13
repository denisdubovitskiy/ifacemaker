## Ifacemaker

### Usage

```shell
ifacemaker \
  --source-pkg github.com/mattermost/mattermost-server/v5@v5.39.3 \
  --module-path model \
  --result-pkg client \
  --struct-name Client4 \
  --interface-name Client4 \
  --output mattermost/client.go
```

### Parameters

* `--source-pkg` - A source package in which the desired struct is located.
* `--module-path` - A full path to the struct package where desired struct resides.
  Should start from the source package's root.
* `--result-pkg` - A name for the resulting package.
* `--struct-name` - A name of the struct from which an interface should be generated.
* `--interface-name` - A name for resulting interface.
* `--output` - A filename in which a result interface is going to be stored.
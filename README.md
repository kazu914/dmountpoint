 This software is released under the MIT License, see [LICENSE](./LICENSE)
 
 # What's `dps`
 
 `dmountpoint` shows the lists of the processes running in your all containers.

# Screenshot
![screenshot](https://github.com/kazu914/dmountpoint/blob/images/screenshot.png)

# How to install

```
wget -O - https://github.com/kazu914/dmountpoint/releases/download/v1.0.0/dmountpoint_1.0.0_linux_amd64.tar.gz | tar zxvf -
```

# How to use

```
./dmountpoint
```

When you get the blow error,

```
panic: Error response from daemon: client version <x.xx> is too new. Maximum supported API version is <y.yy>
```

you have to use `--apiversion` flag to avoid it.

```
./dmountpoint --apiversion <y.yy>
```


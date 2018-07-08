### The signal server of [hlsjs-p2p-engine](https://github.com/cdnbye/hlsjs-p2p-engine)

### build
- install dependency

- go build main.go

### test
```
var hlsjsConfig = {
    p2pConfig: {
        wsSignalerAddr: 'ws://localhost:8081/ws',
        // Other p2pConfig options if applicable
    }
};
// Hls constructor is overriden by included bundle
var hls = new Hls(hlsjsConfig);
// Use `hls` just like the usual hls.js ...
```

### go语言版的 [hlsjs-p2p-engine](https://github.com/cdnbye/hlsjs-p2p-engine)信令服务器


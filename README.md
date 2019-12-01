### The signal server of [hlsjs-p2p-engine](https://github.com/cdnbye/hlsjs-p2p-engine)

#### install dependency
```bash
git clone https://github.com/cdnbye/gosignaler.git

go get github.com/gorilla/websocket
```

#### compile
```bash
cd gosignaler && make
```

#### deploy
Upload binary file to server, create `cert` directory with `crt.pem` and `crt.key`, then start service:
```bash
./admin.sh start
```

### test
```
import Hls from 'cdnbye';
var hlsjsConfig = {
    p2pConfig: {
        wsSignalerAddr: 'ws://YOUR_SIGNAL',
        // Other p2pConfig options provided by hlsjs-p2p-engine
    }
};
// Hls constructor is overriden by included bundle
var hls = new Hls(hlsjsConfig);
// Use `hls` just like the usual hls.js ...
```

### go语言版的 CDNBye 信令服务器，可用于Web、安卓、iOS SDK等所有CDNBye产品
#### 安装依赖
```bash
git clone https://github.com/cdnbye/gosignaler.git

go get github.com/gorilla/websocket
```

#### 编译二进制文件
```bash
cd gosignaler && make
```

#### 部署
将编译生成的二进制文件上传至服务器，并在同级目录创建`cert`文件夹，将证书和秘钥文件分别改名为`crt.pem`和`crt.key`放入cert，之后启动服务：
```bash
./admin.sh start
```

### 测试
```
import Hls from 'cdnbye';
var hlsjsConfig = {
    p2pConfig: {
        wsSignalerAddr: 'ws://YOUR_SIGNAL',
        // Other p2pConfig options provided by hlsjs-p2p-engine
    }
};
// Hls constructor is overriden by included bundle
var hls = new Hls(hlsjsConfig);
// Use `hls` just like the usual hls.js ...
```




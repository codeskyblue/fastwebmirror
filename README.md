# fastwebmirror
NOT Finished!! mirror web resources when request. Good at cache some download site.


## Install
```bash
git clone https://github.com/codeskyblue/fastwebmirror
cd fastwebmirror
go build
```

## Usage
```bash
./fastwebmirror --port=3000 -w /:dl.google.com/go/
```

When request <http://localhost:3000/go.1.11.2.src.tar.gz> file will be download from <https://dl.google.com/go/go.1.11.2.src.tar.gz>, and cached in local.


## LICENSE
[MIT](LICENSE)

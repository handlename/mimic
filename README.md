# mimic

simple http mock server.

## Example

`config.sample.json` is [here](https://github.com/handlename/mimic/config.sample.json).

```bash
$ mimic --config=config.sample.json
2014/08/29 15:15:38 registered: GET /
2014/08/29 15:15:38 registered: POST /api/message
2014/08/29 15:15:38 registered: GET /api/message
2014/08/29 15:15:38 mimic server started at localhost:3390
```

```bash
$ curl -X POST http://localhost:3390/api/message
{"status":"ok"}%                                                                                                                                                                                                               -- 2014/08/29 15:16:30 --

$ curl -X GET http://localhost:3390/api/message
{"message":"Hi, I'm mimic!!"}%                                                                                                                                                                                                 -- 2014/08/29 15:16:35 --

$ curl -X GET http://localhost:3390/hoge
404 page not found
```

## Installation

Download binary for your environment from [release page](https://github.com/handlename/mimic/releases)
and place it in `$PATH` directory.

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[handlename](https://github.com/handlename)

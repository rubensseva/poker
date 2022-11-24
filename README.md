# Poker analyzer
## How to run

``` sh
$ docker build -t poker-analyzer .
$ docker run -p 3333:3333 poker-analyzer
```

Then you can make a request to the server with a command like this:
``` sh
$ curl localhost:3333/analyze-hand/3r-5r-6s-3k-3s
```


# Kooper Zap ⚡️

Simple package to use your zap logger for kooper.

It is just simple interface matching.

## How to

```go

import kooperzap "github.com/thomas-tacquet/kooper-zap"
...

logger := kooperzap.NewKooperLogger(myZapInstance)

kooperCfg := controller.Config{
    Name:   "example",
    Logger: logger,
}
```

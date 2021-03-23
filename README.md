# log

A wrapper on top of Golang log package which includes levelling and rotated file output. Complete documentation at [Godoc](https://godoc.org/github.com/arknable/log).

## Usage

Following example uses default logger implementation.

```
import "github.com/arknable/log"

... etc

l := log.New()

l.Debug("This is a debug message")
l.Info("This is", "an info message")
l.Warning("This is a warning message")
l.Error("This is", "an error message")
```

There is also package-level logger.
```
import "github.com/arknable/log"

... etc

log.Debug("This is a debug message")
log.Info("This is", "an info message")
log.Warning("This is a warning message")
log.Error("This is", "an error message")
```

By default, log messages sent to stdout. To use custom output use `SetOutput()` function of the logger.
```
import "github.com/arknable/log"

... etc

file, err := os.Open(......)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

l := log.New()
l.SetOutput(io.MultiWriter(os.Stdout, file))
```

or you can set custom output to package-level logger.
```
import "github.com/arknable/log"

... etc

file, err := os.Open(......)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

log.SetOutput(io.MultiWriter(os.Stdout, file))
```

If you need file output that has name rolled every day, there is function `DailyFileOutput`.
```
import "github.com/arknable/log"

... etc

file, err := DailyFileOutput("test_daily_file", "./")
if err != nil {
    t.Fatal(err)
}
defer file.Close()

logger := New()
logger.SetOutput(io.MultiWriter(os.Stdout, file))
```

## Sample Output
```
2021/03/23 05:35:19   DEBUG This is a debug message
2021/03/23 05:35:19    INFO This is an info message
2021/03/23 05:35:19 WARNING This is a warning message
2021/03/23 05:35:19   ERROR This is an error message
2021/03/23 05:35:19   DEBUG This is a debugf message
2021/03/23 05:35:19    INFO This is an infof message
2021/03/23 05:35:19 WARNING This is a warningf message
2021/03/23 05:35:19   ERROR This is an errorf message
```
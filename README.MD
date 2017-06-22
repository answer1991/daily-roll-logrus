### How To Use

````go

 logger := drl.GetLogger("test-module")

 logger.Info("hello world")

````

### Where Log File Exist

By default, log files is in ``$HOME/logs/``.

Change log files store path, set ``DAILY_ROLL_LOGRUS_LOG_PATH`` environment.
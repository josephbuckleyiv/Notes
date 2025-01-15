# Dealing with Time in C#

## HTTP/S
Prefer using UnixTime, as this long/double (seconds vs. millisecond accuracy) is unopinionated and exact. Assume any communication that requires Date before 1970 will be passed some other way.
```
public static double Unix()
{
  var timeSpan = (DateTime.UtcNow - new DateTime(1970, 1, 1, 0, 0, 0, 0));
  return (double)timeSpan.TotalMilliseconds;
}
```
UTC is the co-ordinated standard for regulating clock times throughout the world. UtcNow returns the least opinionated clocktime 'by definition.' Assume endpoints know what to do with time, e.g a browser may convert it for the user's pleasure. That way a Huawei in Timbuktu and a Canadian supercomputer both are served the same digits, and they choose what they want to do with it. // 
As a result, we slot it into the query parameters, with maybe something like 'zis:
```
var query = $"https://localhost:69?time={eunuchsTime}"
```

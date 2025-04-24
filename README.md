## Dead Simple CPU Usage Calculator
CPU Usage is calculated with the help of the `/proc` filesystem in linux. 

``` bash
mpstat 1 | awk '/all/ {print 100 - $NF}'
```
 gives the same output. It gives stream of output. Takes CPU usage for 1 second.

``` bash
mpstat 1 1 | awk '/all/ {print 100 - $NF}'
```

If you want only 1 stat. 

``` bash
mpstat | awk '/all/ {print 100 - $NF}'
```
From System start to current time

`mpstat` is part of [sysstat](https://github.com/sysstat/sysstat), and you might want to use that instead of this. 

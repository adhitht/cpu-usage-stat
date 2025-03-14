## Dead Simple CPU Usage Calculator
CPU Usage is calculated with the help of the `/proc` filesystem in linux. 

`mpstat 1 | awk '/all/ {print 100 - $NF}'` gives the same output. 

`mpstat` is part of `sysstat`, and you might want to use that instead of this. 
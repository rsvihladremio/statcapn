# statcapn

**Alpha Status**

metrics collection but easy. Tired of working with environments with no metrics collection setup, limited install rights, in hostile environments and especially at the frequency, duration, and level of aggregation you want? Now you can use StatCapn the easy to install metrics capture utility written in go. Originally conceived to capture usage stats out of a K8s pod with no monitoring or metrics collection setup

## Install

Download a binary here https://github.com/rsvihladremio/statcapn/releases unzip it and just run it (works on Windows, Mac and Linux)

## How to use

by default it outputs some metrics to standard out

```
./bin/statcapn
                Timestamp	    usr %%	    sys %%	 iowait %%	  other %%	    idl %%	     Queue	Latency (ms)	Read (MB/s)	Write (MB/s)	Free Mem (GB)
2023-07-03T16:39:01+02:00	     5.01%	     6.26%	     0.00%	     0.00%	    88.74%	      0.00	      0.00	      0.00	      0.00	      0.96
2023-07-03T16:39:02+02:00	     4.62%	     4.87%	     0.00%	     0.00%	    90.51%	      0.00	      0.00	      0.00	      0.00	      1.05
2023-07-03T16:39:03+02:00	     5.24%	     5.62%	     0.00%	     0.00%	    89.14%	      0.00	      0.00	      0.00	      0.00	      1.03
2023-07-03T16:39:04+02:00	     6.63%	     6.51%	     0.00%	     0.00%	    86.86%	      0.00	      0.00	      0.00	      0.00	      1.05
2023-07-03T16:39:05+02:00	     8.91%	    10.04%	     0.00%	     0.00%	    81.05%	      0.00	      0.00	      0.00	      0.00	      1.02
2023-07-03T16:39:06+02:00	     5.27%	     5.90%	     0.00%	     0.00%	    88.83%	      0.00	      0.00	      0.00	      0.00	      1.02
2023-07-03T16:39:07+02:00	     8.49%	     9.24%	     0.00%	     0.00%	    82.27%	      0.00	      0.00	      0.00	      0.00	      0.99
2023-07-03T16:39:08+02:00	     7.64%	     8.52%	     0.00%	     0.00%	    83.83%	      0.00	      0.00	      0.00	      0.00	      0.95
2023-07-03T16:39:09+02:00	     8.10%	     4.99%	     0.00%	     0.00%	    86.91%	      0.00	      0.00	      0.00	      0.00	      0.95
2023-07-03T16:39:10+02:00	     8.55%	     4.40%	     0.00%	     0.00%	    87.04%	      0.00	      0.00	      0.00	      0.00	      0.96
2023-07-03T16:39:11+02:00	     6.60%	     4.86%	     0.00%	     0.00%	    88.54%	      0.00	      0.00	      0.00	      0.00	      0.98
2023-07-03T16:39:12+02:00	     4.64%	     5.64%	     0.00%	     0.00%	    89.72%	      0.00	      0.00	      0.00	      0.00	      1.03
```
control+c when you have seen enough

### Json output

A simple enough command just give it a file that ends in .json. In the following example I have given 10 seconds for a duration (however one can also just control+c when done, the content is not buffered so the file updates as it collect measurements.

```./bin/statcapn -d 10 ~/Downloads/metrics.json```

```json
{"collectionTimestamp":"2023-07-03T16:41:32.556974+02:00","userCPUPercent":5.53,"systmeCPUPercent":8.29,"idleCPUPercent":86.18,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":973,"cachedRAMMB":0}
{"collectionTimestamp":"2023-07-03T16:41:33.567511+02:00","userCPUPercent":6.72,"systmeCPUPercent":7.96,"idleCPUPercent":85.32,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":946,"cachedRAMMB":0}
{"collectionTimestamp":"2023-07-03T16:41:34.571189+02:00","userCPUPercent":8.89,"systmeCPUPercent":8.14,"idleCPUPercent":82.98,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":963,"cachedRAMMB":0}
{"collectionTimestamp":"2023-07-03T16:41:35.573744+02:00","userCPUPercent":13,"systmeCPUPercent":19.12,"idleCPUPercent":67.88,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":875,"cachedRAMMB":0}
{"collectionTimestamp":"2023-07-03T16:41:36.575854+02:00","userCPUPercent":9.49,"systmeCPUPercent":8.61,"idleCPUPercent":81.9,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":936,"cachedRAMMB":0}
{"collectionTimestamp":"2023-07-03T16:41:37.578162+02:00","userCPUPercent":6.88,"systmeCPUPercent":8.76,"idleCPUPercent":84.36,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":986,"cachedRAMMB":0}
{"collectionTimestamp":"2023-07-03T16:41:38.580396+02:00","userCPUPercent":5.24,"systmeCPUPercent":4.49,"idleCPUPercent":90.26,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":1012,"cachedRAMMB":0}
{"collectionTimestamp":"2023-07-03T16:41:39.583214+02:00","userCPUPercent":6.63,"systmeCPUPercent":5.51,"idleCPUPercent":87.86,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":1029,"cachedRAMMB":0}
{"collectionTimestamp":"2023-07-03T16:41:40.584502+02:00","userCPUPercent":8.52,"systmeCPUPercent":5.14,"idleCPUPercent":86.34,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":1038,"cachedRAMMB":0}
{"collectionTimestamp":"2023-07-03T16:41:41.585611+02:00","userCPUPercent":12.42,"systmeCPUPercent":4.89,"idleCPUPercent":82.69,"niceCPUPercent":0,"ioWaitCPUPercent":0,"irqCPUPercent":0,"softIRQCPUPercent":0,"stealCPUPercent":0,"guestCPUPercent":0,"guestCPUNicePercent":0,"queueDepth":0,"diskLatency":0,"readBytes":0,"writeBytes":0,"freeRAMMB":1069,"cachedRAMMB":0}
```

### Full help

```
./bin/statcapn -h
statcapn v0.1.1-ce0e9c4

standard usage:
	statcapn -i <interval> -d <duration_seconds> metrics.txt

For json output:
	statcapn -i <interval> -d <duration_seconds> metrics.json

flags:

  -d int
    	number of seconds for duration of all collection (default 9223372036854775807)
  -i int
    	number of seconds between execution of collection (default 1)
```

## License

Apache License, Version 2.0

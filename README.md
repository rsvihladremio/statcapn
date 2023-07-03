# statcapn

**Alpha Status**

metrics collection but easy. Tired of working with environments with no metrics collection setup, limited install rights, in hostile environments and especially at the frequency, duration, and level of aggregation you want? Now you can use StatCapn the easy to install metrics capture utility written in go. Originally conceived to capture usage stats out of a K8s pod with no monitoring or metrics collection setup I decided to use it


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

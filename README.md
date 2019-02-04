# Задание:
> Есть односвязный список размера n (1, 2, 3, 4, 5, 6, 7, 8).
Необходимо перепаковать его так, что после первого элемента шел последний,
после второго - предпоследний и т.д. (1, 8, 2, 7, 3, 6, 4, 5).

# Сравнение времени выполнения и выделяемой памяти для 4 сортировок:
```
goos: windows
goarch: amd64
pkg: github.com/sidyakina/SortingLOS
BenchmarkLOS_Sort1_10-4      	10000000	       201 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort2_10-4      	10000000	       199 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort3_10-4      	 1000000	      1132 ns/op	      96 B/op	       6 allocs/op
BenchmarkLOS_Sort4_10-4      	10000000	       208 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort1_100-4     	  100000	     17891 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort2_100-4     	  200000	      9355 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort3_100-4     	  200000	      9640 ns/op	     816 B/op	      51 allocs/op
BenchmarkLOS_Sort4_100-4     	  200000	      9200 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort1_1000-4    	    1000	   1857106 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort2_1000-4    	    2000	    981056 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort3_1000-4    	   10000	    144508 ns/op	    8016 B/op	     501 allocs/op
BenchmarkLOS_Sort4_1000-4    	   10000	    156908 ns/op	    8016 B/op	     501 allocs/op
BenchmarkLOS_Sort1_10000-4   	      10	 245114020 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort2_10000-4   	      20	 161809255 ns/op	       0 B/op	       0 allocs/op
BenchmarkLOS_Sort3_10000-4   	    2000	   1198068 ns/op	   80016 B/op	    5001 allocs/op
BenchmarkLOS_Sort4_10000-4   	    2000	   1173067 ns/op	   80016 B/op	    5001 allocs/op


```


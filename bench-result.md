# Benchmark result

## Local machine

### Machine #1

- MacBook Air (M1, 2020)
- 8GB RAM

### commit hash `06513349873fe2a37fdc669f9ac206253b3bcbfc`

`./kash -mode client -command benchmark -n 1000 -parallel 50 -key-size 5`

```
-> Benchmark result
# Worker 0
    sum: 1.009950591s
    avg: 1.00995ms
    p50: 829.75µs
    p90: 2.0675ms
    p99: 4.005958ms
# Worker 1
    sum: 962.906068ms
    avg: 962.906µs
    p50: 759.375µs
    p90: 2.031333ms
    p99: 3.815416ms
# Worker 2
    sum: 1.003774223s
    avg: 1.003774ms
    p50: 837.167µs
    p90: 2.077167ms
    p99: 3.669208ms
# Worker 3
    sum: 1.014555651s
    avg: 1.014555ms
    p50: 818.583µs
    p90: 2.153208ms
    p99: 4.065458ms
# Worker 4
    sum: 1.001279727s
    avg: 1.001279ms
    p50: 787.75µs
    p90: 2.0835ms
    p99: 4.022666ms
# Worker 5
    sum: 1.013679585s
    avg: 1.013679ms
    p50: 846.834µs
    p90: 2.053167ms
    p99: 3.593667ms
# Worker 6
    sum: 1.009297129s
    avg: 1.009297ms
    p50: 827.125µs
    p90: 2.081959ms
    p99: 3.748542ms
# Worker 7
    sum: 974.805319ms
    avg: 974.805µs
    p50: 820.5µs
    p90: 1.983792ms
    p99: 3.599958ms
# Worker 8
    sum: 998.931739ms
    avg: 998.931µs
    p50: 765.583µs
    p90: 2.116333ms
    p99: 3.652167ms
# Worker 9
    sum: 994.45079ms
    avg: 994.45µs
    p50: 817.834µs
    p90: 2.039125ms
    p99: 3.861375ms
# Worker 10
    sum: 980.910448ms
    avg: 980.91µs
    p50: 800.417µs
    p90: 2.004375ms
    p99: 3.507875ms
# Worker 11
    sum: 995.328475ms
    avg: 995.328µs
    p50: 800.458µs
    p90: 2.059959ms
    p99: 3.389208ms
# Worker 12
    sum: 1.011414696s
    avg: 1.011414ms
    p50: 806.125µs
    p90: 2.067459ms
    p99: 4.044875ms
# Worker 13
    sum: 1.012539378s
    avg: 1.012539ms
    p50: 818.333µs
    p90: 2.127875ms
    p99: 3.640458ms
# Worker 14
    sum: 1.001220843s
    avg: 1.00122ms
    p50: 829.125µs
    p90: 2.018833ms
    p99: 3.463958ms
# Worker 15
    sum: 1.020405672s
    avg: 1.020405ms
    p50: 827.917µs
    p90: 2.151291ms
    p99: 4.169166ms
# Worker 16
    sum: 972.015164ms
    avg: 972.015µs
    p50: 772.708µs
    p90: 2.042083ms
    p99: 3.375959ms
# Worker 17
    sum: 1.009574679s
    avg: 1.009574ms
    p50: 796.833µs
    p90: 2.08725ms
    p99: 3.66875ms
# Worker 18
    sum: 1.0114183s
    avg: 1.011418ms
    p50: 821.5µs
    p90: 2.088583ms
    p99: 3.895083ms
# Worker 19
    sum: 1.008968893s
    avg: 1.008968ms
    p50: 815.959µs
    p90: 2.115875ms
    p99: 3.837375ms
# Worker 20
    sum: 1.011241192s
    avg: 1.011241ms
    p50: 822µs
    p90: 2.097458ms
    p99: 3.610042ms
# Worker 21
    sum: 992.059464ms
    avg: 992.059µs
    p50: 812.292µs
    p90: 1.958625ms
    p99: 3.65525ms
# Worker 22
    sum: 972.013628ms
    avg: 972.013µs
    p50: 823.25µs
    p90: 1.964625ms
    p99: 3.422208ms
# Worker 23
    sum: 1.014770372s
    avg: 1.01477ms
    p50: 786.833µs
    p90: 2.116833ms
    p99: 3.985792ms
# Worker 24
    sum: 1.003313743s
    avg: 1.003313ms
    p50: 833.458µs
    p90: 2.067625ms
    p99: 3.823167ms
# Worker 25
    sum: 1.00803901s
    avg: 1.008039ms
    p50: 829.75µs
    p90: 2.106042ms
    p99: 3.611833ms
# Worker 26
    sum: 1.017567786s
    avg: 1.017567ms
    p50: 827.875µs
    p90: 2.196834ms
    p99: 3.5715ms
# Worker 27
    sum: 973.657507ms
    avg: 973.657µs
    p50: 813.75µs
    p90: 1.931041ms
    p99: 3.587792ms
# Worker 28
    sum: 1.007376427s
    avg: 1.007376ms
    p50: 834.417µs
    p90: 2.132917ms
    p99: 3.846292ms
# Worker 29
    sum: 962.165749ms
    avg: 962.165µs
    p50: 800.125µs
    p90: 1.941125ms
    p99: 3.589834ms
# Worker 30
    sum: 972.506726ms
    avg: 972.506µs
    p50: 790.25µs
    p90: 2.030375ms
    p99: 3.788417ms
# Worker 31
    sum: 1.004170165s
    avg: 1.00417ms
    p50: 811.125µs
    p90: 2.109542ms
    p99: 3.754208ms
# Worker 32
    sum: 1.004343004s
    avg: 1.004343ms
    p50: 827.292µs
    p90: 2.083792ms
    p99: 3.796208ms
# Worker 33
    sum: 975.790552ms
    avg: 975.79µs
    p50: 805.042µs
    p90: 2.019209ms
    p99: 3.423542ms
# Worker 34
    sum: 1.018014663s
    avg: 1.018014ms
    p50: 810.709µs
    p90: 2.162833ms
    p99: 3.866417ms
# Worker 35
    sum: 1.005591791s
    avg: 1.005591ms
    p50: 780.958µs
    p90: 2.111625ms
    p99: 4.130084ms
# Worker 36
    sum: 1.01397928s
    avg: 1.013979ms
    p50: 785.375µs
    p90: 2.050583ms
    p99: 3.685375ms
# Worker 37
    sum: 997.433538ms
    avg: 997.433µs
    p50: 833.792µs
    p90: 1.972709ms
    p99: 3.884792ms
# Worker 38
    sum: 1.004703174s
    avg: 1.004703ms
    p50: 837.167µs
    p90: 2.070917ms
    p99: 3.904417ms
# Worker 39
    sum: 1.003982755s
    avg: 1.003982ms
    p50: 821.625µs
    p90: 2.049666ms
    p99: 3.3885ms
# Worker 40
    sum: 993.793095ms
    avg: 993.793µs
    p50: 815.583µs
    p90: 2.053333ms
    p99: 3.428416ms
# Worker 41
    sum: 1.016305541s
    avg: 1.016305ms
    p50: 847.542µs
    p90: 2.097084ms
    p99: 3.819042ms
# Worker 42
    sum: 1.003284583s
    avg: 1.003284ms
    p50: 803.917µs
    p90: 2.129875ms
    p99: 3.890375ms
# Worker 43
    sum: 1.013726989s
    avg: 1.013726ms
    p50: 825.75µs
    p90: 2.072542ms
    p99: 3.754708ms
# Worker 44
    sum: 1.018156279s
    avg: 1.018156ms
    p50: 826.625µs
    p90: 2.121375ms
    p99: 3.656917ms
# Worker 45
    sum: 1.002678815s
    avg: 1.002678ms
    p50: 821.458µs
    p90: 2.100833ms
    p99: 3.834209ms
# Worker 46
    sum: 1.000676544s
    avg: 1.000676ms
    p50: 841.209µs
    p90: 2.035166ms
    p99: 3.703625ms
# Worker 47
    sum: 1.018503681s
    avg: 1.018503ms
    p50: 796.25µs
    p90: 2.091916ms
    p99: 3.862041ms
# Worker 48
    sum: 986.978606ms
    avg: 986.978µs
    p50: 802.958µs
    p90: 2.082959ms
    p99: 3.629541ms
# Worker 49
    sum: 1.011583988s
    avg: 1.011583ms
    p50: 815.416µs
    p90: 2.106292ms
    p99: 3.340417ms
```

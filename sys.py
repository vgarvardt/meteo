#!/usr/bin/env python3
import subprocess

# result.stdout = b"temp=42.9'C\n"
result = subprocess.run(['vcgencmd', 'measure_temp'], stdout=subprocess.PIPE)
cpu_temp = float(result.stdout.decode('utf-8').split("=")[1].split("'")[0])
print("CPU temperature:", round(cpu_temp, 1))

# result.stdout = 0.01,0.03,0.00
result = subprocess.run("top -b | head -n 1 | awk '{print $10 $11 $12}'", shell=True, stdout=subprocess.PIPE)
la_1min, la_5min, la_15min = list(map(float, result.stdout.decode('utf-8').split(",")))
print("Load average:", round(la_1min, 2), round(la_5min, 2), round(la_15min, 2))

result = subprocess.run("free --kilo | grep Mem | awk '{print $2,$3,$4,$5,$6,$7}'", shell=True, stdout=subprocess.PIPE)
mem_total_kb, mem_used_kb, mem_free_kb, mem_shared_kb, mem_cache_kb, mem_available_kb = list(
    map(int, result.stdout.decode('utf-8').split(" ")))
print(
    "Memory:\n\ttotal: {0}kb\n\tused: {1}kb\n\tfree: {2}kb\n\tshared: {3}kb\n\tcache: {4}kb\n\tavailable: {5}kb".format(
        mem_total_kb, mem_used_kb, mem_free_kb, mem_shared_kb, mem_cache_kb, mem_available_kb))

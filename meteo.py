#!/usr/bin/env python3
import hashlib
import json
import os
import subprocess
from datetime import datetime, timezone

import paho.mqtt.publish as publish
from Adafruit_IO import Client, Feed, RequestError
from sense_hat import SenseHat

import sys

ADAFRUIT_IO_USERNAME = os.environ.get('AIO_USERNAME') or sys.exit('AIO_USERNAME env var is not defined\n')
ADAFRUIT_IO_KEY = os.environ.get('AIO_KEY') or sys.exit('AIO_KEY env var is not defined\n')

MQTT_HOST = os.environ.get('MQTT_HOST') or sys.exit('MQTT_HOST env var is not defined\n')
MQTT_PORT = int(os.environ.get('MQTT_PORT', 1883))
MQTT_CLIENT_ID = os.environ.get('MQTT_HOST', 'bedroom-sense-hat')

QOS_AT_MOST_ONCE = 0
QOS_AT_LEAST_ONCE = 1
QOS_EXACTLY_ONCE = 2

now = datetime.now(timezone.utc).astimezone().isoformat()
measurement_id = hashlib.sha256(now.encode('utf-8')).hexdigest()

# Create an instance of the REST client.
aio = Client(ADAFRUIT_IO_USERNAME, ADAFRUIT_IO_KEY)


def get_or_create_feed(key, name, description=''):
    try:
        return aio.feeds(key)
    except RequestError as e:
        print("Got an error when tried to get feed", name, e)
        feed = Feed(key=key, name=name, description=description)
        return aio.create_feed(feed)


sense = SenseHat()
sense.clear()

humidity = round(sense.get_humidity(), 1)
print("Humidity", humidity)
pressure = round(sense.get_pressure(), 1)
print("Pressure", pressure)
temp_humidity = round(sense.get_temperature_from_humidity(), 1)
temp_pressure = round(sense.get_temperature_from_pressure(), 1)
temp_avg = round((temp_humidity + temp_pressure) / 2, 1)
print("Temperature {0:.1f} (H {1:.1f} / P {2:.1f})".format(temp_avg, temp_humidity, temp_pressure))

humidity_feed = get_or_create_feed("bedroom-humidity", "Bedroom Humidity")
pressure_feed = get_or_create_feed("bedroom-pressure", "Bedroom Pressure")
temp_humidity_feed = get_or_create_feed("bedroom-temperature-h", "Bedroom Temperature (H)",
                                        "Temperature from Humidity sensor")
temp_pressure_feed = get_or_create_feed("bedroom-temperature-p", "Bedroom Temperature (P)",
                                        "Temperature from Pressure sensor")
temp_feed = get_or_create_feed("bedroom-temperature", "Bedroom Temperature",
                               "Temperature average from from Pressure and Humidity sensors")

aio.send_data(humidity_feed.key, humidity, precision=1)
aio.send_data(pressure_feed.key, pressure, precision=1)
aio.send_data(temp_humidity_feed.key, temp_humidity, precision=1)
aio.send_data(temp_pressure_feed.key, temp_pressure, precision=1)
aio.send_data(temp_feed.key, temp_avg, precision=1)

# result.stdout = b"temp=42.9'C\n"
result = subprocess.run(['vcgencmd', 'measure_temp'], stdout=subprocess.PIPE)
cpu_temp = round(float(result.stdout.decode('utf-8').split("=")[1].split("'")[0]), 1)
print("CPU temperature:", cpu_temp)

# result.stdout = 0.01,0.03,0.00
result = subprocess.run("top -b | head -n 1 | awk '{print $10 $11 $12}'", shell=True, stdout=subprocess.PIPE)
la_1min, la_5min, la_15min = list(map(lambda x: round(float(x), 2), result.stdout.decode('utf-8').split(",")))
print("Load average:", la_1min, la_5min, la_15min)

# result.stdout = 971051 54910 797986 6471 118153 856424
result = subprocess.run("free --kilo | grep Mem | awk '{print $2,$3,$4,$5,$6,$7}'", shell=True, stdout=subprocess.PIPE)
mem_total_kb, mem_used_kb, mem_free_kb, mem_shared_kb, mem_cache_kb, mem_available_kb = list(
    map(int, result.stdout.decode('utf-8').split(" ")))
print(
    "Memory:\n\ttotal: {0}kb\n\tused: {1}kb\n\tfree: {2}kb\n\tshared: {3}kb\n\tcache: {4}kb\n\tavailable: {5}kb".format(
        mem_total_kb, mem_used_kb, mem_free_kb, mem_shared_kb, mem_cache_kb, mem_available_kb))

# result.stdout = 7386872 1397124 5657820 20%
result = subprocess.run("df | grep root | awk '{print $2,$3,$4,$5}'", shell=True, stdout=subprocess.PIPE)
disk_total_kb, disk_used_kb, disk_available_kb, disk_use_prct = list(
    map(lambda x: int(x.strip('%\n')), result.stdout.decode('utf-8').split(" ")))
print(
    "Disk:\n\ttotal: {0}kb\n\tused: {1}kb\n\tavailable: {2}kb\n\tuse: {3}%".format(
        disk_total_kb, disk_used_kb, disk_available_kb, disk_use_prct))

raw = {
    "home/bedroom/climate": dict(humidity=humidity, pressure=pressure, temp_humidity=temp_humidity,
                                 temp_pressure=temp_pressure, temp_avg=temp_avg),
    "sys/bedroom/temperature": dict(cpu=cpu_temp),
    "sys/bedroom/la": dict(min1=la_1min, min5=la_5min, min15=la_15min),
    "sys/bedroom/mem": dict(total_kb=mem_total_kb, used_kb=mem_used_kb, free_kb=mem_free_kb, shared_kb=mem_shared_kb,
                            cache_kb=mem_cache_kb, available_kb=mem_available_kb),
    "sys/bedroom/disk": dict(total_kb=disk_total_kb, used_kb=disk_used_kb, available_kb=disk_available_kb,
                             use_prct=disk_use_prct),
}

msgs = []
for topic, val in raw.items():
    msgs.append(dict(
        topic=topic,
        payload=json.dumps(dict(when=now, data=val, measurement_id=measurement_id)),
        qos=QOS_AT_LEAST_ONCE),
    )

print("Publishing metrics to MQTT:", msgs)
publish.multiple(msgs, hostname=MQTT_HOST, port=MQTT_PORT, client_id=MQTT_CLIENT_ID)
print("All done!")

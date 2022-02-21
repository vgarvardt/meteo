#!/usr/bin/env python3
import hashlib
import os
import subprocess
from datetime import datetime, timezone

import paho.mqtt.publish as publish
from Adafruit_IO import Client, Feed, RequestError
from sense_hat import SenseHat

from measurement.v1.meteo_pb2 import Climate, System

import sys

ADAFRUIT_IO_USERNAME = os.environ.get('AIO_USERNAME') or sys.exit('AIO_USERNAME env var is not defined\n')
ADAFRUIT_IO_KEY = os.environ.get('AIO_KEY') or sys.exit('AIO_KEY env var is not defined\n')

# MQTT_HOST = os.environ.get('MQTT_HOST') or sys.exit('MQTT_HOST env var is not defined\n')
# MQTT_PORT = int(os.environ.get('MQTT_PORT', 1883))
# MQTT_CLIENT_ID = os.environ.get('MQTT_HOST', 'bedroom-sense-hat')

QOS_AT_MOST_ONCE = 0
QOS_AT_LEAST_ONCE = 1
QOS_EXACTLY_ONCE = 2

now = datetime.now(timezone.utc).astimezone().isoformat()
measurement_id = hashlib.sha256(now.encode('utf-8')).hexdigest()


def get_or_create_feed(client, key, name, description=''):
    """
    :param client: Client
    :param key: string
    :param name: string
    :param description: string
    :return: Feed
    """
    try:
        return client.feeds(key)
    except RequestError as e:
        print("Got an error when tried to get feed", name, e)
        feed = Feed(key=key, name=name, description=description)
        return client.create_feed(feed)


def get_climate(sense):
    """
    :param sense: SenseHat
    :return: Climate
    """

    result = Climate()

    result.humidity = sense.get_humidity()
    result.pressure = sense.get_pressure()
    result.temperature_humidity = sense.get_temperature_from_humidity()
    result.temperature_pressure = sense.get_temperature_from_pressure()

    return result


def send_climate_to_ada(climate, client):
    """
    :param climate: Climate
    :param client: Client
    """

    humidity = round(climate.humidity, 1)
    print("Humidity", humidity)
    pressure = round(climate.pressure, 1)
    print("Pressure", pressure)
    temp_humidity = round(climate.temperature_humidity, 1)
    temp_pressure = round(climate.temperature_pressure, 1)
    temp_avg = round((temp_humidity + temp_pressure) / 2, 1)
    print("Temperature {0:.1f} (H {1:.1f} / P {2:.1f})".format(temp_avg, temp_humidity, temp_pressure))

    humidity_feed = get_or_create_feed(client, "bedroom-humidity", "Bedroom Humidity")
    pressure_feed = get_or_create_feed(client, "bedroom-pressure", "Bedroom Pressure")
    temp_humidity_feed = get_or_create_feed(client, "bedroom-temperature-h", "Bedroom Temperature (H)",
                                            "Temperature from Humidity sensor")
    temp_pressure_feed = get_or_create_feed(client, "bedroom-temperature-p", "Bedroom Temperature (P)",
                                            "Temperature from Pressure sensor")
    temp_feed = get_or_create_feed(client, "bedroom-temperature", "Bedroom Temperature",
                                   "Temperature average from from Pressure and Humidity sensors")

    aio.send_data(humidity_feed.key, humidity, precision=1)
    aio.send_data(pressure_feed.key, pressure, precision=1)
    aio.send_data(temp_humidity_feed.key, temp_humidity, precision=1)
    aio.send_data(temp_pressure_feed.key, temp_pressure, precision=1)
    aio.send_data(temp_feed.key, temp_avg, precision=1)


def get_system():
    """
    :return: System
    """

    result = System()

    # out.stdout = b"temp=42.9'C\n"
    out = subprocess.run(['vcgencmd', 'measure_temp'], stdout=subprocess.PIPE)
    result.cpu_temperature = float(out.stdout.decode('utf-8').split("=")[1].split("'")[0])

    # top -b | head -n 1 -> top - 12:39:15 up 44 min,  1 user,  load average: 0.20, 0.84, 0.81
    # out.stdout = 0.03,0.56,0.71
    out = subprocess.run("top -b | head -n 1 | awk '{print $11 $12 $13}'", shell=True, stdout=subprocess.PIPE)
    result.LA.min1, result.LA.min5, result.LA.min15 = list(map(lambda x: float(x), out.stdout.decode('utf-8').split(",")))

    # out.stdout = 971051 54910 797986 6471 118153 856424
    out = subprocess.run("free --kilo | grep Mem | awk '{print $2,$3,$4,$5,$6,$7}'", shell=True, stdout=subprocess.PIPE)
    result.Memory.total_kb, result.Memory.used_kb, result.Memory.free_kb, result.Memory.shared_kb, result.Memory.cache_kb, result.Memory.available_kb = list(
        map(int, out.stdout.decode('utf-8').split(" ")))

    # out.stdout = 7386872 1397124 5657820 20%
    out = subprocess.run("df | grep root | awk '{print $2,$3,$4,$5}'", shell=True, stdout=subprocess.PIPE)
    result.Disk.total_kb, result.Disk.used_kb, result.Disk.available_kb, result.Disk.use_prct = list(
        map(lambda x: int(x.strip('%\n')), out.stdout.decode('utf-8').split(" ")))

    return result


# raw = {
#     "home/bedroom/climate": dict(humidity=humidity, pressure=pressure, temp_humidity=temp_humidity,
#                                  temp_pressure=temp_pressure, temp_avg=temp_avg),
#     "sys/bedroom/temperature": dict(cpu=cpu_temp),
#     "sys/bedroom/la": dict(min1=la_1min, min5=la_5min, min15=la_15min),
#     "sys/bedroom/mem": dict(total_kb=mem_total_kb, used_kb=mem_used_kb, free_kb=mem_free_kb, shared_kb=mem_shared_kb,
#                             cache_kb=mem_cache_kb, available_kb=mem_available_kb),
#     "sys/bedroom/disk": dict(total_kb=disk_total_kb, used_kb=disk_used_kb, available_kb=disk_available_kb,
#                              use_prct=disk_use_prct),
# }
#
# msgs = []
# for topic, val in raw.items():
#     msgs.append(dict(
#         topic=topic,
#         payload=json.dumps(dict(when=now, data=val, measurement_id=measurement_id)),
#         qos=QOS_AT_LEAST_ONCE),
#     )
#
# print("Publishing metrics to MQTT:", msgs)
# # publish.multiple(msgs, hostname=MQTT_HOST, port=MQTT_PORT, client_id=MQTT_CLIENT_ID)
# print("All done!")

if __name__ == '__main__':
    sense = SenseHat()
    sense.clear()

    # Create an instance of the REST client.
    aio = Client(ADAFRUIT_IO_USERNAME, ADAFRUIT_IO_KEY)

    climate = get_climate(sense)
    send_climate_to_ada(climate, aio)

    system = get_system()

    cpu_temp = round(system.cpu_temperature, 1)
    print("CPU temperature:", cpu_temp)

    la_1min, la_5min, la_15min = round(system.LA.min1, 2), round(system.LA.min5, 2), round(system.LA.min15, 2)
    print("Load average:", la_1min, la_5min, la_15min)

    print("Memory:\n\ttotal: {0}kb\n\tused: {1}kb\n\tfree: {2}kb\n\tshared: {3}kb\n\tcache: {4}kb\n\tavailable: {5}kb".format(
        system.Memory.total_kb, system.Memory.used_kb, system.Memory.free_kb, system.Memory.shared_kb, system.Memory.cache_kb, system.Memory.available_kb
    ))

    print("Disk:\n\ttotal: {0}kb\n\tused: {1}kb\n\tavailable: {2}kb\n\tuse: {3}%".format(
        system.Disk.total_kb, system.Disk.used_kb, system.Disk.available_kb, system.Disk.use_prct
    ))

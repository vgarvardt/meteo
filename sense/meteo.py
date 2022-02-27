#!/usr/bin/env python3
import hashlib
import os
import subprocess
import sys
from datetime import datetime, timezone

import paho.mqtt.publish as publish
from Adafruit_IO import Client, Feed, RequestError
from sense_hat import SenseHat
from google.protobuf.timestamp_pb2 import Timestamp

from measurement.v1.meteo_pb2 import Climate, System, Measurement

ADAFRUIT_IO_USERNAME = os.environ.get('AIO_USERNAME') or sys.exit('AIO_USERNAME env var is not defined\n')
ADAFRUIT_IO_KEY = os.environ.get('AIO_KEY') or sys.exit('AIO_KEY env var is not defined\n')

MQTT_HOST = os.environ.get('MQTT_HOST', 'localhost')
MQTT_PORT = int(os.environ.get('MQTT_PORT', 1883))
MQTT_CLIENT_ID = os.environ.get('MQTT_HOST', 'bedroom-sense-hat')

QOS_AT_MOST_ONCE = 0
QOS_AT_LEAST_ONCE = 1
QOS_EXACTLY_ONCE = 2


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


def create_measurement(now, id):
    """
    :param now: Timestamp
    :param id: string
    :return: Meta
    """

    return Measurement(time=now, id=id)


def get_climate(sense, now, measurement_id):
    """
    :param sense: SenseHat
    :param now: Timestamp
    :param measurement_id: string
    :return: Climate
    """

    return Climate(
        measurement=create_measurement(now, measurement_id),
        humidity=sense.get_humidity(),
        pressure=sense.get_pressure(),
        temperature_humidity=sense.get_temperature_from_humidity(),
        temperature_pressure=sense.get_temperature_from_pressure(),
    )


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


def get_system(now, measurement_id):
    """
    :param now: Timestamp
    :param measurement_id: string
    :return: System
    """

    result = System(
        measurement=create_measurement(now, measurement_id),
    )

    # out.stdout = b"temp=42.9'C\n"
    out = subprocess.run(['vcgencmd', 'measure_temp'], stdout=subprocess.PIPE)
    result.cpu_temperature = float(out.stdout.decode('utf-8').split("=")[1].split("'")[0])

    # top -b | head -n 1 -> top - 12:39:15 up 44 min,  1 user,  load average: 0.20, 0.84, 0.81
    # top -b | head -n 1 -> top - 08:22:02 up 2 days, 19:17,  0 users,  load average: 0.53, 0.70, 0.74
    out = subprocess.run("top -b | head -n 1", shell=True, stdout=subprocess.PIPE)
    _, la_str = list(map(lambda s: s.strip(), out.stdout.decode('utf-8').split("average:")))
    result.la = result.LA()
    result.la.min1, result.la.min5, result.la.min15 = list(map(lambda x: float(x), la_str.split(",")))

    # out.stdout = 971051 54910 797986 6471 118153 856424
    out = subprocess.run("free --kilo | grep Mem | awk '{print $2,$3,$4,$5,$6,$7}'", shell=True, stdout=subprocess.PIPE)
    result.memory = result.Memory()
    result.memory.total_kb, result.memory.used_kb, result.memory.free_kb, result.memory.shared_kb, result.memory.cache_kb, result.memory.available_kb = list(
        map(int, out.stdout.decode('utf-8').split(" ")))

    # out.stdout = 7386872 1397124 5657820 20%
    out = subprocess.run("df / | tail -n1 | awk '{print $2,$3,$4,$5}'", shell=True, stdout=subprocess.PIPE)
    result.disk = result.Disk()
    result.disk.total_kb, result.disk.used_kb, result.disk.available_kb, result.disk.use_prct = list(
        map(lambda x: int(x.strip('%\n')), out.stdout.decode('utf-8').split(" ")))

    return result


def publish_to_mqtt(climate, system):
    """
    :param climate: Climate
    :param system: System
    :return:
    """
    msgs = (
        dict(topic='climate/bedroom', payload=climate.SerializeToString(), qos=QOS_AT_LEAST_ONCE),
        dict(topic='sys/bedroom', payload=system.SerializeToString(), qos=QOS_AT_LEAST_ONCE),
    )
    print("Publishing metrics to MQTT:", msgs)
    publish.multiple(msgs, hostname=MQTT_HOST, port=MQTT_PORT, client_id=MQTT_CLIENT_ID)
    print("All done!")


if __name__ == '__main__':
    measurement_id = hashlib.sha256(datetime.now(timezone.utc).astimezone().isoformat().encode('utf-8')).hexdigest()
    now = Timestamp()
    now.GetCurrentTime()

    sense = SenseHat()
    sense.clear()

    # Create an instance of the REST client.
    aio = Client(ADAFRUIT_IO_USERNAME, ADAFRUIT_IO_KEY)

    climate = get_climate(sense, now, measurement_id)
    send_climate_to_ada(climate, aio)

    system = get_system(now, measurement_id)

    cpu_temp = round(system.cpu_temperature, 1)
    print("CPU temperature:", cpu_temp)

    la_1min, la_5min, la_15min = round(system.la.min1, 2), round(system.la.min5, 2), round(system.la.min15, 2)
    print("Load average:", la_1min, la_5min, la_15min)

    print("Memory:\n\ttotal: {0}kb\n\tused: {1}kb\n\tfree: {2}kb\n\tshared: {3}kb\n\tcache: {4}kb\n\tavailable: {5}kb".format(
        system.memory.total_kb, system.memory.used_kb, system.memory.free_kb, system.memory.shared_kb, system.memory.cache_kb, system.memory.available_kb
    ))

    print("Disk:\n\ttotal: {0}kb\n\tused: {1}kb\n\tavailable: {2}kb\n\tuse: {3}%".format(
        system.disk.total_kb, system.disk.used_kb, system.disk.available_kb, system.disk.use_prct
    ))

    publish_to_mqtt(climate, system)

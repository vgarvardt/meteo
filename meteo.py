#!/usr/bin/env python3
import json
import os
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

msgs = [
    {
        'topic': "home/bedroom/humidity",
        'payload': json.dumps(dict(when=now, data=humidity)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "home/bedroom/pressure",
        'payload': json.dumps(dict(when=now, data=pressure)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "home/bedroom/temperature-h",
        'payload': json.dumps(dict(when=now, data=temp_humidity)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "home/bedroom/temperature-p",
        'payload': json.dumps(dict(when=now, data=temp_pressure)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "home/bedroom/temperature-2",
        'payload': json.dumps(dict(when=now, data=temp_avg)),
        'qos': QOS_AT_LEAST_ONCE,
    },
]

publish.multiple(msgs, hostname=MQTT_HOST, port=MQTT_PORT, client_id=MQTT_CLIENT_ID)

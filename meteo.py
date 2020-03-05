#!/usr/bin/env python3
import os
import sys

from Adafruit_IO import Client, Feed, RequestError
from sense_hat import SenseHat

ADAFRUIT_IO_USERNAME = os.environ.get('AIO_USERNAME') or sys.exit('AIO_USERNAME env var is not defined\n')
ADAFRUIT_IO_KEY = os.environ.get('AIO_KEY') or sys.exit('AIO_KEY env var is not defined\n')

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

humidity = sense.get_humidity()
print("Humidity", round(humidity, 1))
pressure = sense.get_pressure()
print("Pressure", round(pressure, 1))
temp_humidity = sense.get_temperature_from_humidity()
temp_pressure = sense.get_temperature_from_pressure()
temp = (temp_humidity + temp_pressure) / 2
print("Temperature {0:.1f} (H {1:.1f} / P {2:.1f})".format(round(temp, 1),
                                                           round(temp_humidity, 1),
                                                           round(temp_pressure, 1)))

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
aio.send_data(temp_feed.key, temp, precision=1)

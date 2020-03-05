#!/usr/bin/env python3
import paho.mqtt.publish as publish

QOS_AT_MOST_ONCE = 0
QOS_AT_LEAST_ONCE = 1
QOS_EXACTLY_ONCE = 2

msgs = [
    {'topic': "home/bedroom/humidity", 'payload': 12.34, 'qos': QOS_AT_LEAST_ONCE, 'retain': False},
    {'topic': "home/bedroom/pressure", 'payload': 23.45, 'qos': QOS_AT_LEAST_ONCE, 'retain': False},
    {'topic': "home/bedroom/temperature-h", 'payload': 34.56, 'qos': QOS_AT_LEAST_ONCE, 'retain': False},
    {'topic': "home/bedroom/temperature-p", 'payload': 45.67, 'qos': QOS_AT_LEAST_ONCE, 'retain': False},
    {'topic': "home/bedroom/temperature-2", 'payload': 56.78, 'qos': QOS_AT_LEAST_ONCE, 'retain': False},
]

publish.multiple(msgs, hostname='localhost', port=1883, client_id='bedroom-sense-hat')

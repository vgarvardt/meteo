#!/usr/bin/env python3
import paho.mqtt.subscribe as subscribe


def on_message_print(client, userdata, message):
    print("%s %s" % (message.topic, message.payload))


subscribe.callback(on_message_print, "home/#",
                   hostname='localhost', port=1883, client_id='test-subscriber-111', clean_session=False, qos=1)

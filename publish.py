#!/usr/bin/env python3
import json
from datetime import datetime, timezone

import paho.mqtt.publish as publish

QOS_AT_MOST_ONCE = 0
QOS_AT_LEAST_ONCE = 1
QOS_EXACTLY_ONCE = 2

now = datetime.now(timezone.utc).astimezone().isoformat()

msgs = [
    {
        'topic': "home/bedroom/humidity",
        'payload': json.dumps(dict(when=now, data=12.34)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "home/bedroom/pressure",
        'payload': json.dumps(dict(when=now, data=23.45)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "home/bedroom/temperature-h",
        'payload': json.dumps(dict(when=now, data=34.56)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "home/bedroom/temperature-p",
        'payload': json.dumps(dict(when=now, data=45.67)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "home/bedroom/temperature-2",
        'payload': json.dumps(dict(when=now, data=56.78)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "sys/bedroom/la-1",
        'payload': json.dumps(dict(when=now, data=0.00)),
        'qos': QOS_AT_LEAST_ONCE,
    },
    {
        'topic': "sys/bedroom/la-15",
        'payload': json.dumps(dict(when=now, data=0.03)),
        'qos': QOS_AT_LEAST_ONCE,
    },
]

publish.multiple(msgs, hostname='localhost', port=1883, client_id='bedroom-sense-hat')

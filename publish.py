#!/usr/bin/env python3
import hashlib
import json
from datetime import datetime, timezone

import paho.mqtt.publish as publish

QOS_AT_MOST_ONCE = 0
QOS_AT_LEAST_ONCE = 1
QOS_EXACTLY_ONCE = 2

now = datetime.now(timezone.utc).astimezone().isoformat()
measurement_id = hashlib.sha256(now.encode('utf-8')).hexdigest()

raw = {
    "home/bedroom/climate": dict(humidity=12.34, pressure=23.45, temp_humidity=34.56,
                                 temp_pressure=45.67, temp_avg=56.78),
    "sys/bedroom/temperature": dict(cpu=123),
    "sys/bedroom/la": dict(min1=234, min5=345, min15=456),
    "sys/bedroom/mem": dict(total_kb=1234, used_kb=2345, free_kb=3456, shared_kb=4567,
                            cache_kb=5678, available_kb=6789),
    "sys/bedroom/disk": dict(total_kb=4321, used_kb=5432, available_kb=6543,
                             use_prct=7654),
}

msgs = []
for topic, val in raw.items():
    msgs.append(dict(
        topic=topic,
        payload=json.dumps(dict(when=now, data=val, measurement_id=measurement_id)),
        qos=QOS_AT_LEAST_ONCE),
    )

publish.multiple(msgs, hostname='localhost', port=1883, client_id='bedroom-sense-hat')

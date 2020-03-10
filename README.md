# meteo

## How to start Docker containers at boot time

Source: https://toub.es/2017/08/08/how-to-start-a-docker-container-at-boot-time/ (https://web.archive.org/web/20180831050539/https://toub.es/2017/08/08/how-to-start-a-docker-container-at-boot-time/) 

Copy file `rpi-meteo.service` to `/etc/systemd/system/`.

Enable the service on boot:

```
sudo systemctl enable rpi-meteo.service
```

Check current status

```
sudo systemctl status rpi-meteo.service
```

Disable the service from boot

```
sudo systemctl disable rpi-meteo.service
```

## How to disable LED on Pi 3

Source: https://www.raspberrypi.org/forums/viewtopic.php?t=149126

Add the following to the `/boot/config.txt`:

```
# Disable Ethernet LEDs
dtparam=eth_led0=14
dtparam=eth_led1=14

# Disable the PWR LED
dtparam=pwr_led_trigger=none
dtparam=pwr_led_activelow=off

# Disable the Activity LED
dtparam=act_led_trigger=none
```

# meteo

## How to start Docker containers at boot time

Source: https://toub.es/2017/08/08/how-to-start-a-docker-container-at-boot-time/ (https://web.archive.org/web/20180831050539/https://toub.es/2017/08/08/how-to-start-a-docker-container-at-boot-time/) 

Copy file `rpi-meteo` to `/etc/systemd/system/`.

Enable the service on boot:

```
sudo systemctl enable rpi-meteo
```

Check current status

```
sudo systemctl status rpi-meteo
```

Disable the service from boot

```
sudo systemctl disable rpi-meteo
```

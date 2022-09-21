#!/bin/bash
chmod +x redis-checkhealth
mv redis-checkhealth /usr/bin/redis-checkhealth
cp redis-checkhealth.service /usr/lib/systemd/system/
mv .redis-ch.conf /etc/
systemctl daemon-reload
systemctl enable redis-checkhealth --now
systemctl start redis-checkhealth
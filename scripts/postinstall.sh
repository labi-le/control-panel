#!/bin/bash
sudo touch /etc/systemd/system/control-panel.service
cat <<EOF > /etc/systemd/system/control-panel.service
[Unit]
Description=control-panel
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service
[Service]
Type=simple
ExecStart=/opt/control-panel/panel
WorkingDirectory=/opt/control-panel/
[Install]
WantedBy=multi-user.target
EOF
sudo chmod 664 /etc/systemd/system/control-panel.service
sudo systemctl daemon-reload
sudo systemctl enable --now control-panel.service
sudo systemctl restart control-panel.service

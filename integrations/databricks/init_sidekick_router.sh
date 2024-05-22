#!/bin/bash
set -ex

# Check if sidekick bin is present, if not download it
SIDEKICK_BIN=/usr/bin/sidekick-router
if [ -f "$SIDEKICK_BIN" ]; then
    echo "$SIDEKICK_BIN already installed."
else
    wget https://github.com/project-n-oss/sidekick-router/releases/latest/download/sidekick-router-linux-amd64.tar.gz
    tar -xzvf sidekick-router-linux-amd64.tar.gz -C /usr/bin
fi
chmod +x $SIDEKICK_BIN
$SIDEKICK_BIN --help > /dev/null

cat > /databricks/driver/conf/style-path-spark-conf.conf <<EOL
[driver] {
  "spark.hadoop.fs.s3a.path.style.access" = "true"
  "spark.hadoop.fs.s3a.bucket.<BUCKET_NAME>.endpoint" = "http://localhost:7075"
  "spark.hadoop.fs.s3a.bucket.<BUCKET_NAME>.endpoint.region" = "<BUCKET_REGION>"
}
EOL

# Add any spark or env config here:
# --------------------------------------------------

# --------------------------------------------------

# Create service file for the sidekick-router process
SERVICE_FILE="/etc/systemd/system/sidekick-router.service"
touch $SERVICE_FILE

cat > $SERVICE_FILE << EOF
[Unit]
Description=Sidekick Router service file

[Service]
Environment=SIDEKICKROUTER_APP_NOCRUNCHFAILOVER=true
Environment=SIDEKICKROUTER_APP_CLOUDPLATFORM=AWS
ExecStart=$SIDEKICK_BIN serve -p 7075
Restart=always

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable sidekick-router
systemctl start sidekick-router


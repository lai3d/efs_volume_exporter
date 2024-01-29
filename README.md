## EFS Volume Exporter


[![Docker](https://img.shields.io/badge/docker-master-brightgreen.svg)](https://hub.docker.com/repository/docker/dockerpandamaster/efs_volume_exporter) 
[![Docker Pulls](https://img.shields.io/docker/pulls/dockerpandamaster/efs_volume_exporter.svg)](https://hub.docker.com/r/dockerpandamaster/efs_volume_exporter)

Only for AWS EFS volume monitoring!

## Wiki

Refer project [wiki](https://github.com/lai3d/efs_volume_exporter/wiki) for more details

## Running


### Docker Locally

```bash 
docker run --rm -p 9888:9888 -it dockerpandamaster/efs_volume_exporter --volume-dir=bin:/bin
```
### Deploy It in Cloud
Add as a sidecar

```yaml 
        - name: volume-exporter
          image:  dockerpandamaster/efs_volume_exporter
          imagePullPolicy: "Always"
          args:
            - --volume-dir=prometheus:/prometheus
          ports:
          - name: metrics-volume
            containerPort: 9888
          volumeMounts:
          - mountPath: /prometheus
            name: prometheus-data
            readOnly: true
```

## Config

|Flag |	Description|
| ---------------------------- | -------------------------------------------- | 
| web.listen-address |	Address to listen on for web interface and telemetry. Default is 9888|
| web.telemetry-path |	Path under which to expose metrics. Default is /metrics|
| volume-dir	 | volumes to report, the format is volumeName:VolumeDir, For example ==> logs:/app/logs, you can use this flag multiple times to provide multiple volumes|


## Exporterd Metrics

| metrics	| Type |	Description |
| --------------------------------------------------------- | ----------- |  ------------------------------------- |
| volume_bytes_used{volume_name=”someName”, volume_path=”/some/path”} |	Gauge |	Used size of volume/disk | 



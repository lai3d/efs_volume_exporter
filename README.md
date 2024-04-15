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

### Deploy It in Kubernetes

Add as a sidecar in deployment

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

Expose port 9888 in service

```yaml
  - name: metrics-volume
    port: 9888
    protocol: TCP
    targetPort: 9888
```

Complete Service example

```yaml
apiVersion: v1
kind: Service
metadata:
  name: cdn
spec:
  ports:
  - port: 80
    targetPort: 80
    protocol: TCP
    name: http
  - name: metrics-volume
    port: 9888
    protocol: TCP
    targetPort: 9888
  selector:
    app: nginx
```

In prometheus.yml, add a job to scrape metrics

```yaml
    scrape_configs:
    - job_name: 'cdn_volume_metrics'

      scrape_interval: 15s
      scrape_timeout: 5s

      static_configs:
        - targets: ['cdn.staging:9888']
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



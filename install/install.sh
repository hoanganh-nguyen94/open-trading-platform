

VERSION="1.0.19"
DOCKERREPO="ettec/opentp:"
TAG=-$VERSION

#Kafka

echo installing Kafka...

kubectl create ns kafka

helm repo add bitnami https://charts.bitnami.com/bitnami --force-update

helm install kafka-opentp --namespace kafka bitnami/kafka --version 21.4.6

kubectl apply --wait -f kafka_cmdline_client.yaml


#Postgres

echo installing Postgresql database...

kubectl create ns postgresql

helm install opentp --wait --namespace postgresql bitnami/postgresql --set-file pgHbaConfiguration=./pb_hba_no_sec.conf --set volumePermissions.enabled=true

echo loading data into Postgresql database...
export POSTGRES_PASSWORD=$(kubectl get secret --namespace postgresql opentp-postgresql -o jsonpath="{.data.postgres-password}" | base64 --decode)

kubectl run opentp-postgresql-client --rm --tty -i --restart='Never' --namespace postgresql --image hoanganhnguyen1994/dataload:postgresql-15-data-loader-client-1.0.19 --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql --host opentp-postgresql -U postgres -d postgres -p 5432 -a -f ./opentp.db

#Envoy
echo installing Envoy...

kubectl create ns envoy
helm install opentp-envoy --wait --namespace=envoy ./charts/envoy -f envoy-config-helm-values.yaml 


kubectl patch service envoy --namespace envoy --type='json' -p='[{"op": "replace", "path": "/spec/sessionAffinity", "value": "ClientIP"}]'


#Orders topic
kubectl exec -it --namespace=kafka cmdlineclient -- sh -c "kafka-topics --zookeeper kafka-opentp-zookeeper:2181 --topic orders --create --partitions 1 --replication-factor 1"
kubectl exec -it --namespace=kafka kafka-opentp-client -- bash -c "kafka-topics.sh --zookeeper zookeeper-opentp:2181 --topic orders --create --partitions 1 --replication-factor 1"

kubectl exec --tty -i kafka-opentp-client --namespace kafka -- bash
kafka-topics.sh --create --zookeeper kafka-topics --zookeeper zookeeper-opentp:2181 --topic orders --create --partitions 1 --replication-factor 1

kubectl run kafka-opentp-client --restart='Never' --image docker.io/bitnami/kafka:3.5.1-debian-11-r11 --namespace kafka --command -- sleep infinity
kubectl cp --namespace kafka /path/to/client.properties kafka-opentp-client:/tmp/client.properties
kubectl exec --tty -i kafka-opentp-client --namespace kafka -- sh -c "kafka-topics --zookeeper zookeeper-opentp:2181 --topic orders --create --partitions 1 --replication-factor 1"

kubectl --namespace default exec -it $POD_NAME -- kafka-topics.sh --create --zookeeper ZOOKEEPER-SERVICE-NAME:2181 --replication-factor 1 --partitions 1 --topic mytopic
kubectl get pods --namespace kafka -l "app.kubernetes.io/name=kafka,app.kubernetes.io/instance=kafka,app.kubernetes.io/component=kafka" -o jsonpath="{.items[0].metadata.name}"


#Opentp

echo installing Open Trading Platform...

kubectl create ns opentp
helm install --wait --timeout 1200s otp-1.0.19 ../helm-otp-chart/ --set dockerRepo="ettec/opentp:" --set dockerTag="-1.0.19" --namespace opentp

#Instructions to start client
OTPPORT=$(kubectl get svc --namespace=envoy -o go-template='{{range .items}}{{range.spec.ports}}{{if .nodePort}}{{.nodePort}}{{"\n"}}{{end}}{{end}}{{end}}')

echo Open Trading Platform is running. To start a client point your browser at port $OTPPORT and login as trader1 








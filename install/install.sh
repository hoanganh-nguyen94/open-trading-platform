
DOCKERREPO="ettec/opentp:"
TAG=-$VERSION

#Kafka

echo installing Kafka...

kubectl create ns kafka

helm repo add bitnami https://charts.bitnami.com/bitnami

helm install kafka-opentp --wait --namespace kafka  bitnami/kafka --version 21.4.6

#install kafka cmd line client


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
kubectl exec -it --namespace=kafka cmdlineclient -- /bin/bash -c "kafka-topics --zookeeper kafka-opentp-zookeeper:2181 --topic orders --create --partitions 1 --replication-factor 1"
kubectl exec -it --namespace=kafka cmdlineclient -- sh -c "kafka-topics --zookeeper kafka-opentp-zookeeper:2181 --topic orders --create --partitions 1 --replication-factor 1"


#Opentp

echo installing Open Trading Platform...

kubectl create ns opentp
helm install --wait --timeout 1200s otp-1.0.19 ../helm-otp-chart/ --set dockerRepo="ettec/opentp:" --set dockerTag="-1.0.19" --namespace opentp

#Instructions to start client
export OTPPORT=$(kubectl get svc --namespace=envoy -o go-template='{{range .items}}{{range.spec.ports}}{{if .nodePort}}{{.nodePort}}{{"\n"}}{{end}}{{end}}{{end}}')
echo $OTPPORT

echo Open Trading Platform is running. To start a client point your browser at port $OTPPORT and login as trader1 








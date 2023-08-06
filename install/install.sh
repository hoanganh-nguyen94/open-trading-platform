

VERSION="1.0.19"
DOCKERREPO="ettec/opentp:"
TAG=-$VERSION

#Kafka

echo installing Kafka...

kubectl create ns kafka

helm repo add bitnami https://charts.bitnami.com/bitnami --force-update

helm install zookeeper-opentp --namespace kafka bitnami/zookeeper --version 11.4.9 -set replicaCount=3 --set auth.enabled=false --set allowAnonymousLogin=true
helm install kafka-opentp --wait --namespace kafka bitnami/kafka --version 23.0.7 --set zookeeperChrootPath=zookeeper-opentp.kafka.svc.cluster.local

#install kafka cmd line client

 
kubectl apply --wait -f kafka_cmdline_client.yaml


#Postgres

echo installing Postgresql database...

kubectl create ns postgresql

helm install opentp --wait --namespace postgresql bitnami/postgresql --set-file pgHbaConfiguration=./pb_hba_no_sec.conf --set volumePermissions.enabled=true

echo loading data into Postgresql database...
export POSTGRES_PASSWORD=$(kubectl get secret --namespace postgresql opentp-postgresql -o jsonpath="{.data.postgres-password}" | base64 --decode)

kubectl run opentp-postgresql-client --rm --tty -i --restart='Never' --namespace postgresql --image  ettec/opentp:data-loader-client-1.0.19 --env="POSTGRESQL_PASSWORD=$POSTGRES_PASSWORD" --command -- psql --host opentp-postgresql -U postgres -d postgres -p 5432 -a -f ./opentp.db

#Envoy

echo installing Envoy...

kubectl create ns envoy
helm install opentp-envoy --wait --namespace=envoy ./charts/envoy -f envoy-config-helm-values.yaml 


kubectl patch service envoy --namespace envoy --type='json' -p='[{"op": "replace", "path": "/spec/sessionAffinity", "value": "ClientIP"}]'


#Orders topic
kubectl exec -it --namespace=kafka cmdlineclient -- sh -c "kafka-topics --zookeeper zookeeper-opentp:2181 --topic orders --create --partitions 1 --replication-factor 1"


#Opentp

echo installing Open Trading Platform...


helm install --wait --timeout 1200s otp-1.0.19 ../helm-otp-chart/ --set dockerRepo="ettec/opentp:" --set dockerTag="-1.0.19"

#Instructions to start client
OTPPORT=$(kubectl get svc --namespace=envoy -o go-template='{{range .items}}{{range.spec.ports}}{{if .nodePort}}{{.nodePort}}{{"\n"}}{{end}}{{end}}{{end}}')

echo Open Trading Platform is running. To start a client point your browser at port $OTPPORT and login as trader1 








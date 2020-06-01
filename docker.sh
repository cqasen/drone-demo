#!/bin/bash
while [[ $# -ge 1 ]]; do
	case $1 in
		-n|--name )
			name=$2
			shift 2
			;;
		-r|--replicas )
			replicas=$2
			shift 2
			;;
		-p|--publish )
			publish=$2
			shift
			;;
    -c|--container )
      container=$2
      shift
      ;;
    -t|--tag )
      tag=$2
      shift
      ;;
		* )
			shift
			;;
	esac
done
echo "name = $name"
echo "replicas = $replicas"
echo "publish = $publish"
echo "container = $container"
echo "tag = $tag"

SERVICE_ID=`docker service ps $name|grep $name |awk '{print $2}'`
if [ -n "${SERVICE_ID}" ];then
  echo "更新服务"
  docker service update --image $container:$tag $name
else
  echo "创建服务"
  docker service create --name $name --replicas $replicas -t --publish $publish $container:$tag
fi
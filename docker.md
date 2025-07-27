if i expose port in Dockerfile than it choose random port to communicate with the other containers
docker pull
docker create
docker run
docker stop
docker start
docker commit <container-id> name:tag //this is for creating image from an running image

why we need docker?why container based softwae delivery?

bare metals,vms,containers  vm have kernal but container dont
Docker under the hood:
Namespaces --> mnt/file,network,uts,pid,ipc,user

Control gorups,Cgroups-->we can control resources and monitorring  with this like CPU,Memory,Disk,Network

Union+overlay filesystem-->when i pull a images i pill all data but seconed time only the changes and it also suit for building ,,,overlay 2 is stoage type


for image history we can see docker image history image_name

Copy on write-->>each container create with copied files from images and we can change the content and see with docker diff image_name,and commit that with docker commit container_id new_name

dockerd provide all like immages,logs,stats,
containerd is the runtime to run image

Docker Architecture:

docker engine+docker client its a client server model

docker system info
docker system events   //server side logs

docker ps,-l,-n 2,-a
docker run -it image //interactive mode to interact with the contaienr

docker run -idt iamge

docker logs container_name  //log_path in docker inspect
docker inspect contaianer_name will inspect everything
docker exec -it container bashj
docker diff container_name or id to see difference in file
docker run -d --rm image_name to auto delete after exit of container

docker images history image_name


port mapping:
1. docker run -p 8000:80 image
2. docker run -P ..>for random port
3.docker run -p 80 image ...> here 80 is container port and will expose this port to any host port
control lifecycle of containers
docker create 
docker start
docker stop
docker rm -f  >forcefully
docker rm container1 contaienr2  -->stoped container remove
docker stop-->sigterm--10s--->sigkill


docker system df to see the resource usage

docker container prune  --> to remove all stoped containers

docker ssystem prun:

all stoped container
all dangling image
all network not used
all dangling build cache



Limiting Resources:
docker stats

docker container stats ---> show cpu ,memory net io,PID 

docker container stats --no-stream=true it will not stream 

memory limit -->docker container update -m 100M container_name

docker run -d --cpu-shares 512 image_name --cpu 1 --> cpu share is relative value,,1024 for 100%,here more shared container will get more CPU

docker cp . dontainer_id:/app --> this is fo copy the files in docker container


Docke networkng:

bridge 
host n
none
overlay
mcvlan
custom netwrok with user defined
docker network create -d bridge/overlay network name 
user defined bridge net can communicate with DNS name but default bridge network can commincate with only with ip address
docker network connect ,disconnect network name
docker network inspect netork_name


Docker volumes:
/var/lib/docker/volumes/
docker volume create volume_name whichh will use the defualt folder for store data
Named volumes which we create with a custom name like docker volume cfeate test_volume
docker run -v /host/path:/container/path nginx

docker run -v volume_name:container_pat image_name
for mount
docker run -d --name=test --mount source=volume_name,destination=data path in container image_name

we can add more thing in mount where volume provide us a few options
we can inpect a volume with docker inscpect volume name


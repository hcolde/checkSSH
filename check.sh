#/bin/bash
# replace 22 with the port you set
info=`lsof -i :22 | egrep -o '[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'`
OLD_IFS="$IFS" 
IFS="\n" 
arr=($info) 
IFS="$OLD_IFS"

ips=""
for s in ${arr[@]} 
do
	if [[ $s != "180.76.76.76" ]];then # except ur own ip
		ips=$ips$s";"
	fi
done

if [[ $ips != "" ]];then
	/root/mail -msg $ips
fi

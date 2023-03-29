#!/bin/bash
ListeningPort=`netstat -an | grep "safeis-ods-check task -c config_tron_check_online.yml" | wc -l`
if [ $ListeningPort -eq 0 ]
then
{
  echo "`date` : listener server is down">>/opt/check/eth/online/listen.log
  #nohup /opt/check/eth/online/safeis-ods-check task -c config_tron_check_online.yml > /opt/check/eth/online/out.log 2>&1 &
}
else
{
 echo "`date` : server运行正常" >>/opt/check/eth/online/listen.log
}
fi
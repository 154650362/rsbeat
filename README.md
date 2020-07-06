哨兵配置方式:    

IP:PORT:{SentinelMasterName}


rsbeat:
  # Defines how often an event is sent to the output
  period: 1s
  redis: ["127.0.01:6379:myredis"]
  slowerThan: 100
  
  
 redis cluster配置方式:    

IP:PORT:{clusterName}
//这里clusterName为自定义, 一套集群内唯一

rsbeat:
  # Defines how often an event is sent to the output
  period: 1s
  redis: ["127.0.01:6379:mycluster","127.0.01:6380:mycluster"]
  slowerThan: 100

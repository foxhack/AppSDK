{
    "encoding" : "GBK",
    "uuid": "0001-0001-600406",
    "templateVersion": "1.00",
    "type": "line monitor app v1",
    "model": "NS3900X",
    "manufactureId": "600406",
    "manufactureName": "NARI",
    "description":"line monitor app v1",
    "config": {
            "port": 2405,
            "mount": ["/home/NEdge/104file/upload:/home/NEdge/upload"]
     },
    "channel": [
       {
            "name": "input",
            "protocol":"zmq",
            "host": "localhost",
            "port": 1883,
            "user": "mosquitto",
            "password": "",
            "qos": 0,
            "keepAlive": 3600,
            "topicPath": "/Input"
        },
        {
            "name": "output",
            "protocol":"mqtt",
            "host": "localhost",
            "port": 1883,
            "user": "mosquitto",
            "password": "",
            "qos": 0,
            "keepAlive": 3600,
            "topicPath": "/Output"
        },
        {
            "name": "request",
            "protocol":"mqtt",
            "host": "localhost",
            "port": 1883,
            "user": "mosquitto",
            "password": "",
            "qos": 0,
            "keepAlive": 3600,
            "topicPath": "/Request"
        },
        {
            "name": "response",
            "protocol":"mqtt",
            "host": "localhost",
            "port": 1883,
            "user": "mosquitto",
            "password": "",
            "qos": 0,
            "keepAlive": 3600,
            "topicPath": "/Response"
        }
    ],
    "inputs": [
        {
            "name": "ycset",
            "description": "YC dataSet",
            "type": "DataSet",
            "maxResCount": 100,
            "coefficient": 1.0,
            "offset": 0,
            "reference": ""
        },
		    {
            "name": "yxset",
            "description": "YX dataSet",
            "type": "DataSet",
            "maxResCount": 10,
            "coefficient": 1.0,
            "offset": 0,
            "reference": ""
        },
		    {
            "name": "humidity",
            "description": "humidity value",
            "type": "Double",
            "maxResCount": 0,
            "coefficient": 1.0,
            "offset": 0,
            "reference": ""
        }
    ],
    "outputs": [
        {
            "name": "ResultCode",
            "description": "result code",
            "type": "String",
            "attrType":"YC"
        },
        {
            "name": "ResultDesc",
            "description": "result description",
            "type": "String",
            "attrType":"YC"
        }
	],
	"parameter": [
		{
			"name": "MaxClientNum",
			"description": "最大客户端个数",
			"type": "Int",
			"unit": "",
			"readWrite": "R|W",
			"minimum": 1,
			"maximum": 100,
			"step": 1,
			"defaultValue": 10
		},
		{
			"name": "T0",
			"description": "T0",
			"type": "Int",
			"unit": "",
			"readWrite": "R|W",
			"minimum": 0,
			"maximum": 60,
			"step": 1,
			"defaultValue": 30
		},
		{
            "name": "T1",
            "description": "T1",
            "type": "Int",
            "unit": "",
            "readWrite": "R|W",
            "minimum": 0,
            "maximum": 60,
            "step": 1,
            "defaultValue": 35
        },
        {
            "name": "T2",
            "description": "T2",
            "type": "Int",
            "unit": "",
            "readWrite": "R|W",
            "minimum": 0,
            "maximum": 60,
            "step": 1,
            "defaultValue": 10
        },
        {
            "name": "T3",
            "description": "T3",
            "type": "Int",
            "unit": "",
            "readWrite": "R|W",
            "minimum": 0,
            "maximum": 60,
            "step": 1,
            "defaultValue": 20
        },
        {
            "name": "K",
            "description": "K",
            "type": "Int",
            "unit": "",
            "readWrite": "R|W",
            "minimum": 0,
            "maximum": 60,
            "step": 1,
            "defaultValue": 12
        },
        {
            "name": "W",
            "description": "W",
            "type": "Int",
            "unit": "",
            "readWrite": "R|W",
            "minimum": 0,
            "maximum": 60,
            "step": 1,
            "defaultValue": 8
        }
	]
}

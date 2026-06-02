I want to configure logs to go to file

- get dependency to lumberjack 
- create a new structure for log configuration to be used in ConfigData
- make it use the new structure that I created in config.json as follows

"log": {
  "fileLogEnabled": true,
  "fileName": "/home/daniel/IdeaProjects/smtp-admin-ui/log/smtpadmin.log",
  "maxSize": 100,
  "maxBackups": 10,
  "maxDays": 180,
  "compress": true
}

if fileLogEnabled is not true, do not create file log
maxSize is in megabytes (if lumberjack has it different type or meaning please advise and correct) 

if fileLogEnabled is true, then enable lumber jack but keep console logging as well

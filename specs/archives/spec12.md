- Create type ConfigData
- this holds externalized configuration
- when application starts, expect environment variable CONFIG_FILE that points to the config file, which currently should be /config/config.json
- Change ConfigData to accommodate all fields from config.json at this time
- Create http server listening to the address in ConfigData
- Add dependency to mariadb sql driver
- Create sql db based on the "db" parameters in the config data. Check the connection. Does not work, get out.
- Create instance of Svc with the above sql db
- Create struct of type Controller that will hold all the endpoints
- Controller gets an instance of Svc
- Create one endpoint in Controller that will be "GET /" that returns string "ok"
- Make a folder called "endpoints" and ensure each endpoint goes in one file each
- For now create one file for the root "GET /" and we will see for later
- File for endpoint must have a name with method (post, get) and something about the endpoint like get_root



# BrickQueryEngine

- Setup Fuseki
    1. Run Fuseki standalone server (https://jena.apache.org/documentation/serving_data/#download-fuseki1)
        1.1. You may need to increase java memory limit. You can change it inside "fuseki-server".
        1.2. You may want to increase http timeout time. You can adjust it inside "/run/config.ttl". This ttl file is generated after you run the server for the first time.
    2. Register a dataset
        2.1. In a browser, go to localhost:3030
        2.2. Click Manage datasets 
        2.3. Click add new dataset
        2.4. Make a dataset with name "Brick" either in-memory or persistent. (If you do not want to set this up again, go with "persistent")
        2.5. Upload a building ttl file.
 
- Setup Mongodb
    1. Thie engine connects to a mongodb server via "127.0.0.1:27017". Make sure Mongodb listen to the correct port.

- Directory structure
    1. /queries has SPARQL queries as files (from RUN_APPS.py).
    2. /src has go src codes.
    3. /config has an example fuseki config file

 
- Test run
    1. execute "go run \src\query_engine.go" at the root directory
    2. You must see 10 random triples stored in Fuseki currently.

- Kinda architecture
    1. Code is simple and comments in query_engine.go would be enough.


- TODO
    1. Error handling
    2. Parse query result
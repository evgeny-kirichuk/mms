# mms

Setup DC1 cluster:

```
docker-compose up -d
```

Setup DC2 cluster:

```
docker-compose -f docker-compose-dc2.yml up -d
```

Setup DC3 cluster:

```
docker-compose -f docker-compose-dc3.yml up -d
```

## Fill the database with initial data:

```
CREATE KEYSPACE catalog WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','DC1' : 3};

use catalog;


CREATE TABLE mutant_data (
   first_name text,
   last_name text,
   address text,
   picture_location text,
   PRIMARY KEY((first_name, last_name)));

   insert into mutant_data ("first_name","last_name","address","picture_location") VALUES ('Bob','Loblaw','1313 Mockingbird Lane', 'http://www.facebook.com/bobloblaw');
insert into mutant_data ("first_name","last_name","address","picture_location") VALUES ('Bob','Zemuda','1202 Coffman Lane', 'http://www.facebook.com/bzemuda');
insert into mutant_data ("first_name","last_name","address","picture_location") VALUES ('Jim','Jeffries','1211 Hollywood Lane', 'http://www.facebook.com/jeffries');
```
```
CREATE KEYSPACE tracking WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','DC1' : 3};

use tracking;

CREATE TABLE tracking_data (
       first_name text,
       last_name text,
       timestamp timestamp,
       location varchar,
       speed double,
       heat double,
       telepathy_powers int,
       primary key((first_name, last_name), timestamp))
       WITH CLUSTERING ORDER BY (timestamp DESC)
       AND COMPACTION = {'class': 'TimeWindowCompactionStrategy',
           'base_time_seconds': 3600,
           'max_sstable_age_days': 1};

           INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 08:05+0000','New York',1.0,3.0,17) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 09:05+0000','New York',2.0,4.0,27) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 10:05+0000','New York',3.0,5.0,37) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 10:22+0000','New York',4.0,12.0,47) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 11:05+0000','New York',4.0,9.0,87) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 12:05+0000','New York',4.0,24.0,57) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 08:05+0000','Cincinatti',2.0,6.0,5) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 09:05+0000','Cincinatti',4.0,1.0,10) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 10:05+0000','Cincinatti',6.0,1.0,15) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 10:22+0000','Cincinatti',8.0,3.0,6) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 11:05+0000','Cincinatti',10.0,2.0,3) ;
INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 12:05+0000','Cincinatti',12.0,10.0,60) ;
```

```
exit
```

Configuring the keyspaces for Multi-DC
```
  docker exec -it scylla-node1 cqlsh

  ALTER KEYSPACE catalog WITH REPLICATION = {'class': 'NetworkTopologyStrategy', 'DC1':3, 'DC2':3};

  ALTER KEYSPACE tracking WITH REPLICATION = {'class': 'NetworkTopologyStrategy', 'DC1':3, 'DC2':3};

  ALTER KEYSPACE system_auth WITH replication = { 'class' : 'NetworkTopologyStrategy', 'DC1' : 3, 'DC2' : 3};

  ALTER KEYSPACE system_distributed WITH replication = { 'class' : 'NetworkTopologyStrategy', 'DC1' : 3, 'DC2' : 3};

  ALTER KEYSPACE system_traces WITH replication = { 'class' : 'NetworkTopologyStrategy', 'DC1' : 3, 'DC2' : 3};
```

```
  exit

  docker exec -it scylla-node4 nodetool rebuild -- DC1

  docker exec -it scylla-node5 nodetool rebuild -- DC1

  docker exec -it scylla-node6 nodetool rebuild -- DC1
```

```
  docker exec -it scylla-node4 cqlsh

  describe catalog;

  describe tracking;

  select * from catalog.mutant_data;

  select * from tracking.tracking_data;
```

### Importing the MMS keyspaces and data automatically

```
docker exec scylla-node1 cqlsh -f /mutant-data.txt
```


## Development Notes
[how to get all keyspaces](https://stackoverflow.com/questions/18712967/how-to-list-all-the-available-keyspaces-in-cassandra):
```
SELECT * FROM system_schema.keyspaces;
```
response example:
```
cqlsh:system_schema> SELECT * FROM system_schema.keyspaces;
keyspace_name| durable_writes | replication
-------------+----------------+--------------------------------
I2Tutorials  | True           | {'class': 'org.apache.cassandra.locator.SimpleStrategy', 'replication_factor':'3'}
```
[system keyspace doc](https://docs.datastax.com/en/dse/5.1/cql/cql/cql_using/useQuerySystem.html)

[gocql doc](https://pkg.go.dev/github.com/gocql/gocql)

---
[how to get all tables in keyspace](https://github.com/gocql/gocql/issues/193#issuecomment-473714574)

[TableMetadata doc](https://pkg.go.dev/github.com/gocql/gocql#TableMetadata)

---
[run docker go development environment](https://levelup.gitconnected.com/setup-simple-go-development-environment-with-docker-b8b9c0d4e0a8)
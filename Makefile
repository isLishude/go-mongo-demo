setup:
	@mkdir -p tmp/master && touch tmp/master/mongod.log
	@mkdir -p tmp/slave && touch tmp/slave/mongod.log
	mongod --replSet test --dbpath tmp/master --logpath tmp/master/mongod.log --port 27017 --fork
	mongod --replSet test --dbpath tmp/slave --logpath tmp/slave/mongod.log --port 27018 --fork
	mongo --port 27017 --eval 'rs.initiate({ _id: "test", members: [{ _id: 0, host: "localhost:27017" }, { _id: 1, host: "localhost:27018" }] });'
	mongo --port 27018 --eval 'rs.slaveOk()'

teardown:
	killall mongod
	rm -rf tmp

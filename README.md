# Go Api
A dockerized go api that saves registered links and monitors their status. The api supports registering users and links and uses jwt authorization for authorizing 
usage of the api. The monitoring part of the api uses go routines and channels to further fasten the updating process.
# Third party libraries
We used go-redis for connecting to the redis database and jwt library for managing the tokens.

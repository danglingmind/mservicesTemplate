# This directory is just an example directory showing how you can use any persistence(DB) as a third party by implementing the interface which you have defined in the "domain/repository"

## Every file is the interface implementation of the each entity, user_repository.go is concrete implementation of user's.


# You can have multiple dbs too in that case have one file for each db instead of db.go like 'postgres.go' 'mysql.go' 
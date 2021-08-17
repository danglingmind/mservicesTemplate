# Domain
## This is the domain of your project which will represent your real world entities and aggregates of your bussiness.
- *Entity* : Real world entities which will be mapped to the DB example User, Order, Employee etc.
- *Aggregates* : Complex or combination DB objects which are required for your bussiness ( might not have the same real world entity ).
- *Repository* : Repository contains all the CRUD interfaces for real entities and aggregates.
- *Service* : Service will conatain the interfaces for operations which are not CRUD but required by the 
business.

</br>
Example : 

Sample blank files are there in each directory for user entity, to give you an idea how the interface
and entity will interact.
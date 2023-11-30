# TodoList

1. Create a PGSQL database to test with this tables:

    - [x] players
    - [x] queues
    - [x] game session

2. Create a simple API REST with the following methods and verbs:

    ● players

    - [x] get all: Return a list of players.
    - [x] get by id: Return a concrete player.
    - [x] create: Return the data of the player and the ID generated.
    - [x] update by id (only name)
    - [x] delete by id

    ● queues

    - [x] get all: Returns a list of queues.
    - [x] get by id: Return to the concrete queue.
    - [x] create: Create a queue, please take into account that you can create 10  queues (by a file configuration) .Finally i used the pgSQL DB
    - [x] update by id (only name)
    - [x] delete by id

    ● game session

    - [x] get all: Returns a list of game sessions.
    - [x] create or update (it needs to specify queue ID and player ID): The logic is explained above. Return the ID of the game session and the related data.
    - [x] get by status: Return the list of queues filtered by status.

3. Create a testing for future changes
    ● players

    - [x] get all
    - [ ] get by id
    - [ ] create
    - [ ] update by id 
    - [ ] delete by id

    ● queues

    - [x] get all
    - [ ] get by id:
    - [ ] create
    - [ ] update by id 
    - [ ] delete by id

    ● game session

    - [x] get all
    - [ ] create 
    - [ ] update 
    - [ ] get by status
# Task Tracer CLI

A Task Tracker Command Line Interface created with **GO** language as a beginner project while learning the syntax and the structural typing shape of the go language. 
this project is done in reference to roadmap.sh
https://roadmap.sh/projects/task-tracker


# Running the CLI

after cloning into the project, open up a terminal and CD into the program with 
```
cd task-tracker
```
and you're good to **GO**

## Structure 
the project is using ``.json`` file as the task list and the CLI will update the ``.json`` file

## Command list

to see the command list available for the program you can use
```
./task-cli
```
this will show list of available command which is ``list``,``add``, ``update``, ``delete``,``mark-done``,``mark-in-progress`` for the CRUD of the tasks.


## Notable commands and information

``list`` will show all the tasks in the task list, and you can also pass in an argument to filter the list by ``status`` such as
```
./task-cli list done
```
this will show all the task in the list with status ``done``

the status itself is divided by 3 category which is ``done``, ``in-progress`` and ``done``

##
**Happy Listing**
##



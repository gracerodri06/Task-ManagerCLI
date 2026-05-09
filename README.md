Welcome to Task Manager app!

What is it?

The Task Manager app is a tool where you can manage your tasks. 
The tasks will be stored in a local file on your computer. This file is created at the first time that you use the tool.

Each task has the following fields:
- ID - unique identifier for the task
- Description
- Status:
      To-do = task was not initialized
      In-Progress = task are being executed
      Done = task were finalized
- Date/time of task creation
- Date/time when a task was updated


How it works?

After starting the app, it will be waiting for you command. Each command has its format. 

This is the command list:

1) "add"
   Adds a new record on the task list, with a current creation date/time and status "To-do".

   Format: "add" + task ID + task description
   Example: add 1 "Go to the supermarket and buy biscuits"

3) "update"
   Updates the description of a record on the task list. Also, the field updatedAt is updated with the current date/time.
   
   Format: "update" + task iD + task description
   Example: update 1 "Go to the supermarket and buy rice"

5) "mark-done"

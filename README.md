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
   Adds a new task on the task list, with a current creation date/time and status "todo". 
   The task ID is generated automatically.

   Format: "add" + task description
   Example: add "Go to the supermarket and buy biscuits"

3) "update"
   Updates the description of a task. Also, updates the field updatedAt with the current date/time.
   
   Format: "update" + task ID + task description
   Example: update 1 "Go to the supermarket and buy rice"

4) "delete"
   Deletes a task from the list.

   Format: "delete" + taskID
   Example: delete 10

5) "mark-in-progress"
   Change the status of a task to in-progress. 

   Format: "mark-in-progress" + task ID
   Example: mark-in-progress 5

6) "mark-done"
   Change the status of a task to done. 

   Format: "mark-done" + task ID
   Example: mark-done 5

7) "list"
   Lists the tasks.
   This command has 4 options available:

   7.1) "list"
         Displays all the tasks from the list.

   7.2) "list todo"
         Displays all the tasks on the list with status "todo".


   7.3) "list in-progress"
         Displays all the tasks on the list with status "in-progress".


   7.4) "list done"
         Displays all the tasks on the list with status "done".

   Example: list in-progress

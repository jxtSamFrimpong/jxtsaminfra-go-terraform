
I want to use golang to build my own custom terraform provider with a resource to change the instance type of my aws cloud ec2 instance currently I havent used goland before but I'm okay with other programming languages, I want this provider to follow OOP principles and keep it simple these are the specs for my custom provider ================================================================= Main Class or however its going to be like takes instance id as input ===================== methods below: 

takes instance id as input
=====================
get_instance(takes intance id) -> instance_info_object: try's to get instance(to check for existence) info from aws
if instance exists return instance info object
if not, stop and throw error

get_instance_state(instance_id) -> instance_state: try's to get the instance's state from aws
if the instance is running, it returns "running"
if its stopped it returns "stopped"
if its terminated it stops and throws an error
if its pending it stops and throws an error
if its in a state that is not one of the mentioned, it stops and throws an error 

get_instance_into_stopped_state(instance_id) --> instance_state: try's to stop the instance on aws
gets instance's state using get_instance_state
if instance is stopped it returns "stopped"
if the instance is "running" it try's to stop the aws instance
   waits for the instance to be stopped
   if successful, it uses get_instance_state to get the instance state again, if its the instance is stopped it returns "stopped" else the code stops and throws an error

get_instance_into_running_state(instance_id) --> instance_state: try's to start an instance on aws
gets instance's state using get_instance_state
if instance is stopped on aws, it try's to start the aws instance
   waits for the instance to be started
   if successful, it uses get_instance_state to get the instance state again, if its the instance is started it returns "running" else the code stops and throws an error
if instance is already running it returns "running"

change_instance_type(instance_id, target_instance_type): --> instance_info: try's to change the instance type of an existing aws instance to a desired type
uses get_instance to get the instances info
  compares the instance type of the existing instance and the target instance type to make sure they are different, if they are the same we stop the program and throw an error
uses get_instance_into_stopped_state to get instance into stopped state, if its successful (thus returns "stopped") we'll print a successful message that the instance was stopped, else we'll print a descriptive error message and throw an error
try's to change the instance type of stopped aws instance into the target_instance_type, 
    waits for it to succeed and print successful message, 
      else print descriptive error message if it doesnt succeed 
regardless of whether the preceeding step succeeds or not run get_instance_into_running_state
print summary message



================================================================= at the end of the day this is what my resource resource in my provider should achieve. PS(I wont be using aws cli access keys cause my environment is already having access to my aws cloud)
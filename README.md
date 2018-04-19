# Command on Hook
Runs a command on webhook from github.


## Environment Variables
- SECRET *(Same as in the webhook settings on Github)*
- COMMAND *(Command to run)*
- ARGUMENTS *(Arguments to add after the command)*
- COMMANDPATH *(Path the command should be run in)*

## Docker

`docker run --rm -it -e SECRET=somereallyobscuresecretthatnobodywillguess -e COMMAND=git -e ARGUMENTS=pull -e COMMANDPATH=/home/project -v ${PWD}:/home/project -v ~/.ssh/:/root/.ssh/ -p 80:80 oisann/commandonhook:latest`

This container runs `git pull` in the /home/project which has the contents of the folder you started the container from. It adds your ssh key, this is needed to pull private repos. The secret is the same as in the webhook settings on the github repo.

![alt text](https://i.imgur.com/QL1fR7K.png "Example Github webhook settings")

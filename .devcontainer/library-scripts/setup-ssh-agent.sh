#!/bin/bash
read -r -d '' SSH_AGENT <<"EOF"
# setup ssh agent
if [ -z "$SSH_AUTH_SOCK" ]; then
   # Check for a currently running instance of the agent
   RUNNING_AGENT="`ps -ax | grep 'ssh-agent -s' | grep -v grep | wc -l | tr -d '[:space:]'`"
   if [ "$RUNNING_AGENT" = "0" ]; then
        # Launch a new instance of the agent
        ssh-agent -s &> $HOME/.ssh/ssh-agent
   fi
   eval `cat $HOME/.ssh/ssh-agent`
fi

cat ~/.ssh_host/id_ed25519 | ssh-add -k -

# empty
EOF

echo "$SSH_AGENT" > tmp.txt

sed -i $'/^# if running bash$/{e cat tmp.txt\n}' ~/.profile

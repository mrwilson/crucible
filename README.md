# Crucible - A simple wrapper around ansible-playbook

## Build Instructions

    go get github.com/mrwilson/crucible

## Uses

Download a playbook from github.

    crucible get <user>/<repo> # downloads to $HOME/.crucible/<user>/<repo>

Run a playbook in the crucible repo.

    crucible <user>/<repo> ... # accepts all ansible-playbook flags

You are advised to follow [best practices](http://www.ansibleworks.com/docs/playbooks_best_practices.html) when laying out your playbooks.

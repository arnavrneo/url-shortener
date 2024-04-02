Docker
=========

Role for running the applications using docker containers.

Requirements
------------

- Ansible installed.
- Rename the `env.example` to `.env` in the directories `frontend`, `expressjs` & `go` under `files` and populate the `.env` with the secrets relevant to specified variables.

Role Variables
--------------

To change the directory on the remote servers where script should be copied, edit the `script_dest` in `main.yml` under `vars` directory.

License
-------

BSD

Author Information
------------------

Author: Arnav Raina (github.com/arnavrneo)

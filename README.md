# Go

## Install
- Method 1: Manual
    1. Download tar.gz from https://golang.org/dl/
    2. Extract:
        ```
        $ sudo -C /usr/local -xzf go1.14.12.linux-amd64.tar.gz
        ```

    3. Add to PATH env var. Edit ```$HOME/.profile```, add the following line:
        ```
        $ export PATH=$PATH:/usr/local/go/bin
        ```

    3. Verify
        ```
        $ go version
        ```

- Method 2: Using Snap
    ```
    $ sudo snap install go --channel=1.14/stable --classic
    ```

## Workspace
1. Install recommended packages
    - ```ctrl + shift + P```
    - type: go install/update tools
    - install tools

2. To fix issues with ```gopls```, just add a new workspace with ```src```

#Go

## Install
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
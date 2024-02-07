# LMS API
###### This repo is rest api backend written in golang for lms app (https://github.com/radifanfariz/lms)
#### How to Run This Repo:
- ### Creating Go Workspace

    Creating Go Workspace by making these 3 folder: *bin*, *pkg*, *src* 
    <br/>

  _src_: The directory that contains Go source files. A source file is a file that you write using the Go programming language. Source files are used by the Go compiler to create an executable binary file.
  _bin_: The directory that contains executables built and installed by the Go tools. Executables are binary files that run on your system and execute tasks. These are typically the programs compiled by your source code or another downloaded Go source code.

    <br/>
    
    >The src subdirectory may contain multiple version control repositories (such as Git, Mercurial, and Bazaar). You will see directories like github.com or golang.org when your program imports third party libraries. If you are using a code repository like github.com, you will also put your projects and source files under that directory. This allows for a canonical import of code in your project. Canonical imports are imports that reference a fully qualified package, such as github.com/digitalocean/godo.

    ```.
    ├── bin
    │   ├── buffalo                                      # command executable
    │   ├── dlv                                          # command executable
    │   └── packr                                        # command executable
    └── src
    └── github.com
        └── digitalocean
            └── godo
                ├── .git                            # Git repository metadata
                ├── account.go                      # package source
                ├── account_test.go                 # test source
                ├── ...
                ├── timestamp.go
                ├── timestamp_test.go
                └── util
                    ├── droplet.go
                    └── droplet_test.go


    ```

- ### Setting GOROOT and GOPATH
  Verify your GOROOT and GOPATH in environment variable

    ```
    $env:GOPATH
    ```

    example output:

    ```
    Output
    C:\Users\sammy\go
    ```

    When Go compiles and installs tools, it will put them in the \$GOPATH/bin directory. For convenience, it’s common to add the workspace’s bin subdirectory to your $PATH. You can do this using the setx command in PowerShell:

    ```
    setx PATH "$($env:path);$GOPATH\bin"
    ```

    This will now allow you to run any programs you compile or download via the Go tools anywhere on your system.

    Now that you have the root of the workspace created and your $GOPATH environment variable set, you will create your future projects with the following directory structure. This example assumes you are using github.com as your repository:

    ```
    $GOPATH/src/github.com/username/project
    ```

    If you were working on the https://github.com/name/go_project project, you would put it in the following directory:

    ```
    $GOPATH/src/github.com/name/go_project
    ```

    Structuring your projects in this manner will make projects available with the go get tool. It will also help readability later.

    You can verify this by using the go get command to fetch the godo library:

    ```
    go get github.com/digitalocean/godo
    ```

    You can see it successfully downloaded the godo package by listing the directory:

    ```
    ls $env:GOPATH/src/github.com/digitalocean/godo
    ```

    You will receive output similar to this:

    ```

    Output
        Directory: C:\Users\sammy\go\src\github.com\digitalocean\godo


    Mode                LastWriteTime         Length Name
    ----                -------------         ------ ----
    d-----        4/10/2019   2:59 PM                util
    -a----        4/10/2019   2:59 PM              9 .gitignore
    -a----        4/10/2019   2:59 PM             69 .travis.yml
    -a----        4/10/2019   2:59 PM           1592 account.go
    -a----        4/10/2019   2:59 PM           1679 account_test.go
    -rw-r--r--  1 sammy  staff   2892 Apr  5 15:56 CHANGELOG.md
    -rw-r--r--  1 sammy  staff   1851 Apr  5 15:56 CONTRIBUTING.md
    .
    .
    .
    -a----        4/10/2019   2:59 PM           5076 vpcs.go
    -a----        4/10/2019   2:59 PM           4309 vpcs_test.go
    ```
- ### Build and run this repo
In your src folder in Go Workspace do it:
```
git clone <main or dev branch of this repo>
cd lms-api
go build
./lms-api
```
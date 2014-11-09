##  Setting up Go

Go is very simple to setup and get coding in as long you decide to let go and give into the go structure. (Written for a linux/unix system)

- Download go from [Google's server](http://golang.org/dl/) I'm currently using 1.3.3 (latest at the time this talk was written)
- unzip the tarball
- set your GOROOT, GOPATH to the correct location (/usr/local/go for Mac OS X)
- Add your GOROOT and GOPATH to your Path env variable


<small>The only thing that I REALLY miss from python to go is the virtualenv setup. It's really nice to be able to isolate your installed libraries from other projects. Go doesn't have a built out answer to this problem right now, so I wrote a couple of shell scripts that make a go virtualenv. It's heavy and probably not the best solution in the world but it works well. It downloads a binrary release of the go tools into a hidden dir of the project. Then sets the $GOPATH and $GOROOT equal to that path. So it requires you source the activate_go_vm.sh in the root directory. It will isolate all of the libs and runtimes to the version you develop against. I recommend use godep to manage the lib versions along with the scripts.</small>

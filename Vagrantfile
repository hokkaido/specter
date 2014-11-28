# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.box = "ubuntu/trusty64"

  config.vm.network "forwarded_port", guest: 8080, host: 8080
  
  config.vm.provision "shell", privileged: false, inline: <<SCRIPT

    sudo apt-get update
    sudo apt-get install git

    # install go
    if [ ! -f /tmp/go1.3.3.linux-amd64.tar.gz ]; then
      sudo wget -O /tmp/go1.3.3.linux-amd64.tar.gz https://storage.googleapis.com/golang/go1.3.3.linux-amd64.tar.gz
      sudo tar -C /usr/local -xzf /tmp/go1.3.3.linux-amd64.tar.gz
    fi

    grep '^export GOPATH' ~/.bashrc || echo export GOPATH=~/go >> ~/.bashrc
    grep '^export PATH' ~/.bashrc || echo export PATH=\$PATH:/usr/local/go/bin:~/go/bin:/vagrant/script >> ~/.bashrc
    GOPATH=~/go go get github.com/tools/godep
    mkdir -p ~/go/src/github.com/hokkaido
    ln -s /vagrant ~/go/src/github.com/hokkaido/specter
    grep ^cd ~/.bashrc || echo cd ~/go/src/github.com/hokkaido/specter >> ~/.bashrc
SCRIPT

  config.vm.provision "docker" do |d|
    d.run "test-1", image: "busybox"
    d.run "test-2", image: "busybox"
  end
  
end
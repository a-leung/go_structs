Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.provision 'shell', inline: <<-SHELL
    echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
    cd /tmp
    curl --silent -O https://storage.googleapis.com/golang/go1.8.4.linux-amd64.tar.gz
    tar -zxf go1.8.4.linux-amd64.tar.gz
    rm go1.8.4.linux-amd64.tar.gz
    mv go /usr/local/

    su - vagrant
    cd /home/vagrant
    mkdir -p /home/vagrant/go/{pkg,bin,src}
    cd /home/vagrant/go/src
    ln -s /vagrant go_structs
    chown -R vagrant:vagrant /home/vagrant/go

    # install git
    add-apt-repository ppa:git-core/ppa -y
    apt-get update
    apt-get install git -y
  SHELL
end

module LocalCommand
  class Config < Vagrant.plugin("2", :config)
    attr_accessor :command
  end

  class Plugin < Vagrant.plugin("2")
    name "local_shell"

    config(:local_shell, :provisioner) do
      Config
    end

    provisioner(:local_shell) do
      Provisioner
    end
  end

  class Provisioner < Vagrant.plugin("2", :provisioner)
    def provision
      result = system "#{config.command}"
    end
  end
end

Vagrant.configure("2") do |config|
  config.vm.box = "eugenmayer/opnsense"

  config.ssh.sudo_command = "%c"
  config.ssh.shell = "/bin/sh"
  config.ssh.password = "opnsense"
  config.ssh.username = "root"
  config.ssh.port = "10022"
  # we need to use rsync, no vbox drivers for bsd
  config.vm.synced_folder ".", "/vagrant", disabled: true

  config.vm.define 'opnsense', autostart: false do |test|
    test.vm.provider 'virtualbox' do |vb|
      vb.customize ['modifyvm',:id, '--nic1', 'intnet', '--nic2', 'nat'] # swap the networks around
      vb.customize ['modifyvm', :id, '--natpf2', "ssh,tcp,127.0.0.1,10022,,22" ] #port forward
      vb.customize ['modifyvm', :id, '--natpf2', "https,tcp,127.0.0.1,10443,,443" ] #port forward
      vb.customize ['modifyvm', :id, '--natpf2', "openvpn,tcp,127.0.0.1,11194,,1194" ] # openvpn
      #vb.customize ['modifyvm', :id, '--natpf1', "https,tcp,127.0.0.1,1443,,443" ] #port forward
    end

    # install dev tools
    test.vm.provision "shell",
      inline: "pkg update && pkg install -y vim-lite joe nano gnu-watch git tmux screen",
      run: "once"

    # replace the public ssh key for the root user with the one vagrant deployed for comms before we restart - or we lock vagrant out
    test.vm.provision "inject-pubkey-into-config", type: "local_shell", command: "export PUB=$(ssh-keygen -f .vagrant/machines/opnsense/virtualbox/private_key -y | base64) && xmlstarlet ed --inplace -u '/opnsense/system/user/authorizedkeys' -v \"$PUB\" config.xml"
    # apply our configuration so we have a configured radius with users and clients and an active openvpn server
    test.vm.provision "file", source: "./config.xml", destination: "/conf/config.xml"
    test.vm.provision "shell",
      inline: "echo 'rebooting to apply config' && reboot"

    test.vm.provision "sleep-for-reboot", type: "local_shell", command: "echo 'waiting for the reboot' && sleep 50"

    test.vm.provision "shell",
      inline: "export openvpn_version=0.0.4 && curl -Lo os-openvpn-devel-${openvpn_version}.txz https://github.com/EugenMayer/opnsense-openvpn-plugin/raw/master/dist/os-openvpn-devel-${openvpn_version}.txz && pkg add os-openvpn-devel-${openvpn_version}.txz"
    test.vm.provision "shell",
      inline: "export unbound_version=0.0.1 && curl -Lo os-unbound-devel-${unbound_version}.txz https://github.com/EugenMayer/opnsense-unbound-plugin/raw/master/dist/os-unbound-devel-${unbound_version}.txz && pkg add os-unbound-devel-${unbound_version}.txz"
  end
end

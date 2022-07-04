# my5G-non3GPP-IoTSDGw

<!-- TODO: replace the repo name below `template` with your repo name -->
![GitHub](https://img.shields.io/github/license/my5G/template?color=blue) 
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/my5G/template)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/my5G/template) 
![GitHub last commit](https://img.shields.io/github/last-commit/my5G/template)
![GitHub contributors](https://img.shields.io/github/contributors/my5G/template)

<img width="12%" src="static/img/my5g-logo.png" alt="my5g-core"/>

----
<!-- TODO: add here general description of the project -->
----
## Description
LoRaWan IoT network integration project via non-3GPP access with 5G Network core

<!-- TODO: add here steps to install the project -->

----
## Installation

**Requirements**
* The installation can be done directly over the host operating system (OS) or inside a virtual machine (VM).
System requirements:
  * CPU type: x86-64 (specific model and number of cores only affect performance)
  * RAM: 4 GB
  * Disk space: 40 GB
  * Ubuntu 18.04 LTS

* Hardware minimum
  * CPU: Intel i5 processor
  * RAM: 4GB
  * Hard drive: 40G
  * NIC card: 1Gbps ethernet card

* Hardware recommended
  * CPU: Intel i7 processor
  * RAM: 8GB
  * Hard drive: 160G
  * NIC card: 10Gbps ethernet card

**Recommended environment**
* Software (minimum) You can use actual versions like Ubuntu 20.04
   * OS: Ubuntu 18.04
   * gcc 7.3.0
   * Go 1.14.4 linux/amd64
   * kernel version 5.0.0-23-generic or higher (for UPF)

    Notes:
    - You can use uname -r to check your current kernel version.
    - Also tested on Ubuntu 20.04 with 5.4.0-53-generic kernel version.
    - You can use go version to check your current Go version.

**Steps**

Install python-minimal:
```
sudo apt update && apt install python-minimal -y
```

Install git:
```
sudo apt -y install git
```

Clone this repository:
```
git clone https://github.com/my5G/my5G-non3GPP-IoTSDGw.git
```

Install Ansible:
```
sudo apt -y install ansible
```

----
**Check**

<!-- TODO: add here steps to test the project --->

**More Information**

<!-- TODO: add here other comments that may be important (Optional) !-->

**Questions**
 
For questions, support, interacting with community members or share new ideas, use the [Discussions](../../discussions). If you want to report a bug or request a new feature, create an [issue](../../issues/new). Please, before creating a new issue, make sure it's not duplicating another existing one.

**Acknowledgments**

<!-- TODO: add here acknowledges to other projects used or external contributors -->

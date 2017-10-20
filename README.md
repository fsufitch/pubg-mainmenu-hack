# PUBG Main Menu MITM Demo 

This repository contains the demo source code for the vulnerability in PUBG's Main Menu, as featured [here](https://medium.com/@fsufitch/playerunknowns-battlegrounds-main-menu-is-vulnerable-to-hacking-d483b00a7036).

A live version can be found at [http://104.239.228.225/index.html](http://104.239.228.225/index.html).

You can visit it with a web browser, or you can fool your own PUBG into loading it by adding the following line to your hosts file:
```
104.239.228.225 front.battlegroundsgame.com
```

In a Windows Environment, the hosts file is found at `%SystemRoot%\System32\drivers\etc\hosts`. Removing the line from the hosts file will disable this demo.

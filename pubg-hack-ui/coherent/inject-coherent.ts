interface PUWindow extends Window{
  _dev_: boolean;
  broConfiguration: any;
  production: boolean;
  engine: any;
}

declare var window: PUWindow;

export function injectCoherent() {
  window._dev_ = false;
  window.broConfiguration = {
    "gate": "https://prod-live-front.playbattlegrounds.com/index.html",
    "uri": "https://prod-live-entry.playbattlegrounds.com/publicproxy",
    "entry": "ws://prod-live-entry.playbattlegrounds.com:81/userproxy",
    "secureEntry": "wss://prod-live-entry.playbattlegrounds.com/userproxy",
    "useSecure": true,
    "credential": {
      "steam": "provider={provider}&ticket={ticket}&playerNetId={playerNetId}&cc={cc}&clientGameVersion={clientGameVersion}",
    },
    "verreq": "2.6.36",
    "LayoutMode": "",
    "LanpartyHome": "http://52.59.158.42/Dreamhack/Game/{accountId}/day1_qualifying/1",
  };
  window.production = true;
  window.engine = {
    BindingsReady: () => <any>undefined,
  };
  require('./coherent.js');
}
